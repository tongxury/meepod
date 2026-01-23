package auth_comp

import (
	"context"
	"gitee.com/meepo/backend/kit/components/auth/verification"
	"gitee.com/meepo/backend/kit/components/sdk/digest"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"time"
)

type IAuthComp interface {
	SendProof(ctx context.Context, authKey string) (string, error)
	GenerateToken(ctx context.Context, authKey, proof string) (string, string, error)
	FindAuthIdByToken(ctx context.Context, token string) (string, error)
}

type IAuthStore interface {
	// 通过用户填入的手机号或者email等 获取到用户的详细信息
	FindAuthIdByAuthKey(ctx context.Context, authKey string) (string, error)

	// 保存登录用户的Token和Id的对应关系, 用户后续可通过Token进行接口访问
	SaveAuthIdAndTokenMapping(ctx context.Context, token, authId string) error
	// 通过Token拿到登录用户的ID
	FindAuthIdByToken(ctx context.Context, token string) (string, error)
}

func Assemble(authVerifyComp verification_comp.IAuthVerifyComp, authStore IAuthStore) IAuthComp {
	return &instance{
		authVerifyComp: authVerifyComp,
		authStore:      authStore,
	}
}

type instance struct {
	authVerifyComp verification_comp.IAuthVerifyComp
	authStore      IAuthStore
}

func (i *instance) SendProof(ctx context.Context, authKey string) (string, error) {
	return i.authVerifyComp.SendProof(ctx, authKey)
}

func (i *instance) GenerateToken(ctx context.Context, authKey, proof string) (string, string, error) {
	if reason, err := i.authVerifyComp.VerifyProof(ctx, authKey, proof); err != nil {
		return "", reason, xerror.Wrap(err)
	}

	id, err := i.authStore.FindAuthIdByAuthKey(ctx, authKey)
	if err != nil {
		return "", "", xerror.Wrap(err)
	}

	token := digest.MD5(authKey + time.Now().String())

	if err := i.authStore.SaveAuthIdAndTokenMapping(ctx, token, id); err != nil {
		return "", "", xerror.Wrap(err)
	}

	return token, "", nil
}

func (i *instance) FindAuthIdByToken(ctx context.Context, token string) (string, error) {
	return i.authStore.FindAuthIdByToken(ctx, token)
}
