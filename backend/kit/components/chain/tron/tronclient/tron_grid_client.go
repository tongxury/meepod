package tronclient

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
)

func NewTronGridClient(url, apiKey string) (*TronGridClient, error) {
	return &TronGridClient{url: url, apiKey: apiKey}, nil
}

type TronGridClient struct {
	url    string
	apiKey string
}

func (t *TronGridClient) options() httpcli.Options {
	return httpcli.Options{Headers: map[string]string{
		"TRON_PRO_API_KEY": t.apiKey,
		"TRON-PRO-API-KEY": t.apiKey,
	}}
}

func (t *TronGridClient) GetEventsByNum(ctx context.Context, num uint64) (Events, error) {
	//GET https://api.trongrid.io/v1/blocks/47693937/events
	url := fmt.Sprintf("%s/v1/blocks/%d/events?only_confirmed=true&limit=200&fingerprint=", t.url, num)

	var rsp Events

	for {
		resp, err := t.doGetEventsByNum(ctx, url)
		if err != nil {
			return nil, err
		}

		if len(resp.Data) != 0 {
			rsp = append(rsp, resp.Data...)
		}

		if resp.Meta.Fingerprint == "" {
			break
		}

		url = url + resp.Meta.Fingerprint
	}

	return rsp, nil
}

func (t *TronGridClient) doGetEventsByNum(ctx context.Context, url string) (*GridResult[Events], error) {
	resBytes, _, err := httpcli.Client().GET(ctx, url, t.options())
	if err != nil {
		return nil, err
	}

	var rsp GridResult[Events]
	err = conv.J2S(resBytes, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}

//func (t *TronGridClient) GetNowBlock(ctx context.Context) {
//
//	url := fmt.Sprintf("%s/wallet/getnowblock", t.url)
//
//	resBytes, _, err := httpcli.Client().GET(ctx, url, t.options())
//	if err != nil {
//		return nil, err
//	}
//
//	var rsp GridResult[Events]
//	err = conv.J2S(resBytes, &rsp)
//	if err != nil {
//		return nil, err
//	}
//
//	return rsp.Data, nil
//}
