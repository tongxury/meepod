package digest

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

type MD5_ struct {
}

func Md5() *MD5_ {
	return &MD5_{}
}

func (m *MD5_) SumHex(data []byte) (string, error) {
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str, nil
}
