package aliyun

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewOSSClient(endpoint, accessKey, accessSecret string) (*OSSClient, error) {
	client, err := oss.New(endpoint, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &OSSClient{client: client}, nil
}

type OSSClient struct {
	client *oss.Client
}

func (c *OSSClient) Upload(ctx context.Context, bucketName, imageName string, imageBytes []byte) (string, error) {
	// 获取存储空间。
	bucket, err := c.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// 上传Byte数组。
	err = bucket.PutObject(imageName, bytes.NewReader(imageBytes))
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.oss-cn-beijing.aliyuncs.com/%s", bucketName, imageName)

	return url, nil
}
