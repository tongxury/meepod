package types

import "gitee.com/meepo/backend/kit/components/sdk/conv"

// PrizeGrade 中奖明细
type PrizeGrade struct {
	ItemId string
	Grade  int
	Count  string
	Amount string
}

type PrizeGrades []PrizeGrade

func (ts PrizeGrades) FilterItemId(itemId string) PrizeGrades {

	var tmp PrizeGrades

	for _, x := range ts {
		if x.ItemId == itemId {
			tmp = append(tmp, x)
		}
	}

	return tmp
}

func (ts PrizeGrades) Find(grade int) (*PrizeGrade, bool) {

	for _, t := range ts {
		if t.Grade == grade {
			return &t, true
		}
	}

	return nil, false
}

func (ts PrizeGrades) GetGradeAmount(grade int) float64 {

	for _, t := range ts {
		if t.Grade == grade {
			return conv.Float64(t.Amount)
		}
	}

	return 0
}
