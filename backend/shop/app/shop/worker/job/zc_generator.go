package job

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/go-pg/pg/v10"
	"golang.org/x/net/context"
)

func GenerateZ14() {

	ctx := context.Background()
	// 获取 matches 表中同步到的状态为pending 最新数据和期数 生成对应的issue数据
	pendingMatches, err := new(db.Match).FindPendingMatches(ctx, enum.MatchCategory_Z14.Value)
	if err != nil {
		slf.WithError(err).Errorw("FindPendingMatches err")
		return
	}

	// 未同步到
	if len(pendingMatches) == 0 {
		return
	}

	// 可能拿到多个issue
	groups := map[string]db.Matches{}

	for _, x := range pendingMatches {
		groups[x.Issue] = append(groups[x.Issue], x)
	}

	// 开始时间默认为爬取到场次的时间 todo

	var issues db.Issues
	for x, matches := range groups {

		y1 := db.Issue{
			Id:        enum.ItemId_rx9 + "-" + x,
			ItemId:    enum.ItemId_rx9,
			Index:     x,
			Result:    "",
			PrizedAt:  matches[0].CloseAt, // todo
			CloseAt:   matches[0].CloseAt,
			StartedAt: matches[0].CreatedAt,
		}

		y2 := db.Issue{
			Id:        enum.ItemId_sfc + "-" + x,
			ItemId:    enum.ItemId_sfc,
			Index:     x,
			Result:    "",
			PrizedAt:  matches[0].CloseAt, // todo
			CloseAt:   matches[0].CloseAt,
			StartedAt: matches[0].CreatedAt,
		}

		issues = append(issues, &y1, &y2)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		err := issues.UpsertBatch(ctx, tx)
		if err != nil {
			return err
		}

		_, indexes := issues.Ids()

		_, err = new(db.Match).UpdateStatus(ctx, tx, enum.MatchCategory_Z14.Value, indexes, enum.MatchStatus_UnStart.Value)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		slf.WithError(err).Errorw("Insert Issues err")
	}

}
