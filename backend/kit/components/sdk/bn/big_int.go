package bn

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"math"
	"math/big"
)

type BigInt struct {
	raw      *big.Int
	decimals uint
}

func MustInt(raw any, decimals ...uint) *BigInt {
	bn, _ := NewInt(raw, decimals...)
	return bn
}

func NewInt(raw any, decimals ...uint) (*BigInt, bool) {

	dec := uint(0)
	if len(decimals) > 0 {
		dec = decimals[0]
	}

	var bn *big.Int
	var ok bool

	switch t := raw.(type) {
	case string:
		bn, ok = new(big.Int).SetString(t, 10)
	case int64:
		bn = new(big.Int).SetInt64(t)
		ok = true
	case int:
		bn = new(big.Int).SetInt64(int64(t))
		ok = true
	case *big.Int:
		bn = t
		ok = true
	case big.Int:
		bn = &t
		ok = true
	}

	if !ok {
		return nil, false
	}

	return &BigInt{raw: bn, decimals: dec}, true
}

func (t *BigInt) Divide(val *BigInt) *BigFloat {

	bf := new(big.Float).Quo(
		new(big.Float).Set(t.Value()),
		new(big.Float).Set(val.Value()),
	)

	ff, _ := NewFloat(bf)
	return ff
}

func (t *BigInt) Raw() *big.Int {
	return t.raw
}

func (t *BigInt) Value() *big.Float {

	raw := new(big.Float).SetInt(t.raw)
	d := new(big.Float).SetFloat64(math.Pow10(int(t.decimals)))

	z := new(big.Float).Quo(raw, d)
	//bf, _ := z.Float64()

	//bf := big.NewFloat(t.val * math.Pow10(int(t.decimals)))
	//
	//bi, _ := bf.Int(nil)
	return z
}

type BN struct {
	raw      *big.Int
	decimals uint
}

// Deprecated
func FromStr(val string, decimals ...uint) *BN {

	x, _ := new(big.Float).SetString(val)
	var dec uint
	if len(decimals) > 0 {
		dec = decimals[0]
	}

	y := big.NewFloat(math.Pow10(int(dec)))

	raw := new(big.Float).Mul(x, y)
	rawInt, _ := raw.Int(new(big.Int))

	return &BN{
		raw:      rawInt,
		decimals: dec,
	}
}

// Deprecated
func FromRaw(raw *big.Int, decimals uint) *BN {
	return &BN{raw: raw, decimals: decimals}
}

// Deprecated
func BigNumber(raw *big.Int, decimals uint) *BN {
	return &BN{raw: raw, decimals: decimals}
}

func (b *BN) RawVal() *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	return b.raw
}

func (b *BN) StrVal() string {
	if b == nil {
		return "0"
	}
	return conv.String(b.Val())
}

func (b *BN) Val() float64 {
	if b == nil {
		return 0
	}

	if b.raw == nil {
		return 0
	}

	bigD := math.Pow10(int(b.decimals))

	x, _ := new(big.Float).SetString(b.raw.String())
	y := new(big.Float).SetFloat64(bigD)

	z := new(big.Float).Quo(x, y)
	rsp, _ := z.Float64()
	return rsp
}
