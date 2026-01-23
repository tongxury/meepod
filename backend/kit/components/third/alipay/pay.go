package alipay

import (
	"context"
	"crypto"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/encryptor"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"github.com/go-resty/resty/v2"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	options Options
}

func NewAlipayClient(options Options) *Client {
	return &Client{
		options: options,
	}
}

func (t *Client) CheckSign(params url.Values, sign string) error {

	params.Del("sign_type")
	params.Del("sign")

	toSign := t.toSign(params)

	return encryptor.RsaCheckSign(toSign, sign, t.options.AlipayPublicKey, crypto.SHA256)
}

//curl http://localhost:6066/api/payment/v1/topup-callback -d 'gmt_create=2023-07-06 14:29:34&charset=utf-8&seller_email=jmtechcor@163.com&subject=购买&sign=gKCr9pSjHozi4+Sd/VbmYLuNUVcDK3Nkum8yPIpvg+sAfyz7F819odhZ/O7XjNB2I4eY3KwqICNo84Y//zZWBqfSTsrctddYGFCn6gOXuKD380ucijg6Dah1YoQzfzVEkkpf30F439D86ZX/LCHvNEe6cExyZqzyjrdDgcYTYS8kfiQ8o1pEsZyYxJ7ZgMSj0Vt9qFqx/b/otc93qSisvkii6Zr78v3AFuHw5C1UzvdTKIgqq4if85JqHJnLYE/pBRhbFCqX74Bfle/AJZYOEcMfm2hfvf48/mkGTLeK9js/97iQJX5x78sC7ccJhtXnajao6j1dS0rZGRBKdEBA==&buyer_id=2088642537420913&invoice_amount=1.00&notify_id=2023070601222142936020911477706832&fund_bill_list=[{"amount":"1.00","fundChannel":"ALIPAYACCOUNT"}]&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=1.00&buyer_pay_amount=1.00&app_id=2021004101679599&sign_type=RSA2&seller_id=2088641357051477&gmt_payment=2023-07-06 14:29:35&notify_time=2023-07-06 14:43:13&passback_params={"amount":1,"orderId":"14","storeId":"1qaz","subject":"购买","userId":"14"}&version=1.0&out_trade_no=14&tol_amount=1.00&trade_no=2023070622001420911402182536&auth_app_id=2021004102670035&buyer_logon_id=133****0292&point_amount=0.00'

func (t *Client) genSign(params url.Values) string {

	toSign := t.toSign(params)

	return encryptor.RsaSign(toSign, t.options.PrivateKey, crypto.SHA256)
}

func (t *Client) toSign(params url.Values) string {

	keys := helper.Keys(params)
	sort.Strings(keys)

	var parts []string

	for _, key := range keys {
		if key != "" && len(params[key]) > 0 && params[key][0] != "" {

			v, _ := url.PathUnescape(params[key][0])
			parts = append(parts, fmt.Sprintf("%s=%s", key, v))
		}
	}

	return strings.Join(parts, "&")
}

func (t *Client) baseQueries() url.Values {

	queries := url.Values{}

	queries.Set("app_id", t.options.AppId)
	queries.Set("format", "json")
	queries.Set("charset", "utf-8")
	queries.Set("sign_type", "RSA2")
	queries.Set("timestamp", time.Now().Format(time.DateTime))
	queries.Set("version", "1.0")

	return queries
}

type TradeParams struct {
	NotifyUrl      string
	AppAuthToken   string
	OrderId        string
	TotalAmount    float64
	Subject        string
	TimeExpire     string // 1m~15d
	PassbackParams map[string]any
	ProviderPID    string
}

func (t *Client) GenerateTradeQrCode(ctx context.Context, params TradeParams) (string, error) {

	if params.TotalAmount <= 0 {
		return "", nil
	}

	queries := t.baseQueries()

	queries.Set("method", "alipay.trade.precreate")
	queries.Set("app_auth_token", params.AppAuthToken)
	queries.Set("notify_url", params.NotifyUrl)
	queries.Set("biz_content", conv.M2J(map[string]any{
		"out_trade_no":    params.OrderId,
		"total_amount":    params.TotalAmount,
		"time_expire":     params.TimeExpire,
		"subject":         params.Subject,
		"passback_params": conv.M2J(params.PassbackParams),
		"extend_params":   map[string]any{"sys_service_provider_id": params.ProviderPID},
	}))

	queries.Set("sign", t.genSign(queries))

	response, err := resty.New().R().SetContext(ctx).Get(t.options.Entrypoint + "?" + queries.Encode())
	if err != nil {
		return "", err
	}

	var resp T2
	err = conv.B2S(response.Body(), &resp)
	if err != nil {
		return "", err
	}

	return resp.AlipayTradePrecreateResponse.QrCode, nil
}

func (t *Client) GeneratePayUrl(ctx context.Context, params TradeParams) (string, error) {

	if params.TotalAmount <= 0 {
		return "", nil
	}

	queries := t.baseQueries()

	queries.Set("method", "alipay.trade.wap.pay")
	queries.Set("app_auth_token", params.AppAuthToken)
	queries.Set("notify_url", params.NotifyUrl)
	queries.Set("biz_content", conv.M2J(map[string]any{
		"out_trade_no":    params.OrderId,
		"total_amount":    params.TotalAmount,
		"time_expire":     params.TimeExpire,
		"subject":         params.Subject,
		"passback_params": conv.M2J(params.PassbackParams),
	}))

	queries.Set("sign", t.genSign(queries))

	return t.options.Entrypoint + "?" + queries.Encode(), nil
}

func (t *Client) RefreshAppAuthToken(ctx context.Context, refreshToken string) (*T, error) {
	queries := t.baseQueries()

	queries.Set("method", "alipay.open.auth.token.app")
	queries.Set("biz_content", conv.M2J(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}))

	queries.Set("sign", t.genSign(queries))

	response, err := resty.New().R().SetContext(ctx).Get(t.options.Entrypoint + "?" + queries.Encode())
	if err != nil {
		return nil, err
	}

	var resp T
	err = conv.B2S(response.Body(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
func (t *Client) GetAppAuthToken(ctx context.Context, authCode string) (*T, error) {

	queries := t.baseQueries()

	queries.Set("method", "alipay.open.auth.token.app")
	queries.Set("biz_content", conv.M2J(map[string]string{
		"grant_type": "authorization_code",
		"code":       authCode,
	}))

	queries.Set("sign", t.genSign(queries))

	response, err := resty.New().R().SetContext(ctx).Get(t.options.Entrypoint + "?" + queries.Encode())
	if err != nil {
		return nil, err
	}

	var resp T
	err = conv.B2S(response.Body(), &resp)
	if err != nil {
		return nil, err
	}

	if resp.AlipayOpenAuthTokenAppResponse.Code != "10000" || resp.AlipayOpenAuthTokenAppResponse.SubCode != "" {
		return nil, fmt.Errorf("%s %s", resp.AlipayOpenAuthTokenAppResponse.SubCode, resp.AlipayOpenAuthTokenAppResponse.SubMsg)
	}

	return &resp, nil
}

type T2 struct {
	AlipayTradePrecreateResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		OutTradeNo string `json:"out_trade_no"`
		QrCode     string `json:"qr_code"`
	} `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

type T struct {
	AlipayOpenAuthTokenAppResponse struct {
		Code    string      `json:"code"`
		Msg     string      `json:"msg"`
		SubCode string      `json:"sub_code"`
		SubMsg  string      `json:"sub_msg"`
		Tokens  []AuthToken `json:"tokens"`
	} `json:"alipay_open_auth_token_app_response"`
	Sign string `json:"sign"`
}

type AuthToken struct {
	AppAuthToken    string `json:"app_auth_token"`
	AppRefreshToken string `json:"app_refresh_token"`
	AuthAppId       string `json:"auth_app_id"`
	ExpiresIn       int    `json:"expires_in"`
	ReExpiresIn     int    `json:"re_expires_in"`
	UserId          string `json:"user_id"`
}
