package f3d

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
)

type Tickets []*Ticket

type Ticket struct {
	Cat string   `json:"cat"`
	Bai []string `json:"bai"`
	Shi []string `json:"shi"`
	Gen []string `json:"gen"`
	Ton []string `json:"ton"`
	// 只有开奖结果有这3个字段
	//Sales       string            `json:"sales,omitempty"`
	//PoolMoney   string            `json:"pool_money,omitempty"`
	//PrizeGrades types.PrizeGrades `json:"prize_grades,omitempty"`
}

func (t *Ticket) Amount() (int, float64) {

	var count int
	switch t.Cat {
	case "z1":
		count = len(t.Bai) * len(t.Shi) * len(t.Gen)
	case "z3":
		count = mathd.Cmn(len(t.Ton), 2) * 2
	case "z6":
		count = mathd.Cmn(len(t.Ton), 3)
	}

	return count, float64(count) * 2
}

func (t *Ticket) Split() Tickets {
	return Tickets{t}
}
