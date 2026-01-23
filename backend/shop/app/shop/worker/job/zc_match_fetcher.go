package job

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/match_adapter"
	"github.com/go-pg/pg/v10"
	"golang.org/x/net/context"
	"time"
)

func FetchZjcMatches() {
	ctx := context.Background()

	matches, err := match_adapter.GetLotteryAdapter().ListZjcMatches("")
	if err != nil {
		slf.WithError(err).Errorw("fetchZcMatchesV1 err")
		return
	}

	err = new(db.Match).UpsertMatches(ctx, new(db.Match).FromIssueMatches(matches))
	if err != nil {
		slf.WithError(err).Errorw("UpsertBatch err ")
		return
	}

}
func FetchZ14Matches() {

	ctx := context.Background()

	//issue, err := new(db.Match).FindLatestIssue(ctx, enum.MatchCategory_Z14.Value)
	//if err != nil {
	//	slf.WithError(err).Errorw("FindLatestIssue err")
	//	return
	//}
	//
	//nextIssue := conv.String(conv.Int(issue) + 1)

	// 队伍名称有区别  不能直接替换 adapter
	matches, err := match_adapter.GetLotteryAdapter().ListZ14Matches("")
	//matches, err := match_adapter.GetOkoooAdapter().ListZ14Matches("")
	if err != nil {
		slf.WithError(err).Errorw("ListZ14Matches err")
		return
	}

	err = new(db.Match).UpsertMatches(ctx, new(db.Match).FromIssueMatches(matches))
	if err != nil {
		slf.WithError(err).Errorw("UpsertMatches err ")
		return
	}
}

func FetchZjcMatchesResult() {

	ctx := context.Background()

	toMatches, err := new(db.Match).FindShouldBeEndZcMatches(ctx, 1)
	if err != nil {
		slf.WithError(err).Errorw("FindShouldBeEndZcMatches err")
		return
	}

	if len(toMatches) == 0 {
		return
	}

	toMatch := toMatches[0]

	matches, _, err := match_adapter.GetLotteryAdapter().ListZjcMatchResults(toMatch.Issue)
	if err != nil {
		return
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		for _, x := range matches {
			_, err := new(db.Match).UpdateResult(ctx, tx, x)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		slf.WithError(err).Errorw("")
		return
	}
}
func FetchZ14MatchesResult() {

	ctx := context.Background()

	toMatches, err := new(db.Match).FindShouldBeEndZ14Matches(ctx, 1)
	if err != nil {
		slf.WithError(err).Errorw("FindShouldByEndMatches err")
		return
	}

	if len(toMatches) == 0 {
		return
	}

	toMatch := toMatches[0]

	results, prizeGrades, err := match_adapter.GetLotteryAdapter().ListZ14MatchResults(toMatch.Issue)
	if err != nil {
		return
	}

	if len(results) == 0 {
		return
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		for _, x := range results {
			_, err := new(db.Match).UpdateResult(ctx, tx, x)
			if err != nil {
				return err
			}
		}

		_, err = new(db.Issue).UpdatePrizeResult(ctx, tx, enum.ItemId_rx9+"-"+toMatch.Issue, "",
			prizeGrades.FilterItemId(enum.ItemId_rx9), time.Now())
		if err != nil {
			return err
		}

		_, err = new(db.Issue).UpdatePrizeResult(ctx, tx, enum.ItemId_sfc+"-"+toMatch.Issue, "",
			prizeGrades.FilterItemId(enum.ItemId_sfc), time.Now())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		slf.WithError(err).Errorw("UpdateResults err")
		return
	}

}
