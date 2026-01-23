package swap

type SushiPair struct {
	Id        string
	Token0    *AToken
	Token1    *AToken
	Timestamp string
	Block     string
	Version   string
	Swapped   bool
}

type SushiPairs []*SushiPair

func (ps SushiPairs) AsAPairs() APairs {

	var rsp APairs
	for _, x := range ps {

		y := APair{
			Id:                   x.Id,
			Token0:               x.Token0,
			Token1:               x.Token1,
			CreatedAtTimestamp:   x.Timestamp,
			CreatedAtBlockNumber: x.Block,
			Version:              "sushiswap",
		}

		rsp = append(rsp, &y)
	}

	return rsp
}

type SushiSwap RawV2Swap

type SushiSwaps RawV2Swaps
