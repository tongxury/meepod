package ssq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSsqOption_Prize(t *testing.T) {

	target := &Ticket{
		Red:  []string{"01", "02", "03", "04", "05", "06"},
		Blue: []string{"01"},
	}

	var cases = []struct {
		Case     *SingleTicket
		Target   *Ticket
		Expected int
	}{
		{Case: &SingleTicket{Red: []string{"11", "12", "13", "14", "15", "16"}, Blue: []string{"11"}}, Target: target, Expected: 0},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "05", "06"}, Blue: []string{"01"}}, Target: target, Expected: 1},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "05", "06"}, Blue: []string{"02"}}, Target: target, Expected: 2},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "05", "16"}, Blue: []string{"01"}}, Target: target, Expected: 3},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "05", "16"}, Blue: []string{"02"}}, Target: target, Expected: 4},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "15", "16"}, Blue: []string{"01"}}, Target: target, Expected: 4},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "04", "15", "16"}, Blue: []string{"11"}}, Target: target, Expected: 5},
		{Case: &SingleTicket{Red: []string{"01", "02", "03", "14", "15", "16"}, Blue: []string{"01"}}, Target: target, Expected: 5},
		{Case: &SingleTicket{Red: []string{"11", "12", "13", "14", "15", "16"}, Blue: []string{"01"}}, Target: target, Expected: 6},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Case.prize(c.Target))
	}

}

func TestSsqTicket_Amount(t *testing.T) {

	var cases = []struct {
		Case     *Ticket
		Expected int
	}{
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, DRed: []string{"11"}, Blue: []string{"01"}}, Expected: 6},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, DRed: []string{}, Blue: []string{"01"}}, Expected: 1},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06"}, DRed: []string{}, Blue: []string{"01", "02"}}, Expected: 2},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06", "07"}, DRed: []string{}, Blue: []string{"01"}}, Expected: 7},
		{Case: &Ticket{Red: []string{"01", "02", "03", "04", "05", "06", "07"}, DRed: []string{}, Blue: []string{"01", "02"}}, Expected: 14},
	}

	for _, c := range cases {
		count, _ := c.Case.Amount()
		assert.Equal(t, c.Expected, count)
	}
	for _, c := range cases {
		options := c.Case.SingleTickets()
		assert.Equal(t, c.Expected, len(options))
	}

}

func TestSsqTicket_Split(t *testing.T) {

	var cases = []struct {
		Case     *Ticket
		Expected int
	}{
		{Case: &Ticket{Red: []string{
			"01", "02", "03", "04", "05", "06", "07", "08",
			"01", "02", "03", "04", "05", "06", "07",
		}, DRed: []string{}, Blue: []string{"01"}}, Expected: 1},
		{Case: &Ticket{Red: []string{
			"01", "02", "03", "04", "05", "06", "07", "08",
			"01", "02", "03", "04", "05", "06", "07", "08",
		}, DRed: []string{}, Blue: []string{"01"}}, Expected: 1},
		{Case: &Ticket{Red: []string{
			"01", "02", "03", "04", "05", "06", "07", "08",
			"01", "02", "03", "04", "05", "06", "07", "08",
		}, DRed: []string{}, Blue: []string{"01", "02"}}, Expected: 2},
	}

	for _, c := range cases {
		tickets := c.Case.Split()
		assert.Equal(t, c.Expected, len(tickets))
	}
}
