package helper

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"strings"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"github.com/go-resty/resty/v2"
	redisV9 "github.com/redis/go-redis/v9"
)

type Http struct {
}

func (t *Http) Get(url string, useProxy bool, headers ...map[string]string) ([]byte, error) {

	//ctx := context.Background()
	//ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	c := resty.New()

	c.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	if len(headers) > 0 {
		c.SetHeaders(headers[0])
	}

	if useProxy {

		//proxy, err := t.getProxyV2(ctx)
		//if err != nil {
		//	return nil, err
		//}
		//
		//slf.Debugw("using proxy:", slf.String("url", url), slf.String("proxy", proxy))
		//
		//c.SetProxy(proxy)

	}

	response, err := c.R().Get(url)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}

func (t *Http) getProxyV2(ctx context.Context) (string, error) {

	url := fmt.Sprintf("%s/get/", comp.Flags().GetStr("proxy.pool.url"))

	resp, err := resty.New().R().SetContext(ctx).EnableTrace().Get(url)

	if err != nil {
		return "", xerror.Wrap(err)
	}

	var response struct {
		Anonymous  string `json:"anonymous"`
		CheckCount int    `json:"check_count"`
		FailCount  int    `json:"fail_count"`
		Https      bool   `json:"https"`
		LastStatus bool   `json:"last_status"`
		LastTime   string `json:"last_time"`
		Proxy      string `json:"proxy"`
		Region     string `json:"region"`
		Source     string `json:"source"`
	}

	err = conv.J2S(resp.Body(), &response)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	return "http://" + response.Proxy, nil
}

func (t *Http) getProxy(ctx context.Context) (string, error) {

	redisKey := "proxies.list"

	result, err := comp.SDK().Redis().LPop(ctx, redisKey).Result()
	if err != nil && !errors.Is(err, redisV9.Nil) {
		return "", err
	}
	if result != "" {
		return "http://" + result, nil
	}

	url := "http://proxy.siyetian.com/apis_get.html?token=AesJWLNR1Y45keJdXTq10dORVT41keBhnT31STqFUeNpXQx0keFBjTUF1dOpXU10keFRjTqdGe.AM4IDMxUTN4YTM&limit=10&type=0&time=&split=1&split_text=&repeat=0&isp=0"
	resp, err := resty.New().R().EnableTrace().Get(url)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	proxies := strings.Split(string(resp.Body()), "<br />")

	if len(proxies) == 0 {
		return "", fmt.Errorf("resp format invalid: %s", string(resp.Body()))
	}

	if comp.SDK().Redis().RPush(ctx, redisKey, conv.AnySlice(proxies)[1:]...).Err() != nil {
		slf.WithError(err).Errorw("RPush err")
	}

	return "http://" + proxies[0], nil
}

//114.235.211.171:12923<br />125.107.200.18:14868<br />180.120.179.117:14338<br />180.105.247.73:11425<br />106.111.231.104:13917<br />110.90.137.13:14627<br />222.93.185.40:15912<br />120.41.210.148:11832<br />183.150.102.98:13529<br />115.208.33.34:16463
