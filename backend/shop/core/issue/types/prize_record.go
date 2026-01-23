package types

// PrizeRecord 获奖记录
type PrizeRecord struct {
	Ticket any
	Grade  int
	Count  int
	Amount float64
}

type PrizeRecords []PrizeRecord

func (ts PrizeRecords) Total() (int, float64) {

	var totalAmount float64
	var totalCount int
	for _, t := range ts {
		totalAmount += t.Amount * float64(t.Count)
		totalCount += t.Count
	}

	return len(ts), totalAmount
}
