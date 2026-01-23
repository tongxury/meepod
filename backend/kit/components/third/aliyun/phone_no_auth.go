package aliyun

import (
	"context"
	"errors"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func NewPhoneAuthClient(accessKey, accessSecret string) (*PhoneAuthClient, error) {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", accessKey, accessSecret)
	if err != nil {
		return nil, err
	}

	return &PhoneAuthClient{client: client}, nil
}

type PhoneAuthClient struct {
	client *sdk.Client
}

type PhoneResult struct {
	Message            string
	RequestId          string
	Code               string
	GetMobileResultDTO GetMobileResultDTO
}

type GetMobileResultDTO struct {
	Mobile string
}

func (p *PhoneAuthClient) GetPhoneNo(ctx context.Context, accessToken string) (string, error) {

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dypnsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "GetMobile"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["AccessToken"] = accessToken
	request.QueryParams["OutId"] = "1"

	response, err := p.client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}
	/**
	{
	    "Message": "请求成功",
	    "RequestId": 8906582,
	    "Code": "OK",
	    "GetMobileResultDTO": {
	        "Mobile": 121343241
	    }
	}
	*/

	var rsp PhoneResult
	err = conv.B2S(response.GetHttpContentBytes(), &rsp)
	if err != nil {
		slf.WithError(err).Errorw("response.GetHttpContentBytes() to json err", slf.Reflect("response", response))
		return "", err
	}

	if rsp.Code != "OK" {

		e := errors.New("response.GetHttpContentBytes() Code is not ok")
		slf.WithError(e).Errorw("", slf.Reflect("rsp", rsp))
		return "", e
	}

	return rsp.GetMobileResultDTO.Mobile, nil
}
