package coingecko

import "gitee.com/meepo/backend/kit/components/sdk/helper/mathd"

type Price struct {
	Address string
	Ts      int64
	Price   float64
}

type Prices []*Price

func (ps Prices) Avg() float64 {

	var tmp []float64
	for _, p := range ps {
		tmp = append(tmp, p.Price)
	}

	avg := mathd.Avg(tmp...)

	return avg

}

type Ohlc struct {
	Id    string
	Time  int64
	Open  float64
	High  float64
	Low   float64
	Close float64
}

type Ohlcs []*Ohlc

type Coin struct {
	Id     string
	Symbol string
	Name   string
}

type Coins []*Coin
