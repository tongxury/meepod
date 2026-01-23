package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/go-pg/pg/v10"
	"time"
)

type Match struct {
	tableName    struct{} `pg:"t_matches"`
	Id           string
	League       string
	HomeTeam     string
	HomeTeamTag  string
	GuestTeam    string
	GuestTeamTag string
	Category     string
	Issue        string
	StartAt      time.Time
	CloseAt      time.Time
	CreatedAt    time.Time
	Status       string
	Odds         types.Odds
	RCount       int
	RealOdds     types.Odds
	Result       types.Result
}

type Matches []*Match

func (ts Matches) ToIssueMatches() types.Matches {
	var rsp types.Matches

	for _, x := range ts {

		y := types.Match{
			Id:           x.Id,
			League:       x.League,
			HomeTeam:     x.HomeTeam,
			HomeTeamTag:  x.HomeTeamTag,
			GuestTeam:    x.GuestTeam,
			GuestTeamTag: x.GuestTeamTag,
			Category:     x.Category,
			Issue:        x.Issue,
			StartAt:      x.StartAt,
			CloseAt:      x.CloseAt,
			CreatedAt:    x.CreatedAt,
			Status:       x.Status,
			Odds:         x.Odds,
			RCount:       x.RCount,
			RealOdds:     x.RealOdds,
			Result:       x.Result,
		}
		rsp = append(rsp, &y)
	}

	return rsp
}

func (t *Match) FromIssueMatches(matches types.Matches) Matches {

	var rsp Matches

	for _, x := range matches {

		y := Match{
			Id:           x.Id,
			League:       x.League,
			HomeTeam:     x.HomeTeam,
			HomeTeamTag:  x.HomeTeamTag,
			GuestTeam:    x.GuestTeam,
			GuestTeamTag: x.GuestTeamTag,
			Category:     x.Category,
			Issue:        x.Issue,
			StartAt:      x.StartAt,
			CloseAt:      x.CloseAt,
			CreatedAt:    x.CreatedAt,
			Status:       x.Status,
			Odds:         x.Odds,
			RCount:       x.RCount,
			RealOdds:     x.RealOdds,
			Result:       x.Result,
		}
		rsp = append(rsp, &y)
	}

	return rsp
}

func (ts Matches) GroupByIssue() map[string]Matches {

	rsp := make(map[string]Matches, len(ts))
	for _, x := range ts {
		rsp[x.Issue] = append(rsp[x.Issue], x)
	}

	return rsp
}

func (t *Match) UpsertMatches(ctx context.Context, matches Matches) error {

	if len(matches) == 0 {
		return nil
	}

	_, err := comp.SDK().Postgres().Model(&matches).Context(ctx).
		OnConflict("(issue, home_team, guest_team, category) do nothing").
		Insert()
	if err != nil {
		return err
	}

	return nil

}

func (t *Match) UpsertOdds(ctx context.Context) error {

	if t == nil {
		return nil
	}

	update, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("home_team = ?", t.HomeTeam).
		Where("guest_team = ?", t.GuestTeam).
		Where("start_at = ?", t.StartAt).
		Where("category = ?", t.Category).
		Set(fmt.Sprintf("odds = jsonb_concat(odds, '%s')", conv.S2J(t.Odds))).
		Update()
	if err != nil {
		return err
	}

	if update.RowsAffected() == 0 {
		_, err := comp.SDK().Postgres().Model(t).Context(ctx).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ts Matches) UpsertBatch(ctx context.Context) error {

	if len(ts) == 0 {
		return nil
	}

	_, err := comp.SDK().Postgres().Model(&ts).Context(ctx).
		OnConflict("(issue, home_team, guest_team, category) do update").
		Insert()
	if err != nil {
		return err
	}

	return nil
}

func (t *Match) UpdateResult(ctx context.Context, tx *pg.Tx, match *types.Match) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Set("status = ?", enum.MatchStatus_End.Value).
		Set("result = ?", conv.S2J(match.Result)).
		Set("real_odds = ?", match.RealOdds).
		Where("issue = ?", match.Issue).
		Where("category = ?", match.Category).
		Where("home_team = ?", match.HomeTeam).
		Where("guest_team = ?", match.GuestTeam).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

type ListMatchesParams struct {
	Category   string
	Issues     []string
	Page, Size int64
}

func (t *Match) List(ctx context.Context, params ListMatchesParams) (Matches, int64, error) {

	var rsp Matches
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("status != ?", enum.MatchStatus_Pending.Value)

	if params.Category != "" {
		q = q.Where("category = ?", params.Category)
	}

	if len(params.Issues) > 0 {
		q = q.WhereIn("issue in (?)", params.Issues)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	total, err := q.OrderExpr("start_at desc").SelectAndCount()
	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(total), nil
}

func (t *Match) FindByCategoryAndIssue(ctx context.Context, category, index string) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Where("issue = ?", index).
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) ListMatchesByCloseTs(ctx context.Context, category string, closeTsAfter time.Time) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Where("close_at > ?", closeTsAfter).
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) ListMatchesByStartTs(ctx context.Context, category string, startTsAfter time.Time) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Where("start_at > ?", startTsAfter).
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) FindShouldBeEndZ14Matches(ctx context.Context, limit int) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", enum.MatchCategory_Z14.Value).
		Where("status = ?", enum.MatchStatus_UnStart.Value).
		Where("start_at < ?", time.Now().Add(-100*time.Minute)).
		Limit(limit).
		OrderExpr("issue").
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) FindShouldBeEndZcMatches(ctx context.Context, limit int) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", enum.MatchCategory_Zjc.Value).
		Where("status = ?", enum.MatchStatus_UnStart.Value).
		Where("start_at < ?", time.Now().Add(-100*time.Minute)).
		Limit(limit).
		OrderExpr("issue").
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

//func (t *Match) UpdateResult(ctx context.Context, tx *pg.Tx) (bool, error) {
//
//	update, err := tx.Model(t).Context(ctx).
//		Set("status = ?", enum.MatchStatus_End.Value).
//		Set("real_odds = ?", conv.S2J(t.Odds)).
//		Set("result = ?", conv.S2J(t.Result)).
//		Where("home_team = ?", t.HomeTeam).
//		Where("guest_team = ?", t.GuestTeam).
//		Where("start_at = ?", t.StartAt).
//		Update()
//	if err != nil {
//		return false, err
//	}
//
//	return update.RowsAffected() > 0, nil
//}

func (t *Match) FindPendingMatches(ctx context.Context, category string) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Where("status = ?", enum.MatchStatus_Pending.Value).
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) UpdateStatus(ctx context.Context, tx *pg.Tx, category string, issues []string, status string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).Set("status = ?", status).
		Where("category = ?", category).
		WhereIn("issue in (?)", issues).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Match) FindLatest(ctx context.Context, category string) (Matches, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Order("issue desc").Limit(1).
		Select()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Match) FindLatestIssue(ctx context.Context, category string) (string, error) {

	var rsp Matches
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("category = ?", category).
		Order("issue desc").Limit(1).
		Select()
	if err != nil {
		return "", err
	}

	if len(rsp) == 0 {
		return "", nil
	}

	return rsp[0].Issue, nil
}
