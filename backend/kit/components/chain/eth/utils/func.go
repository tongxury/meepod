package ethutils

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func GetFuncSelector(f string) string {

	f = strings.ReplaceAll(f, " ", "")

	funcBytes := crypto.Keccak256Hash([]byte(f)).Bytes()
	fStr := "0x" + hex.EncodeToString(funcBytes[:4])

	return fStr
}
