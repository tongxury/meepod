package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
)

type OrderGroupShareService struct{}

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

func (t *OrderGroupShareService) Assemble(ctx context.Context, shares db.OrderGroupShares) (types.GroupShares, error) {
	_, groupIds, userIds := shares.Ids()

	dbGroups, err := new(db.OrderGroup).ListByIds(ctx, groupIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbGroupMap := dbGroups.AsMap()
	_, _, groupUserIds, _, _ := dbGroups.Ids()
	userIds = append(userIds, groupUserIds...)

	dbUsers, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	userMap := dbUsers.AsMap()

	var rsp types.GroupShares

	for _, x := range shares {

		user := types.FromDbUser(userMap[x.UserId])

		dbGroup := dbGroupMap[x.GroupId]

		if x.UserId == dbGroup.UserId {
			user.Tags = append(user.Tags, &types.Tag{
				Title: "发起人",
				Color: "orange",
			})
		}

		//group := model.FromDbGroup(dbGroup, userMap[dbGroup.UserId])

		//new(OrderGroupService).GetGroup(ctx, id)

		y := types.GroupShare{
			Id:          x.Id,
			User:        user,
			Volume:      x.Volume,
			Amount:      x.Amount,
			Group:       &types.OrderGroup{Id: x.GroupId},
			CreatedAtTs: x.CreatedAt.Unix(),
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			Status:      enum.OrderGroupShareStatus(x.Status),
			//RewardSummary: x.Extra.
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil
}
