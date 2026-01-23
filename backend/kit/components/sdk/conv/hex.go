package conv

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
	"strings"
)

func HexToInt64(hex string) (int64, error) {
	val := strings.ReplaceAll(hex, "0x", "")
	return strconv.ParseInt(val, 16, 64)
}

func TrimLeftZeros(val string) string {

	if strings.HasPrefix(val, "0x") {
		val = val[2:]
	}

	var trimFrom int
	for i, v := range val {
		if v != '0' {
			trimFrom = i
			break
		}
	}

	return "0x" + val[trimFrom:]
}

func HexToBigInt(hex string) *big.Int {

	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}

	n := new(big.Int)
	if hex == "" {
		return n
	}

	n, _ = n.SetString(hex, 16)

	return n
}

func HexToBytes(hexStr string) ([]byte, error) {

	if !strings.HasPrefix(hexStr, "0x") {
		hexStr = "0x" + hexStr
	}

	decodeBytes, err := hexutil.Decode(hexStr)
	if err != nil {
		return nil, err
	}

	return decodeBytes, nil
}

func HexToBytes32(hexStr string) ([32]byte, error) {

	rsp := [32]byte{}

	bytes, err := HexToBytes(hexStr)
	if err != nil {
		return rsp, err
	}

	copy(rsp[:], bytes)

	return rsp, nil
}
