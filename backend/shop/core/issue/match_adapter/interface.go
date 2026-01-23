package match_adapter

import "gitee.com/meepo/backend/shop/core/issue/types"

type Adapter interface {
	ListZ14Matches(index string) (types.Matches, error)
	ListZ14MatchResults(index string) (types.Matches, types.PrizeGrades, error)
	ListZjcMatches(index string) (types.Matches, error)
	ListZjcMatchResults(index string) (types.Matches, types.PrizeGrades, error)
}
