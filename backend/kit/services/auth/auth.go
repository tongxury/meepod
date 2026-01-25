package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/digest"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Service struct {
}

func (t *Service) SendCode(ctx context.Context, phone, category, sign, templateCode string, testPrefix string) (string, bool, error) {

	code := mathd.RandNumber(1000, 9999)

	// isTest := strings.HasPrefix(phone, testPrefix)
	isTest := true

	if !isTest {
		//resp, err := comp.SDK().AliSMS().Send(ctx, []string{phone}, conv.String(code), sign, templateCode)
		//if err != nil {
		//	return "", false, err
		//}
		//
		//if !resp.IsSuccess() {
		//	return "", false, fmt.Errorf(resp.String())
		//}

		if !comp.SDK().Redis().SetNX(ctx, category+phone, 1, 5*time.Minute).Val() {
			return "", false, errorx.UserMessage("发送频繁，请稍后")

		}

		resp, err := comp.SDK().UCloudSMS().Send(ctx, phone, conv.String(code), sign, templateCode)
		if err != nil {
			return "", false, err
		}

		if resp.GetRetCode() != 0 {
			return "", false, fmt.Errorf(resp.GetMessage())
		}

	}

	key := "auth.code:" + phone

	_, err := comp.SDK().Redis().Set(ctx, key, code, 5*time.Minute).Result()
	if err != nil {
		return "", false, xerror.Wrap(err)
	}

	return conv.String(code), isTest, nil
}

func (t *Service) FindCode(ctx context.Context, phone string) (string, error) {

	key := "auth.code:" + phone
	code, err := comp.SDK().Redis().Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", xerror.Wrap(err)
	}

	return code, nil
}

func (t *Service) GenerateTokenByCode(ctx context.Context, phone, code string, checkUser CheckUserFunc) (string, error) {

	expectedCode, err := t.FindCode(ctx, phone)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	fmt.Println("expectedCode", expectedCode)
	fmt.Println("code", code)
	if expectedCode != code {
		return "", errorx.UserMessage("验证码错误")
	}

	userId, err := checkUser(ctx, phone)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"ts":      time.Now().Unix(),
	})
	tokenString, err := token.SignedString(enum.SecretSignKey)

	//token := digest.MD5(phone + time.Now().String())
	//
	//err = comp.SDK().Redis().Set(ctx, "auth.token:"+token, userId, 30*24*time.Hour).Err()
	if err != nil {
		return "", xerror.Wrap(err)
	}

	//comp.SDK().Redis().Del(ctx, "auth.code:"+phone)

	return tokenString, nil
}

func (t *Service) GenerateTokenByPassword(ctx context.Context, phone, password string, checkUser CheckUserPasswordFunc) (string, error) {

	userId, err := checkUser(ctx, phone, password)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	token := digest.MD5(phone + time.Now().String())

	err = comp.SDK().Redis().Set(ctx, "auth.token:"+token, userId, 30*24*time.Hour).Err()
	if err != nil {
		return "", xerror.Wrap(err)
	}

	comp.SDK().Redis().Del(ctx, "auth.code:"+phone)

	return token, nil
}

func (t *Service) FindUserIdByToken(ctx context.Context, token string) (string, error) {

	userId, err := comp.SDK().Redis().Get(ctx, "auth.token:"+token).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", xerror.Wrap(err)
	}

	return userId, nil
}
