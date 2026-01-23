package eth

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	mapset "github.com/deckarep/golang-set/v2"
	"math"
	"math/big"
	"strings"
	"time"
)

type GBlock struct {
	Number       int64
	TxCount      int64   `json:"transactionCount"`
	TimestampHex string  `json:"timestamp"`
	Trades       GTrades `json:"transactions"`
}

type GBlocks []*GBlock

func (bs GBlocks) AsMap() map[int64]*GBlock {
	rsp := make(map[int64]*GBlock, len(bs))

	for _, b := range bs {
		rsp[b.Number] = b
	}
	return rsp
}

type GTrade struct {
	Hash      string
	Index     int64
	InputData string
	From      *GAccount
	To        *GAccount
	ValueHex  string `json:"value"`
	Status    int64
	Logs      GLogs
	Block     *GBlock
	Contract  *GAccount `json:"createdContract"`
}

func (b *GTrade) IsERC20Trade(address string) bool {
	if b.To == nil {
		return false
	}
	return helper.EqualsIgnoreCase(b.To.Address, address)
}

func (b *GTrade) IsSuccess() bool {
	return b.Status == 1
}

//
//func (b *GTrade) IsNormal() bool {
//	// 发送eth交易
//	return b.ValueHex != "0x" &&
//		b.To != nil && b.To.Code == "0x" &&
//		b.From != nil && b.From.Code == "0x" && len(b.Logs) == 0
//}

func (b *GTrade) Value() float64 {

	if b.ValueHex == "0x" {
		return 0
	}

	hex := b.ValueHex
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}

	bigValue, _ := new(big.Int).SetString(hex, 16)

	realValue := new(big.Float).Quo(new(big.Float).SetInt(bigValue), big.NewFloat(math.Pow10(18)))
	v, _ := realValue.Float64()
	return v
}

func (b *GTrade) Time() time.Time {
	return b.Block.Time()
}

func (b *GBlock) Time() time.Time {
	rsp, _ := conv.HexToInt64(b.TimestampHex)
	return time.Unix(rsp, 0)
}

type GTrades []*GTrade

func (b *GTrade) AsGEthTrade() *GEthTrade {
	return &GEthTrade{
		From:  b.From,
		To:    b.To,
		Value: b.Value(),
		Block: b.Block,
	}
}
func (ts GTrades) AsGEthTrades() GEthTrades {
	rsp := make(GEthTrades, 0, len(ts))
	for _, t := range ts {
		rsp = append(rsp, t.AsGEthTrade())
	}

	return rsp
}

type GEthTrade struct {
	From  *GAccount
	To    *GAccount
	Value float64
	Block *GBlock
}

type GEthTrades []*GEthTrade

func (es GEthTrades) RelatedAccounts() []string {

	tmp := mapset.NewSet[string]()

	for _, e := range es {
		tmp.Add(e.From.Address)
		if e.To != nil {
			tmp.Add(e.To.Address)
		}
	}

	return tmp.ToSlice()
}

type GLog struct {
	Index   int64
	Topics  []string
	Data    string
	Account *GAccount
	Trade   *GTrade `json:"transaction"`
}

type GLogs []*GLog

type GAccount struct {
	Address           string
	CurrentTxCountHex string `json:"transactionCount"`
	BalanceHex        string `json:"balance"`
	Code              string
}

func (g *GAccount) CurrentTxCount() int64 {
	count, err := conv.HexToInt64(g.CurrentTxCountHex)
	if err != nil {
		slf.WithError(err).Errorw("HexToInt64 err", slf.Reflect("account", g))
	}
	return count
}

type GAccounts []*GAccount

func (g GLogs) AddressMap() map[string]*GLog {
	rsp := map[string]*GLog{}
	for _, log := range g {
		rsp[strings.ToLower(log.Account.Address)] = log
	}

	return rsp
}
