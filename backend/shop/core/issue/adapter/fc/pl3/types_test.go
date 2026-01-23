package pl3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	ta, _ := time.LoadLocation("Asia/Shanghai")

	fmt.Println(ta)
}

func TestF3dTicket_Prize(t *testing.T) {

	var cases = []struct {
		Case     *Ticket
		Target   *Ticket
		Expected int
	}{
		{Case: &Ticket{Cat: "z1", Bai: []string{"6"}, Shi: []string{"7"}, Gen: []string{"8"}},
			Target: &Ticket{Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"7"}}, Expected: 0},

		{Case: &Ticket{Cat: "z1", Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"7"}},
			Target: &Ticket{Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"7"}}, Expected: 1},

		{Case: &Ticket{Cat: "z3", Ton: []string{"5", "6", "7"}},
			Target: &Ticket{Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"7"}}, Expected: 0},

		{Case: &Ticket{Cat: "z3", Ton: []string{"5", "6", "7"}},
			Target: &Ticket{Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"6"}}, Expected: 1},

		{Case: &Ticket{Cat: "z3", Ton: []string{"5", "6", "7"}},
			Target: &Ticket{Bai: []string{"7"}, Shi: []string{"6"}, Gen: []string{"6"}}, Expected: 1},

		{Case: &Ticket{Cat: "z6", Ton: []string{"5", "6", "7"}},
			Target: &Ticket{Bai: []string{"5"}, Shi: []string{"6"}, Gen: []string{"7"}}, Expected: 1},

		{Case: &Ticket{Cat: "z6", Ton: []string{"5", "6", "7"}},
			Target: &Ticket{Bai: []string{"6"}, Shi: []string{"5"}, Gen: []string{"7"}}, Expected: 1},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, len(c.Case.prize(c.Target)))
	}

}
