package adapter

import (
	"fmt"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/dlt"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/f3d"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/pl3"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/pl5"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/ssq"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/x7c"
	"gitee.com/meepo/backend/shop/core/issue/adapter/zc/z14"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"time"
)

type IGenerator interface {
	Generate(fromIndex string, fromTime time.Time) (string, time.Time, time.Time)
}

type IParser interface {
	Parse(src string) ([]any, []any, float64, error)
}

type IPrizer interface {
	Prize(src, target string) (types.PrizeRecords, error)
	FetchTarget(index string) (*types.PrizeResult, error)
}

func Generate(itemId string, lastIssue *types.Issue) (*types.Issue, error) {
	// 第一场手动添加
	if lastIssue == nil {
		return nil, nil
	}

	// 拿到结果之后才开下一期  这样即便是时间算错了  也不影响数据
	if !lastIssue.Prized {
		return nil, nil
	}

	startAt := lastIssue.PrizeAt.Add(time.Second)

	var nextIssue string
	var closeAt time.Time
	var prizeAt time.Time

	switch itemId {
	case enum.ItemId_dlt:
		adapter := dlt.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	case enum.ItemId_f3d:
		adapter := f3d.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	case enum.ItemId_ssq:
		adapter := ssq.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	case enum.ItemId_x7c:
		adapter := x7c.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	case enum.ItemId_pl3:
		adapter := pl3.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	case enum.ItemId_pl5:
		adapter := pl5.Generator{}
		nextIssue, closeAt, prizeAt = adapter.Generate(lastIssue.Index, startAt)
	default:
	}

	return &types.Issue{
		Index:   nextIssue,
		StartAt: startAt,
		CloseAt: closeAt,
		PrizeAt: prizeAt,
	}, nil
}

func Parse(itemId string, src string) ([]any, []any, float64, error) {

	switch itemId {
	case enum.ItemId_dlt:
		adapter := dlt.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_f3d:
		adapter := f3d.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_ssq:
		adapter := ssq.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_x7c:
		adapter := x7c.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_pl3:
		adapter := pl3.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_pl5:
		adapter := pl5.Parser{}
		return adapter.Parse(src)
	case enum.ItemId_rx9:
		adapter := z14.Parser{Min: 9}
		return adapter.Parse(src)
	case enum.ItemId_sfc:
		adapter := z14.Parser{Min: 14}
		return adapter.Parse(src)
	default:
	}

	return nil, nil, 0, fmt.Errorf("unknown itemId: %s", itemId)
}

func Prize(itemId, src, target string) (types.PrizeRecords, error) {
	switch itemId {
	case enum.ItemId_dlt:
		adapter := dlt.Prizer{}
		return adapter.Prize(src, target)
	case enum.ItemId_f3d:
		adapter := f3d.Prizer{}
		return adapter.Prize(src, target)
	case enum.ItemId_ssq:
		adapter := ssq.Prizer{}
		return adapter.Prize(src, target)
	case enum.ItemId_x7c:
		adapter := x7c.Prizer{}
		return adapter.Prize(src, target)
	case enum.ItemId_pl3:
		adapter := pl3.Prizer{}
		return adapter.Prize(src, target)
	case enum.ItemId_pl5:
		adapter := pl5.Prizer{}
		return adapter.Prize(src, target)
	default:
	}

	return nil, fmt.Errorf("unknown itemId: %s", itemId)
}

func FetchTarget(itemId, index string) (*types.PrizeResult, error) {
	switch itemId {
	case enum.ItemId_dlt:
		adapter := dlt.Prizer{}
		return adapter.FetchTarget(index)
	case enum.ItemId_f3d:
		adapter := f3d.Prizer{}
		return adapter.FetchTarget(index)
	case enum.ItemId_ssq:
		adapter := ssq.Prizer{}
		return adapter.FetchTarget(index)
	case enum.ItemId_x7c:
		adapter := x7c.Prizer{}
		return adapter.FetchTarget(index)
	case enum.ItemId_pl3:
		adapter := pl3.Prizer{}
		return adapter.FetchTarget(index)
	case enum.ItemId_pl5:
		adapter := pl5.Prizer{}
		return adapter.FetchTarget(index)
	default:
	}

	return nil, fmt.Errorf("unknown itemId: %s", itemId)
}
