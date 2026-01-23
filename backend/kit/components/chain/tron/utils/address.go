package tronutils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	trx_address "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"math/big"
)

func GenerateRandomAddress() (string, string) {
	pri, err := btcec.NewPrivateKey()
	if err != nil {
		return "", ""
	}
	if len(pri.Serialize()) != 32 {
		for {
			pri, err = btcec.NewPrivateKey()
			if err != nil {
				continue
			}
			if len(pri.Serialize()) == 32 {
				break
			}
		}
	}

	addr := trx_address.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
	wif := hex.EncodeToString(pri.Serialize())
	return addr, wif
}

func PrivateKeyToAddress(pk string) (string, error) {
	pkBytes, err := common.Hex2Bytes(pk)
	if err != nil {
		return "", err
	}
	privateKey, _ := btcec.PrivKeyFromBytes(pkBytes)
	addr := trx_address.PubkeyToAddress(privateKey.ToECDSA().PublicKey).String()
	return addr, nil
}

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ToHexAddress(address string) string {
	return hex.EncodeToString(base58Decode([]byte(address)))
}

func ToBase58Address(hexAddress string) string {
	addrByte, err := hex.DecodeString(hexAddress)
	if err != nil {
		return ""
	}

	sha := sha256.New()
	sha.Write(addrByte)
	shaStr := sha.Sum(nil)

	sha2 := sha256.New()
	sha2.Write(shaStr)
	shaStr2 := sha2.Sum(nil)

	addrByte = append(addrByte, shaStr2[:4]...)
	return string(base58Encode(addrByte))
}

func base58Encode(input []byte) []byte {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	reverseBytes(result)
	return result
}

func base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded[:len(decoded)-4]
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
