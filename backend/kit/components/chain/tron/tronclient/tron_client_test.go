package tronclient

import (
	common "github.com/fbsobreira/gotron-sdk/pkg/common"
	"testing"
)

func TestName(t *testing.T) {

	src := "TF6StBpaKsJQRPfBRvTeQYQsh9MzT47qy5"
	bytes, err := common.DecodeCheck(src)
	if err != nil {
		panic(err)
	}

	val := common.Bytes2Hex(bytes)

	println(val)
}
