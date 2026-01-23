package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
)

type OrderGroupShareService struct {
	service.OrderGroupShareService
}

func (t *OrderGroupShareService) ListGroupShares(ctx context.Context, groupId, keeperId, storeId string) (types.GroupShares, error) {

	dbShares, err := new(db.OrderGroupShare).FindByGroupIdAndStatus(ctx, groupId, []string{enum.OrderGroupShareStatus_Payed.Value})
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	shares, err := t.Assemble(ctx, dbShares)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return shares, nil

}
