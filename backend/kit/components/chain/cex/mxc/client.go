package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strings"
	"time"
)

func NewMxcClient() *MxcClient {
	return &MxcClient{
		//apiKey:    "mx0vglVBa3JIIO2T0h",
		//apiSecret: "ffafd73215d841a190279fba80e116aa",
		apiKey:    "mx0vgl8sTp1IKtZvoX",
		apiSecret: "5e63a22b77ba4491bbf2a1bb61aad0d1",
		url:       "https://api.mexc.com",
	}
}

type MxcClient struct {
	apiKey    string
	apiSecret string
	url       string
}

func (c *MxcClient) ListSymbols() {
	uri := c.url + "/api/v3/selfSymbols"

	response, _ := c.PrivateGet(uri, "")

	fmt.Println(response)
}

func (c *MxcClient) PrivateGet(urlStr string, jsonParams string) (interface{}, error) {
	var path string
	timestamp := time.Now().UnixNano() / 1e6
	fmt.Println(timestamp)
	if jsonParams == "" {
		message := fmt.Sprintf("timestamp=%d", timestamp)
		sign := c.ComputeHmac256(message, c.apiSecret)
		path = fmt.Sprintf("%s?timestamp=%d&signature=%s", urlStr, timestamp, sign)
		fmt.Println("message:", message)
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	} else {
		strParams := c.JsonToParamStr(jsonParams)
		message := fmt.Sprintf("%s&timestamp=%d", strParams, timestamp)
		sign := c.ComputeHmac256(message, c.apiSecret)
		path = fmt.Sprintf("%s?%s&timestamp=%d&signature=%s", urlStr, strParams, timestamp, sign)
		fmt.Println("message:", c.ParamsEncode(message))
		fmt.Println("sign:", sign)
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().SetHeaders(map[string]string{
		"X-MEXC-APIKEY": c.apiKey,
		"Content-Type":  "application/json",
	}).Get(path)

	if err != nil {
		return nil, err
	}

	// fmt.Println("Response Info:", resp)
	return resp, nil
}

// urlencode
func (c *MxcClient) ParamsEncode(paramStr string) string {
	return url.QueryEscape(paramStr)
}

// 加密
func (c *MxcClient) ComputeHmac256(Message string, sec_key string) string {
	key := []byte(sec_key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(Message))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *MxcClient) JsonToParamStr(jsonParams string) string {
	//转化json参数->参数字符串
	var paramsarr []string
	var arritem string
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonParams), &m)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for key, value := range m {

		arritem = fmt.Sprintf("%s=%s", key, value)
		paramsarr = append(paramsarr, arritem)
		i++
		if i > len(m) {
			break
		}
	}
	paramsstr := strings.Join(paramsarr, "&")
	return paramsstr
}
