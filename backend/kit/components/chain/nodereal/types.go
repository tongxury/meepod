package nodereal

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type AToken struct {
	Id       string
	Name     string
	Symbol   string
	Decimals string
	PriceETH float64 // warning 不同使用场景 此值可能为空
}

type ATokens []*AToken

func (ts ATokens) ToMap() map[string]AToken {
	rsp := make(map[string]AToken, len(ts))

	for _, t := range ts {
		rsp[t.Id] = *t
	}
	return rsp
}

type ASwap struct {
	Id          string
	Version     string
	Pair        *APair
	Timestamp   int64
	From        string
	Amount0     string
	Amount1     string
	AmountUSD   string
	BlockNumber string
	LogIndex    int
	Hash        string
}

type APair struct {
	Id                   string
	Token0               *AToken
	Token1               *AToken
	CreatedAtTimestamp   string
	CreatedAtBlockNumber string
	Version              string
	Swapped              bool
}

type APairs []*APair

func (p *APair) Swap() {

	t0 := getStability(p.Token0.Id)
	t1 := getStability(p.Token1.Id)

	if t0 > t1 {
		tmp := p.Token0
		p.Token0 = p.Token1
		p.Token1 = tmp
		p.Swapped = true
	}
}

// 按照token0为买卖币种，token1作为稳定币 方便后续处理
// 稳定系数
func getStability(token string) int {
	for i, x := range []string{} {
		if x == token {
			return i + 1
		}
	}
	return 0
}

type ASwaps []*ASwap

func (t ASwaps) Accounts() []string {
	slice := mapset.NewSet[string]()
	for _, v := range t {
		slice.Add(v.From)
	}

	return slice.ToSlice()
}

func (t ASwaps) Hashes() []string {
	slice := mapset.NewSet[string]()
	for _, v := range t {
		slice.Add(v.Hash)
	}

	return slice.ToSlice()
}
