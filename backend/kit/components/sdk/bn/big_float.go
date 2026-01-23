package bn

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type BigFloat struct {
	raw      *big.Float
	decimals uint
}

func MustFloat(raw any, decimals ...uint) *BigFloat {
	bn, _ := NewFloat(raw, decimals...)
	return bn
}

func NewFloat(raw any, decimals ...uint) (*BigFloat, bool) {

	dec := uint(0)
	if len(decimals) > 0 {
		dec = decimals[0]
	}

	var bn *big.Float
	var ok bool

	switch t := raw.(type) {
	case string:
		bn, ok = new(big.Float).SetString(t)
	case int64:
		bn = new(big.Float).SetInt64(t)
		ok = true
	case int:
		bn = new(big.Float).SetInt64(int64(t))
		ok = true
	case *big.Float:
		bn = new(big.Float).Set(t)
		ok = true
	case big.Float:
		bn = new(big.Float).Set(&t)
		ok = true
	}

	if !ok {
		return nil, false
	}

	return &BigFloat{raw: bn, decimals: dec}, true
}

func (t *BigFloat) Value() *big.Float {

	d := new(big.Float).SetFloat64(math.Pow10(int(t.decimals)))
	z := new(big.Float).Quo(t.raw, d)

	return z
}

func (t *BigFloat) FixedFloat(fixed int) float64 {

	f, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", t.Value()), 64)

	return f
}
