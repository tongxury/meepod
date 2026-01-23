package fetcher_adapter

import "gitee.com/meepo/backend/shop/core/issue/types"

type Adapter interface {
	FindDltResultByIndex(index string) (*types.DltResult, error)
	FindSsqResultByIndex(index string) (*types.SsqResult, error)
	FindX7cResultByIndex(index string) (*types.X7cResult, error)
	FindF3dResultByIndex(index string) (*types.F3dResult, error)
	FindPl3ResultByIndex(index string) (*types.Pl3Result, error)
	FindPl5ResultByIndex(index string) (*types.Pl5Result, error)
}
