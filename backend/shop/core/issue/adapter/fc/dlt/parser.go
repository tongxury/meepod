package dlt

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

	// 先根据倍数拆分
	//if multiple > consts.SsqMaxMultiple {
	//
	//	times := multiple / consts.SsqMaxMultiple
	//	left := multiple % consts.SsqMaxMultiple
	//
	//}

	// 再根据金额拆分
	var splitTickets Tickets

	for _, x := range tickets {
		rl := len(x.Red)
		bl := len(x.Blue)
		drl := len(x.DRed)

		if bl < MinBluePerTicket {
			return nil, nil, 0, fmt.Errorf("blue invalid : %s", src)
		}

		if rl+drl < MinRedPerTicket {
			return nil, nil, 0, fmt.Errorf("red invalid : %s", src)
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
