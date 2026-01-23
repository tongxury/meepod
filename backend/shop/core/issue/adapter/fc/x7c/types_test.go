package x7c

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestX7cTicket_Prize(t *testing.T) {

	grades := types.PrizeGrades{
		{Count: "1", Grade: 1},
		{Count: "1", Grade: 2},
		{Count: "1", Grade: 3},
		{Count: "1", Grade: 4},
		{Count: "1", Grade: 5},
		{Count: "1", Grade: 6},
	}

	var cases = []struct {
		Case          *Ticket
		Target        *Ticket
		ExpectedCount int
		ExpectedGrade int
	}{
		{
			Case: &Ticket{
				Swan: []string{"0", "1", "2"},
				Wan:  []string{"0"},
				Qian: []string{"0"},
				Bai:  []string{"0"},
				Shi:  []string{"0"},
				Gen:  []string{"0"},
				Last: []string{"0"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 1,
			ExpectedGrade: 1,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"0"},
				Qian: []string{"0"},
				Bai:  []string{"0"},
				Shi:  []string{"0"},
				Gen:  []string{"0", "1", "2"},
				Last: []string{"1", "2", "3"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 3,
			ExpectedGrade: 2,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"0"},
				Qian: []string{"0"},
				Bai:  []string{"0"},
				Shi:  []string{"0", "1"},
				Gen:  []string{"1", "2", "3"},
				Last: []string{"0", "1"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 3,
			ExpectedGrade: 3,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"0"},
				Qian: []string{"0"},
				Bai:  []string{"0"},
				Shi:  []string{"0", "1"},
				Gen:  []string{"1", "2", "3"},
				Last: []string{"1", "2"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 6,
			ExpectedGrade: 4,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"0"},
				Qian: []string{"0"},
				Bai:  []string{"0", "1"},
				Shi:  []string{"1", "2"},
				Gen:  []string{"1", "2"},
				Last: []string{"1", "2", "3"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 12,
			ExpectedGrade: 5,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"0"},
				Qian: []string{"0", "1"},
				Bai:  []string{"1", "2"},
				Shi:  []string{"1", "2"},
				Gen:  []string{"1", "2"},
				Last: []string{"1", "2"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 16,
			ExpectedGrade: 6,
		},
		{
			Case: &Ticket{
				Swan: []string{"0", "1"},
				Wan:  []string{"1", "2"},
				Qian: []string{"1", "2"},
				Bai:  []string{"1"},
				Shi:  []string{"1"},
				Gen:  []string{"1", "2"},
				Last: []string{"0", "2"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 8,
			ExpectedGrade: 6,
		},

		{
			Case: &Ticket{
				Swan: []string{"1", "2"},
				Wan:  []string{"1", "2"},
				Qian: []string{"1"},
				Bai:  []string{"1"},
				Shi:  []string{"1"},
				Gen:  []string{"1"},
				Last: []string{"0"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 4,
			ExpectedGrade: 6,
		},
		{
			Case: &Ticket{
				Swan: []string{"0"},
				Wan:  []string{"1", "2"},
				Qian: []string{"1", "2"},
				Bai:  []string{"1", "2", "2"},
				Shi:  []string{"1"},
				Gen:  []string{"1"},
				Last: []string{"0"},
			},
			Target: &Ticket{Swan: []string{"0"},
				Wan:         []string{"0"},
				Qian:        []string{"0"},
				Bai:         []string{"0"},
				Shi:         []string{"0"},
				Gen:         []string{"0"},
				Last:        []string{"0"},
				PrizeGrades: grades,
			},
			ExpectedCount: 12,
			ExpectedGrade: 6,
		},
	}

	for _, c := range cases {
		rsp := c.Case.prize(c.Target)
		assert.Equal(t, c.ExpectedGrade, helper.Choose(len(rsp) > 0, rsp[0].Grade, 0))
		assert.Equal(t, c.ExpectedCount, helper.Choose(len(rsp) > 0, rsp[0].Count, 0))
	}

}
