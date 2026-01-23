package adapter

import (
	"fmt"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter/zc/z14"
	"gitee.com/meepo/backend/shop/core/issue/adapter/zc/zjc"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type IZcPrizer interface {
	Prize(src string, matchTarget types.MatchTarget) (types.PrizeRecords, error)
}

func PrizeZc(itemId, src string, matchTarget types.MatchTarget) (types.PrizeRecords, error) {
	switch itemId {
	case enum.ItemId_rx9:
		adapter := z14.Rx9Prizer{}
		return adapter.Prize(src, matchTarget)
	case enum.ItemId_sfc:
		adapter := z14.SfcPrizer{}
		return adapter.Prize(src, matchTarget)
	case enum.ItemId_zjc:
		adapter := zjc.Prizer{}
		return adapter.Prize(src, matchTarget)
	default:
	}

	return nil, fmt.Errorf("unknown itemId: %s", itemId)
}
