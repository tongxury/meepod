package eth

import (
	ethutils "gitee.com/meepo/backend/kit/components/chain/eth/utils"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type EventLog types.Log

type EventLogs []EventLog

func (es EventLogs) FilterByAddress(addresses []string) EventLogs {
	var rsp EventLogs
	m := es.GroupByToken()
	for _, address := range addresses {
		for _, log := range m[address] {
			rsp = append(rsp, log)
		}
	}
	return rsp
}

func (es EventLogs) BlockNumbers() []int64 {

	rsp := mapset.NewSet[int64]()

	for _, e := range es {
		rsp.Add(int64(e.BlockNumber))
	}

	return rsp.ToSlice()
}

func (es EventLogs) RelatedAccounts() []string {

	rsp := mapset.NewSet[string]()

	for _, e := range es {
		rsp.Add(e.From())
		rsp.Add(e.To())
	}

	return rsp.ToSlice()
}

func (es EventLogs) GroupByToken() map[string]EventLogs {

	rsp := map[string]EventLogs{}

	for _, e := range es {
		addr := strings.ToLower(e.Address.Hex())
		rsp[addr] = append(rsp[addr], e)
	}

	return rsp
}

func (es EventLogs) Tokens() []string {
	rsp := mapset.NewSet[string]()
	for _, e := range es {
		rsp.Add(strings.ToLower(e.Address.Hex()))
	}

	return rsp.ToSlice()
}

type TransferEventLog = EventLog
type TransferEventLogs = EventLogs

func (t *TransferEventLog) TokenAddress() string {
	return strings.ToLower(t.Address.Hex())
}

func (t *TransferEventLog) From() string {
	if len(t.Topics) != 3 {
		return ""
	}
	from := ethutils.BytesToAddress(t.Topics[1].Bytes())
	fromAddress := strings.ToLower(from.Hex())
	return fromAddress
}
func (t *TransferEventLog) To() string {
	if len(t.Topics) != 3 {
		return ""
	}

	to := ethutils.BytesToAddress(t.Topics[2].Bytes())
	toAddress := strings.ToLower(to.Hex())
	return toAddress
}

func (t *TransferEventLog) Value() *big.Int {
	hexValue := hexutil.Encode(t.Data)
	hexValue = conv.TrimLeftZeros(hexValue)

	bigAmount := conv.HexToBigInt(hexValue)

	return bigAmount
}

type TokenMetadata struct {
	Address  string
	Symbol   string
	Name     string
	Decimals int
}

type TokenMetadatas []*TokenMetadata

func (ts TokenMetadatas) AddressMap() map[string]*TokenMetadata {
	rsp := map[string]*TokenMetadata{}

	for _, t := range ts {
		rsp[t.Address] = t
	}
	return rsp
}
