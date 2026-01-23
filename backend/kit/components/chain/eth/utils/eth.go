package ethutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
	"strings"
)

func PrivateKeyToAddress(pk string) (*common.Address, error) {

	if strings.HasPrefix(pk, "0x") {
		pk = pk[2:]
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &address, nil
}

func GenerateRandomAddress() (string, string) {

	privateKey, _ := crypto.GenerateKey()

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	privateKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))

	return strings.ToLower(address.Hex()), privateKeyStr
}

func IsHexAddress(address string) bool {
	return common.IsHexAddress(address)
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(address interface{}) bool {
	var addr common.Address
	switch v := address.(type) {
	case string:
		addr = common.HexToAddress(v)
	case common.Address:
		addr = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := addr.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func TrimHex(addr string) string {

	if len(addr) != 66 { // with '0x'
		return addr
	}

	return "0x" + addr[len(addr)-40:]

	//0x000000000000000000000000839e71613f9aa06e5701cf6de63e303616b0dde3
}

func FormatHex(addr string) string {
	if addr == "" {
		return ""
	}
	return strings.ToLower(TrimHex(addr))
}

func BytesToAddress(b []byte) common.Address {
	return common.BytesToAddress(b)
}
