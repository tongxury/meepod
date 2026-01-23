package z14

import (
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicket_PrizeSfc(t *testing.T) {

	var cases = []struct {
		Case          *Ticket
		Target        types.MatchTarget
		ExpectedGrade int
		ExpectedCount int
	}{
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"11": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"12": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"13": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"14": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedGrade: 1,
			ExpectedCount: 1,
		},
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "0"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"11": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"12": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"13": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"14": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedGrade: 2,
			ExpectedCount: 2,
		},
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "0"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"11": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"12": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"13": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"14": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedGrade: 0,
			ExpectedCount: 0,
		},
	}

	for _, c := range cases {
		records := c.Case.prizeSfc(c.Target)
		assert.Equal(t, c.ExpectedGrade > 0, len(records) > 0)
		if len(records) > 0 {
			assert.Equal(t, c.ExpectedGrade, records[0].Grade)
			assert.Equal(t, c.ExpectedCount, records[0].Count)
		}
	}
}

func TestTicket_PrizeRx9(t *testing.T) {

	var cases = []struct {
		Case          *Ticket
		Target        types.MatchTarget
		ExpectedCount int
	}{
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedCount: 1,
		},
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedCount: 10,
		},

		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedCount: 0,
		},

		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"9":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"10": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"11": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"12": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedCount: 7,
		},
		{
			Case: &Ticket{
				Options: map[string]Options{
					"1":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}, types.OddsItem{Result: "1"}}}},
					"2":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"3":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"4":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"5":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"6":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"7":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"8":  {Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "1"}}}},
					"9":  {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"10": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
					"11": {Dan: true, Odds: types.Odds{Items: types.OddsItems{types.OddsItem{Result: "3"}}}},
				},
			},
			Target: types.MatchTarget{
				Matches: types.Matches{
					{Id: "1", Result: types.Result{Value: "3"}},
					{Id: "2", Result: types.Result{Value: "3"}},
					{Id: "3", Result: types.Result{Value: "3"}},
					{Id: "4", Result: types.Result{Value: "3"}},
					{Id: "5", Result: types.Result{Value: "3"}},
					{Id: "6", Result: types.Result{Value: "3"}},
					{Id: "7", Result: types.Result{Value: "3"}},
					{Id: "8", Result: types.Result{Value: "3"}},
					{Id: "9", Result: types.Result{Value: "3"}},
					{Id: "10", Result: types.Result{Value: "3"}},
					{Id: "11", Result: types.Result{Value: "3"}},
					{Id: "12", Result: types.Result{Value: "3"}},
					{Id: "13", Result: types.Result{Value: "3"}},
					{Id: "14", Result: types.Result{Value: "3"}},
				},
			},
			ExpectedCount: 0,
		},
	}

	for _, c := range cases {
		records := c.Case.prizeRx9(c.Target)
		assert.Equal(t, c.ExpectedCount > 0, len(records) > 0)
		if len(records) > 0 {
			assert.Equal(t, c.ExpectedCount, records[0].Count)
		}
	}
}
