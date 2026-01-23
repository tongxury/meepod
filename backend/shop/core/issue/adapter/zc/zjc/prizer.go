package zjc

import "gitee.com/meepo/backend/shop/core/issue/types"

type Prizer struct {
}

func (t *Prizer) Prize(src string, matchTarget types.MatchTarget) (types.PrizeRecords, error) {
	return nil, nil
}
