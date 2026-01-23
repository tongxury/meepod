package oss

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
)

func Resource(key string) string {

	if key == "" {
		return ""
	}

	//return comp.Flags().OSSOptions.Entrypoint + "/images/" + key
	return fmt.Sprintf("https://%s.oss-cn-beijing.aliyuncs.com/%s", comp.Flags().AliOSSOptions.Bucket, key)
}

func Resources(keys []string) []string {

	var rsp []string
	for _, key := range keys {
		if key == "" {
			continue
		}

		rsp = append(rsp, Resource(key))
	}

	return rsp
}
