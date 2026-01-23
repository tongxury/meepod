package ssq

import (
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
)

const (
	//双色球 每周二、四、日21:15开奖
	PrizeHour, PrizeMinute = 21, 15

	MinRedPerTicket  = 6
	MaxRedPerTicket  = 16
	MaxDRedPerTicket = 5
	MaxBluePerTicket = 16
	MinBluePerTicket = 1
	// MaxAmountPerTicket 拆票规则
	MaxAmountPerTicket = 20000 //单注最大金额,单位为元
	MaxMultiple        = 99    //单注最大倍数

	//双色球国庆节和春节内不开奖日期

)

var (
	ExcludeDates = []string{
		// 2023 春节 (1/19-1/28)
		"2023-01-19", // 四
		"2023-01-21", // 日
		"2023-01-24", // 二
		"2023-01-26", // 四
		// 2023 国庆 (10/1-10/4)
		"2023-10-01", // 日
		"2023-10-03", // 二
		// 2024 春节 (2/8-2/17)
		"2024-02-06", // 二
		"2024-02-08", // 四
		"2024-02-11", // 日
		"2024-02-13", // 二
		"2024-02-15", // 四
		// 2024 国庆 (10/1-10/4)
		"2024-10-01", // 二
		"2024-10-03", // 四
		// 2025 春节 (1/27-2/5)
		"2025-01-26", // 日
		"2025-01-28", // 二
		"2025-01-30", // 四
		"2025-02-02", // 日
		"2025-02-04", // 二
		// 2025 国庆 (10/1-10/4)
		"2025-10-01", // 三
		"2025-10-02", // 四
		"2025-10-04", // 日
		// 2026 春节 (2/14-2/23)
		"2026-02-15", // 日
		"2026-02-17", // 二
		"2026-02-19", // 四
		"2026-02-22", // 日
		// 2026 国庆 (10/1-10/4)
		"2026-10-01", // 四
		"2026-10-03", // 日
		// 2027 春节 (预测: 2/6春节, 休市1/30-2/8)
		"2027-02-02", // 二
		"2027-02-04", // 四
		"2027-02-07", // 日
		// 2027 国庆 (10/1-10/4)
		"2027-10-02", // 日
		"2027-10-03", // 二
		// 2028 春节 (预测: 1/26春节, 休市1/20-1/29)
		"2028-01-20", // 四
		"2028-01-23", // 日
		"2028-01-25", // 二
		"2028-01-27", // 四
		// 2028 国庆 (10/1-10/4)
		"2028-10-01", // 日
		"2028-10-03", // 二
	}
)

type Generator struct {
}

func (t *Generator) Generate(fromIndex string, fromTime time.Time) (string, time.Time, time.Time) {
	return t.next(fromIndex, fromTime)
}

func (t *Generator) next(fromIndex string, fromTime time.Time) (string, time.Time, time.Time) {

	//每周二、四、日21:15开奖
	th, tm := PrizeHour, PrizeMinute

	y, m, d := fromTime.Date()
	h, mm, _ := fromTime.Clock()
	week := fromTime.Weekday()

	var nextPrizeAt time.Time

	switch week {
	case 1, 3, 5, 6:
		nextPrizeAt = time.Date(y, m, d+1, th, tm, 0, 0, timed.LocAsiaShanghai)
	case 2, 0:
		if h < th && mm < tm {
			nextPrizeAt = time.Date(y, m, d, th, tm, 0, 0, timed.LocAsiaShanghai)
		} else {
			nextPrizeAt = time.Date(y, m, d+2, th, tm, 0, 0, timed.LocAsiaShanghai)
		}
	case 4:
		if h < th && mm < tm {
			nextPrizeAt = time.Date(y, m, d, th, tm, 0, 0, timed.LocAsiaShanghai)
		} else {
			nextPrizeAt = time.Date(y, m, d+3, th, tm, 0, 0, timed.LocAsiaShanghai)
		}
	}

	// 顺延
	if helper.InSlice(nextPrizeAt.Format(time.DateOnly), ExcludeDates) {
		return t.next(fromIndex, nextPrizeAt)
	}

	// 跨年
	issueIndex := fmt.Sprintf("%d", conv.Int(fromIndex)+1)
	if nextPrizeAt.Year() != fromTime.Year() {
		issueIndex = fmt.Sprintf("%d001", nextPrizeAt.Year())
	}

	return issueIndex, nextPrizeAt.Add(-2 * time.Hour), nextPrizeAt

}
