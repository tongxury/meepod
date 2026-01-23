package encryptor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsa(t *testing.T) {

	pub := `
-----BEGIN rsa public key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYaR0KOVZJRaPuLl2XDCBkiGSk
Dtazv+JpgrnHIMdl5ThXyQpE0adTLVb1ud1M8ezPyZLNU+CA8FHwYhtoxSATOhPE
8SXnsM5wMupXPNXkvsShNJkZ1U3QcjF2ooMlvzVraaTb08XEWxHmH5xW1ghbRzW9
Gias1IqSGMuisjIPnQIDAQAB 
-----END rsa public key-----
	`

	pri := `
-----BEGIN rsa private key-----
MIICXQIBAAKBgQDYaR0KOVZJRaPuLl2XDCBkiGSkDtazv+JpgrnHIMdl5ThXyQpE
0adTLVb1ud1M8ezPyZLNU+CA8FHwYhtoxSATOhPE8SXnsM5wMupXPNXkvsShNJkZ
1U3QcjF2ooMlvzVraaTb08XEWxHmH5xW1ghbRzW9Gias1IqSGMuisjIPnQIDAQAB
AoGAZ0bioPpz/0vIy+Y8q9URsGiW/uRF+kpcltXYKvutrScTGHHNAMK9A6jjkyk8
P3hE93TPJkYdIeuObxWi1wEcKK+34FBsu560/AUm23ZuKCOm36z7Xb+yLpDxcWqJ
MMswhPvQqhmDqlMw9hj5WqkiAgJBzjETEZmmDIc4icNOS1ECQQD6tTvrdvoda/wq
ruW8X12iBS3nUsQSfcX1sLi61f1zRZicSCDwvsDC4ur21G2zHurbAaTUQqDFnf7+
Az5SaPjTAkEA3PqLcUsqpE595wJvyeuFsIw3c/MNiuRavbjwOkrwzEUUNyq4HmeZ
wMFSXWZYoYwzU4S7TGVZ6PKdJUipjoGPzwJBAK3AZSKvdnBlooJCbF29Cjt7s3Ca
X+Eg4c2BCMYUAG+fUEEfjBTNXvKyKX2fg9ecGdBmt0GUW7AZ69tHjC25KpkCQGg/
8uUB9x4IwbDoH2D9Mdb2b3rOIYdy77QtuXdmv28+76iPCMmfSpP7ICZcEFg2UkiG
h+4kqmQRgT2DqCpIyVUCQQCi7LuRGpEjFN+JcqfeV1CszazSTD0GoyxgHaH3pDla
Isy0Dz3nJit1F7jSFH+7/ik7tJPGMpIRa3ZRJiDY3x04
-----END rsa private key-----
	`

	txt := "text"

	cipherText, err1 := RSA().Encrypt([]byte(txt), pub)

	plainText, err2 := RSA().Decrypt(cipherText, pri)

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, txt, string(plainText))

	hex, err3 := RSA().EncryptHex(txt, pub)
	fmt.Println(hex)
	assert.Nil(t, err3)
	decrypt, err4 := RSA().DecryptHex(hex, pri)
	assert.Nil(t, err4)
	assert.Equal(t, decrypt, string(plainText))

}
