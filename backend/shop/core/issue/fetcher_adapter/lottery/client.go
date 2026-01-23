package lottery

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/helper"
)

func client() *lClient {
	return &lClient{}
}

type lClient struct {
}

func (t *lClient) FindByIndex(gameNo, index string) (*Result, error) {

	url := "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=%s&provinceId=0&pageSize=30&isVerify=1&pageNo=1&startTerm=%s&endTerm=%s"

	url = fmt.Sprintf(url, gameNo, index, conv.String(conv.Int(index)+1))

	resultBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var r HttpResult
	err = conv.B2S[HttpResult](resultBytes, &r)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if !r.Success {
		return nil, nil
	}

	for _, x := range r.Value.List {

		if x.LotteryDrawNum != index {
			continue
		}

		//x := result.Value.ListStoreItems[len(result.Value.ListStoreItems)-1]

		// 奖金可能有延迟
		if len(x.PrizeLevelList) == 0 {
			continue
		}

		return &x, nil
	}

	return nil, nil
}
