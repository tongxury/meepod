package dlt

import (
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
)

const (
	//每周一、三、六21:15开奖
	PrizeHour, PrizeMinute = 20, 25

	MinRedPerTicket  = 5
	MaxRedPerTicket  = 16
	MaxDRedPerTicket = 4
	MaxBluePerTicket = 16
	MinBluePerTicket = 2
	// MaxAmountPerTicket 拆票规则
	MaxAmountPerTicket = 20000 //单注最大金额,单位为元
	MaxMultiple        = 99    //单注最大倍数

	//双色球国庆节和春节内不开奖日期

)

var (
	ExcludeDates = []string{
		// 2023 春节 (1/19-1/28)
		"2023-01-21", // 六
		"2023-01-23", // 一
		"2023-01-25", // 三
		"2023-01-28", // 六
		// 2023 国庆 (10/1-10/4)
		"2023-10-02", // 一
		"2023-10-04", // 三
		// 2024 春节 (2/8-2/17)
		"2024-02-05", // 一
		"2024-02-07", // 三
		"2024-02-10", // 六
		"2024-02-12", // 一
		"2024-02-14", // 三
		"2024-02-17", // 六
		// 2024 国庆 (10/1-10/4)
		"2024-10-02", // 三
		// 2025 春节 (1/27-2/5)
		"2025-01-27", // 一
		"2025-01-29", // 三
		"2025-02-01", // 六
		"2025-02-03", // 一
		"2025-02-05", // 三
		// 2025 国庆 (10/1-10/4)
		"2025-10-01", // 三
		"2025-10-04", // 六
		// 2026 春节 (2/14-2/23)
		"2026-02-14", // 六
		"2026-02-16", // 一
		"2026-02-18", // 三
		"2026-02-21", // 六
		"2026-02-23", // 一
		// 2026 国庆 (10/1-10/4)
		"2026-10-03", // 六
		// 2027 春节 (预测: 2/6春节, 休市1/30-2/8)
		"2027-01-30", // 六
		"2027-02-01", // 一
		"2027-02-03", // 三
		"2027-02-06", // 六
		"2027-02-08", // 一
		// 2027 国庆 (10/1-10/4)
		"2027-10-02", // 六
		"2027-10-04", // 一
		// 2028 春节 (预测: 1/26春节, 休市1/20-1/29)
		"2028-01-22", // 六
		"2028-01-24", // 一
		"2028-01-26", // 三
		"2028-01-29", // 六
		// 2028 国庆 (10/1-10/4)
		"2028-10-02", // 一
		"2028-10-04", // 三
	}
)

type Generator struct {
}

func (t *Generator) Generate(fromIndex string, fromTime time.Time) (string, time.Time, time.Time) {
	return t.next(fromIndex, fromTime)
}

func (t *Generator) next(fromIndex string, fromTime time.Time) (string, time.Time, time.Time) {

	th, tm := PrizeHour, PrizeMinute

	y, m, d := fromTime.Date()
	h, mm, _ := fromTime.Clock()
	week := fromTime.Weekday()

	var nextPrizeAt time.Time

	switch week {
	case 2, 4, 5, 0:
		nextPrizeAt = time.Date(y, m, d+1, th, tm, 0, 0, timed.LocAsiaShanghai)
	case 1, 6:
		if h < th && mm < tm {
			nextPrizeAt = time.Date(y, m, d, th, tm, 0, 0, timed.LocAsiaShanghai)
		} else {
			nextPrizeAt = time.Date(y, m, d+2, th, tm, 0, 0, timed.LocAsiaShanghai)
		}
	case 3:
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
