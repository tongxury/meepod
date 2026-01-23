package z14

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
)

type Parser struct {
	Min int
}

func (t *Parser) Parse(src string) ([]any, []any, float64, error) {

	var tickets Tickets
	if err := conv.J2S([]byte(src), &tickets); err != nil {
		return nil, nil, 0, xerror.Wrap(err)
	}

	var splitTickets Tickets
	for _, x := range tickets {

		if len(x.Options) < t.Min {
			return nil, nil, 0, fmt.Errorf("len invalid : %s", src)
		}

		splitTickets = append(splitTickets, x.Split()...)
	}

	// amount
	var totalAmount float64
	for _, x := range tickets {

		_, x.Amount = x.GetAmount(t.Min)
		totalAmount += x.Amount
	}

	return conv.AnySlice(tickets), conv.AnySlice(splitTickets), totalAmount, nil
}
