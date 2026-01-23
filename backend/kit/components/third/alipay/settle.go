package alipay

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"github.com/go-resty/resty/v2"
	"time"
)

type SettleParams struct {
	AppAuthToken      string
	OrderId           string
	TradeNo           string
	TotalAmount       float64
	RoyaltyParameters []RoyaltyParameters
}

type RoyaltyParameters struct {
	RoyaltyType  string  `json:"royalty_type"`
	TransOut     string  `json:"trans_out,omitempty"`
	TransOutType string  `json:"trans_out_type,omitempty"`
	TransInType  string  `json:"trans_in_type"`
	TransIn      string  `json:"trans_in"`
	Amount       float64 `json:"amount"`
	Desc         string  `json:"desc"`
	RoyaltyScene string  `json:"royalty_scene"`
	TransInName  string  `json:"trans_in_name"`
}

//"      \"royalty_type\":\"transfer\"," +
//"      \"trans_out\":\"2088101126765726\"," +
//"      \"trans_out_type\":\"userId\"," +
//"      \"trans_in_type\":\"userId\"," +
//"      \"trans_in\":\"2088101126708402\"," +
//"      \"amount\":0.1," +
//"      \"desc\":\"分账给2088101126708402\"," +
//"      \"royalty_scene\":\"达人佣金\"," +
//"      \"trans_in_name\":\"张三\"" +

func (t *Client) Settle(ctx context.Context, params SettleParams) (*T4, error) {

	queries := t.baseQueries()
	queries.Set("method", "alipay.trade.order.settle")
	queries.Set("app_auth_token", params.AppAuthToken)
	queries.Set("biz_content", conv.M2J(map[string]any{
		"out_request_no":     conv.String(time.Now().UnixNano()),
		"trade_no":           params.TradeNo,
		"royalty_parameters": params.RoyaltyParameters,
		"royalty_mode":       "sync",
		"extend_params": map[string]string{
			"royalty_finish": "true",
		},
	}))

	queries.Set("sign", t.genSign(queries))

	response, err := resty.New().R().SetContext(ctx).Get(t.options.Entrypoint + "?" + queries.Encode())
	if err != nil {
		return nil, err
	}

	var resp T4
	err = conv.B2S(response.Body(), &resp)
	if err != nil {
		return nil, err
	}

	if resp.AlipayTradeOrderSettleResponse.Code != "10000" {
		return nil, fmt.Errorf("%s %s", resp.AlipayTradeOrderSettleResponse.SettleNo,
			resp.AlipayTradeOrderSettleResponse.Msg)
	}

	return &resp, nil
}

type T4 struct {
	AlipayTradeOrderSettleResponse struct {
		Code     string `json:"code"`
		Msg      string `json:"msg"`
		TradeNo  string `json:"trade_no"`
		SettleNo string `json:"settle_no"`
	} `json:"alipay_trade_order_settle_response"`
	Sign string `json:"sign"`
}

type BindParams struct {
	Receivers    []Receiver
	AppAuthToken string
}

type Receiver struct {
	Type          string `json:"type"`
	Account       string `json:"account"`
	Name          string `json:"name"`
	Memo          string `json:"memo"`
	LoginName     string `json:"login_name"`
	BindLoginName string `json:"bind_login_name"`
}

func (t *Client) Bind(ctx context.Context, params BindParams) error {

	queries := t.baseQueries()
	queries.Set("method", "alipay.trade.royalty.relation.bind")
	queries.Set("app_auth_token", params.AppAuthToken)
	queries.Set("biz_content", conv.M2J(map[string]any{
		"receiver_list":  params.Receivers,
		"out_request_no": conv.String(time.Now().UnixNano()),
	}))

	queries.Set("sign", t.genSign(queries))

	response, err := resty.New().R().SetContext(ctx).Get(t.options.Entrypoint + "?" + queries.Encode())
	if err != nil {
		return err
	}

	var resp T3
	err = conv.B2S(response.Body(), &resp)
	if err != nil {
		return err
	}

	if resp.AlipayTradeRoyaltyRelationBindResponse.Code != "10000" {
		return fmt.Errorf("%s %s", resp.AlipayTradeRoyaltyRelationBindResponse.ResultCode, resp.AlipayTradeRoyaltyRelationBindResponse.Msg)
	}

	return nil
}

type T3 struct {
	AlipayTradeRoyaltyRelationBindResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		ResultCode string `json:"result_code"`
	} `json:"alipay_trade_royalty_relation_bind_response"`
	Sign string `json:"sign"`
}
