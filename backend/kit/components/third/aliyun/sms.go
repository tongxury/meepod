package aliyun

import (
	"context"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strings"
)

func NewSMSClient(accessKey, accessSecret string) (*SMSClient, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKey, accessSecret)
	if err != nil {
		return nil, err
	}

	return &SMSClient{client: client}, nil
}

type SMSClient struct {
	client *dysmsapi.Client
	//sign   string
}

func (s *SMSClient) Send(ctx context.Context, phoneNos []string, code string, sign, templateCode string) (*dysmsapi.SendSmsResponse, error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = strings.Join(phoneNos, ",")
	request.SignName = sign
	request.TemplateCode = templateCode
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

	response, err := s.client.SendSms(request)

	if err != nil {
		return nil, err
	}
	if response.Code != "OK" {
		return nil, errors.New("aliyun.oss: " + response.Message)
	}

	return response, nil
}
