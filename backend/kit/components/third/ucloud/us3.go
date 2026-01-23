package ucloud

import (
	"context"
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"io"
)

type IUS3Client interface {
	Upload(ctx context.Context, bucketName, fileName string, reader io.Reader, mimeType string) (url string, err error)
}

var sUS3Client *us3Client = nil

func GetUS3Client() IUS3Client {
	return sUS3Client
}

var mPublicKey, mPrivateKey, mBucketHost, mFileHost string

func InitUS3(publicKey, privateKey, bucketHost, fileHost string) (err error) {
	// 创建US3Client实例。

	mPublicKey = publicKey
	mPrivateKey = privateKey
	mBucketHost = bucketHost
	mFileHost = fileHost

	return err
}

type us3Client struct {
	//client *us3.Client
}

func (c *us3Client) Upload(ctx context.Context, bucketName, fileName string, reader io.Reader, mimeType string) (url string, err error) {

	conf := &ufsdk.Config{
		PublicKey:       mPublicKey,
		PrivateKey:      mPrivateKey,
		BucketHost:      mBucketHost,
		BucketName:      bucketName,
		FileHost:        mFileHost,
		VerifyUploadMD5: false,
	}
	req, err := ufsdk.NewFileRequest(conf, nil)
	if err != nil {
		return "", err
	}

	//err = req.IOPut(reader, fileName, mimeType)

	err = req.IOMutipartAsyncUpload(reader, fileName, mimeType)
	if err != nil {
		return "", err
	}

	url = req.GetPublicURL(fileName)
	return url, err
}
