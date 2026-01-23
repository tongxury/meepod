package ucloud

import (
	"context"
	"github.com/ucloud/ucloud-sdk-go/services/usms"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func NewSMSClient(accessKey, accessSecret string) (*SMSClient, error) {

	cfg := ucloud.NewConfig()
	cfg.Region = "cn-bj2"

	credential := auth.NewCredential()
	credential.PublicKey = accessKey
	credential.PrivateKey = accessSecret

	smsClient := usms.NewClient(&cfg, &credential)

	return &SMSClient{client: smsClient}, nil
}

type SMSClient struct {
	client *usms.USMSClient
	//sign   string
}

func (t *SMSClient) Send(ctx context.Context, phone string, code string, sign, templateCode string) (*usms.SendUSMSMessageResponse, error) {

	message, err := t.client.SendUSMSMessage(&usms.SendUSMSMessageRequest{
		PhoneNumbers:   []string{phone},
		SigContent:     &sign,
		TemplateId:     &templateCode,
		TemplateParams: []string{code},
	})
	if err != nil {
		return nil, err
	}

	return message, nil
}
