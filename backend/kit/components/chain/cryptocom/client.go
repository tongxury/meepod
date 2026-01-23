package cryptocom

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
	"strings"
)

func NewCryptoComClient() *Client {
	return &Client{
		url:    "https://min-api.cryptocompare.com",
		apiKey: "ad9c9191fbbeac998764a1e41e7c3cac2c06cc51ecd366a81cab4d245fa1ade6",
	}
}

type Client struct {
	url    string
	apiKey string
}

//func (c *Client) ()  {
////https://min-api.cryptocompare.com/data/top/totalvolfull?limit=10&tsym=USD
//}

var NotFound = errors.New("not found")

func (c *Client) GetDayAvgPrice(ctx context.Context, symbol string, date string) (float64, error) {

	ts := timed.DateTs(date)

	url := fmt.Sprintf("%s/data/dayAvg?fsym=%s&tsym=USD&toTs=%d&api_key=%s&extraParams=app",
		c.url, symbol, ts, c.apiKey)

	bodyBytes, code, err := httpcli.Client().GET(ctx, url)
	if err != nil {
		return 0, err
	}

	if code != 200 {
		return 0, fmt.Errorf("%s", string(bodyBytes))
	}

	var errRsp ErrResult

	err = conv.B2S(bodyBytes, &errRsp)
	if err != nil {
		return 0, err
	}

	if errRsp.Response == "Error" {
		if strings.Contains(errRsp.Message, "There is no data for") {
			return 0, NotFound
		}
		if strings.Contains(errRsp.Message, "CCCAGG conversion does not exist for this coin pair") {
			return 0, NotFound
		}

		return 0, fmt.Errorf(errRsp.Message)
	}

	var rsp Result
	err = conv.B2S(bodyBytes, &rsp)
	if err != nil {
		return 0, err
	}

	return rsp.USD, nil
}

type Result struct {
	USD            float64 `json:"USD"`
	ConversionType struct {
		Type             string `json:"type"`
		ConversionSymbol string `json:"conversionSymbol"`
	} `json:"ConversionType"`
}

type ErrResult struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data struct {
	} `json:"Data"`
	ParamWithError string `json:"ParamWithError"`
}
