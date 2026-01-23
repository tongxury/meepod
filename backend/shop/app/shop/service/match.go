package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"time"
)

type MatchService struct {
}

func (t *MatchService) ListZ14Matches(ctx context.Context, issue string, page, size int64) (types.Matches, int64, error) {
	dbMatches, total, err := new(db.Match).List(ctx, db.ListMatchesParams{
		Category: enum.MatchCategory_Z14.Value, Issues: []string{issue}, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	matches, err := t.Assemble(ctx, dbMatches)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return matches, total, nil

}

func (t *MatchService) ListZjcMatches(ctx context.Context, page, size int64) (types.Matches, int64, error) {
	dbMatches, total, err := new(db.Match).List(ctx, db.ListMatchesParams{
		Category: enum.MatchCategory_Zjc.Value, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	matches, err := t.Assemble(ctx, dbMatches)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return matches, total, nil

}

func (t *MatchService) ListMatchesByCategoryAndStartTs(ctx context.Context, category string, startTsAfter time.Time) (types.Matches, error) {

	dbMatches, err := new(db.Match).ListMatchesByStartTs(ctx, category, startTsAfter)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return t.Assemble(ctx, dbMatches)
}

func (t *MatchService) ListByCategoryAndIssue(ctx context.Context, category, issue string) (types.Matches, error) {

	dbMatches, err := new(db.Match).FindByCategoryAndIssue(ctx, category, issue)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return t.Assemble(ctx, dbMatches)
}

func (t *MatchService) Assemble(ctx context.Context, matches db.Matches) (types.Matches, error) {

	//issueIds := matches.Ids()
	//
	//issues, err := new(db.Issue).ListMetaByIds(ctx, issueIds)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}
	//issuesTap := issues.AsMap()
	//itemIds := issues.ItemIds()
	//
	//items, err := new(db.Item).ListMetaByIds(ctx, itemIds, false)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}
	//itemsTap := items.AsMap()

	var rsp types.Matches
	for _, x := range matches {

		//issue := issuesTap[x.Issue]
		//item := itemsTap[issue.ItemId]

		y := types.FromDbMatch(x, nil, nil)

		rsp = append(rsp, y)
	}

	return rsp, nil
}
