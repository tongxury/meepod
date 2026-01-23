package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
)

type AlipayService struct {
}

func (t *AlipayService) SaveAuthToken(ctx context.Context, authCode string) error {
	authToken, err := comp.SDK().Alipay().GetAppAuthToken(ctx, authCode)
	if err != nil {
		return xerror.Wrap(err)
	}

	token := authToken.AlipayOpenAuthTokenAppResponse.Tokens[0]

	err = new(db.Store).UpdateAuthToken(ctx, token.UserId, token)
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *AlipayService) GetAuthToken(ctx context.Context, storeId string) (string, error) {

	store, err := new(db.Store).RequireByStoreId(ctx, storeId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	if store.Alipay == nil || store.Alipay.AuthToken == nil {
		return "", xerror.Wrapf("no token found")
	}

	return store.Alipay.AuthToken.AppAuthToken, nil
}
