package aliyun

import (
	"fmt"
	"testing"
)

func TestSMSClient_SendSMS(t *testing.T) {
	var code = "2322"
	println(fmt.Sprintf(`{"code":"%s"}`, code))
}
