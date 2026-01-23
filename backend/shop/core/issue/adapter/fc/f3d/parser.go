package f3d

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
)

type Parser struct {
}

func (t *Parser) Parse(src string) ([]any, []any, float64, error) {
	var tickets Tickets
	if err := conv.J2S([]byte(src), &tickets); err != nil {
		return nil, nil, 0, xerror.Wrap(err)
	}

	var splitTickets Tickets
	for _, x := range tickets {

		switch x.Cat {
		case "z1":
			if len(x.Bai) == 0 || len(x.Shi) == 0 || len(x.Gen) == 0 {
				return nil, nil, 0, fmt.Errorf("invalid : %s", src)
			}
		case "z3":
			if len(x.Ton) < 2 {
				return nil, nil, 0, fmt.Errorf("invalid : %s", src)
			}
		case "z6":
			if len(x.Ton) < 3 {
				return nil, nil, 0, fmt.Errorf("invalid : %s", src)
			}
		}

		splitTickets = append(splitTickets, x.Split()...)
	}

	// amount
	var totalAmount float64
	for _, x := range tickets {
		_, amount := x.Amount()
		totalAmount += amount
	}

	return conv.AnySlice(tickets), conv.AnySlice(splitTickets), totalAmount, nil
}
