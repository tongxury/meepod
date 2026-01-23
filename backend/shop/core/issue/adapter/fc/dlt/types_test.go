package dlt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicket_Prize(t *testing.T) {

	target := &Ticket{
		Red:  []string{"01", "02", "03", "04", "05"},
		Blue: []string{"01", "02"},
	}

	var cases = []struct {
		Case          *Ticket
		Target        *Ticket
		ExpectedGrade int
		ExpectedCount int
	}{
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 1, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 1, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, Blue: []string{"01", "03"}}, Target: target, ExpectedGrade: 2, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05"}, Blue: []string{"04", "03"}}, Target: target, ExpectedGrade: 3, ExpectedCount: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, Blue: []string{"04", "03"}}, Target: target, ExpectedGrade: 3, ExpectedCount: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "15", "16"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 4, ExpectedCount: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16"}, DRed: []string{"15"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 4, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16", "15"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 5, ExpectedCount: 4},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16"}, DRed: []string{"15"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 5, ExpectedCount: 2},

		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "16"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 6, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "16", "15"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 6, ExpectedCount: 3},
		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "16"}, DRed: []string{"15"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 6, ExpectedCount: 2},

		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16"}, Blue: []string{"11", "12"}}, Target: target, ExpectedGrade: 7, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16", "17"}, Blue: []string{"11", "12"}}, Target: target, ExpectedGrade: 7, ExpectedCount: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "16", "17"}, Blue: []string{"11", "12", "13"}}, Target: target, ExpectedGrade: 7, ExpectedCount: 6},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "17"}, DRed: []string{"16"}, Blue: []string{"11", "12", "13"}}, Target: target, ExpectedGrade: 7, ExpectedCount: 3},

		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "17"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "17", "16"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 6},
		{Case: &Ticket{Red: []string{"02", "03", "14", "17"}, DRed: []string{"16", "01"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 4},
		{Case: &Ticket{Red: []string{"02", "03", "14", "17", "15"}, Blue: []string{"01", "02"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"02", "03", "14", "17", "15", "16"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 4},
		{Case: &Ticket{Red: []string{"02", "03", "14", "17"}, DRed: []string{"15", "16"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 8, ExpectedCount: 2},

		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "17"}, Blue: []string{"11", "12", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 3},
		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "17", "18"}, Blue: []string{"11", "12", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 9},
		{Case: &Ticket{Red: []string{"01", "02", "03", "14", "18"}, DRed: []string{"17"}, Blue: []string{"11", "12", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 6},
		{Case: &Ticket{Red: []string{"01", "12", "13", "14", "18"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 1},
		{Case: &Ticket{Red: []string{"01", "12", "13", "14", "18", "19"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 5},
		{Case: &Ticket{Red: []string{"01", "02", "13", "14", "18", "19"}, Blue: []string{"01", "12", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 8},
		{Case: &Ticket{Red: []string{"11", "12", "13", "14", "18", "19"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 6},
		{Case: &Ticket{Red: []string{"13", "14", "18", "19"}, DRed: []string{"11", "12"}, Blue: []string{"01", "02", "13"}}, Target: target, ExpectedGrade: 9, ExpectedCount: 4},
	}

	for i, c := range cases {
		records := c.Case.prize(c.Target)

		assert.Equal(t, c.ExpectedGrade > 0, len(records) > 0)

		if len(records) > 0 {
			assert.Equal(t, c.ExpectedCount, records[0].Count, i)
			assert.Equal(t, c.ExpectedGrade, records[0].Grade, i)
		}
	}

}

func TestTicket_Amount(t *testing.T) {

	var cases = []struct {
		Case     *Ticket
		Expected int
	}{
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05"}, DRed: []string{"11"}, Blue: []string{"01", "02"}}, Expected: 5},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05"}, DRed: []string{"11"}, Blue: []string{"01", "02", "03"}}, Expected: 15},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05"}, Blue: []string{"01", "02", "03"}}, Expected: 3},
	}

	for _, c := range cases {
		count, _ := c.Case.Amount()
		assert.Equal(t, c.Expected, count)
	}
	//for _, c := range cases {
	//	//options := c.Case.SingleTickets()
	//	//assert.Equal(t, c.Expected, len(options))
	//}

}
