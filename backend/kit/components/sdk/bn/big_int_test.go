package bn

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestBigInt_Divide(t *testing.T) {

}

func TestBigN_Val(t *testing.T) {

	cases := []struct {
		bg   *BN
		want float64
	}{
		{bg: BigNumber(big.NewInt(100), 6), want: 0.0001},
		{bg: BigNumber(big.NewInt(0), 0), want: 0},
		{bg: BigNumber(big.NewInt(0), 10), want: 0},
		{bg: BigNumber(big.NewInt(10), 0), want: 10},
		{bg: BigNumber(big.NewInt(100000000), 6), want: 100},
	}

	for _, c := range cases {
		assert.Equal(t, c.want, c.bg.Val())
	}

}
func TestBigN_FromStr(t *testing.T) {

	cases := []struct {
		bg      *BN
		wantVal float64
		wantRaw string
	}{
		{bg: FromStr("100", 6), wantVal: 100, wantRaw: "100000000"},
		{bg: FromStr("100", 0), wantVal: 100, wantRaw: "100"},
		{bg: FromStr("0.01", 3), wantVal: 0.01, wantRaw: "10"},
	}

	for _, c := range cases {
		assert.Equal(t, c.wantVal, c.bg.Val())
		assert.Equal(t, c.wantRaw, c.bg.RawVal().String())
	}

}
