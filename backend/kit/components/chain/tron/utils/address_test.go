package tronutils

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomAddress(t *testing.T) {

	pk, _ := btcec.PrivKeyFromBytes(common.Hex2Bytes(""))

	addr := address.PubkeyToAddress(pk.ToECDSA().PublicKey).String()
	println(addr)
}

func TestToHexAddress(t *testing.T) {
	addr := "TMpdYew1gk8rfq6G94Wv296RSnesnh8Xeu"

	hexaddr := ToHexAddress(addr)

	assert.Equal(t, "4182024132a61acb3c4fddb409eae209a78d8ff03e", hexaddr)

}

func TestToBase58Address(t *testing.T) {

	addr := "4182024132a61acb3c4fddb409eae209a78d8ff03e"

	base58 := ToBase58Address(addr)

	assert.Equal(t, "TMpdYew1gk8rfq6G94Wv296RSnesnh8Xeu", base58)

}
