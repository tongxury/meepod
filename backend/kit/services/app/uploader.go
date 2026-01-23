package app

import (
	"context"
	"encoding/base64"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/tracing"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/kit/services/util/oss"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"strings"
	"sync"
)

func UploadBase64Route(bucket string) func(c *gin.Context) {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var body struct {
			Base64Resources []string
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			gind.BadRequest(c, err)
			return
		}

		var images []string

		wg := sync.WaitGroup{}
		lc := sync.Mutex{}

		for _, x := range body.Base64Resources {
			wg.Add(1)

			go func(ctx context.Context, b64 string) {
				defer helper.DeferFunc(func() {
					wg.Done()
				})

				parts := strings.Split(b64, ",")

				filename := parts[0]
				//t_ := parts[1]
				src := parts[1]

				fileBytes, err := base64.StdEncoding.DecodeString(src)
				if err != nil {
					slf.WithError(err).Errorw("DecodeString err")
					return
				}

				_, err = comp.SDK().AliOSS().Upload(ctx, bucket, filename, fileBytes)
				if err != nil {
					slf.WithError(err).Errorw("Upload ")
					return
				}

				//info, err := comp.SDK().OSS().PutObject(ctx, bucket, filename, bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{})
				//if err != nil {
				//	slf.WithError(err).Errorw("PutObject err")
				//	return
				//}
				//
				//url := info.Key

				lc.Lock()
				images = append(images, filename)
				lc.Unlock()

			}(ctx, x)

		}

		wg.Wait()

		if len(images) != len(body.Base64Resources) {
			gind.Error(c, fmt.Errorf("upload err"))
			return
		}

		var rsp []map[string]string

		for _, image := range images {
			rsp = append(rsp, map[string]string{
				"key": image,
				"url": oss.Resource(image),
			})
		}

		gind.OK(c, gin.H{"files": rsp})
	}
}

func UploaderRoute(bucket string) func(c *gin.Context) {

	return func(c *gin.Context) {

		ctx := c.Request.Context()
		form, err := c.MultipartForm()
		if err != nil {
			gind.BadRequest(c, err)
			return
		}

		formFiles := form.File["file"]

		var fileUrls []string

		wg := sync.WaitGroup{}
		lc := sync.Mutex{}

		for _, x := range formFiles {

			wg.Add(1)
			go func(ctx context.Context, mf *multipart.FileHeader) {
				defer helper.DeferFunc(func() {
					wg.Done()
				})

				file, err := mf.Open()
				if err != nil {
					return
				}

				fileBytes, err := io.ReadAll(file)
				if err != nil {
					return
				}

				//var suffix string
				//nameParts := strings.Split(mf.Filename, ".")
				//if len(nameParts) > 1 {
				//	suffix = "." + nameParts[len(nameParts)-1]
				//}

				url, err := comp.SDK().AliOSS().Upload(ctx, bucket, mf.Filename, fileBytes)
				if err != nil {
					slf.WithError(err).Errorw("Upload err")
					return
				}

				lc.Lock()
				fileUrls = append(fileUrls, url)
				lc.Unlock()

			}(tracing.PropagateContext(ctx), x)

		}

		wg.Wait()

		if len(fileUrls) != len(formFiles) {
			gind.Error(c, fmt.Errorf("upload err"))
			return
		}

		gind.OK(c, fileUrls)
	}
}
