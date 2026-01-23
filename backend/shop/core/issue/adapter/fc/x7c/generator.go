package x7c

import (
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
)

var (
	// Z1RewardPerTicket = float64(1040)
	// Z3RewardPerTicket = float64(346)
	// Z6RewardPerTicket = float64(173)
	PrizeHour, PrizeMinute = 20, 30
	MaxAmountPerTicket     = float64(20000) //单注最大金额,单位为元
	MaxMultiple            = 99
)

var (
	// 国庆前4天
	// 春节前3天 到 初六
	X7cExcludeDates = []string{
		// 2023 春节 (1/19-1/28) - 周二、五、日
		"2023-01-20", // 五
		"2023-01-22", // 日
		"2023-01-24", // 二
		"2023-01-27", // 五
		// 2023 国庆 (10/1-10/4) - 周二、五、日
		"2023-10-01", // 日
		"2023-10-03", // 二
		// 2024 春节 (2/8-2/17) - 周二、五、日
		"2024-02-09", // 五
		"2024-02-11", // 日
		"2024-02-13", // 二
		"2024-02-16", // 五
		// 2024 国庆 (10/1-10/4) - 周二、五、日
		"2024-10-01", // 二
		"2024-10-04", // 五
		// 2025 春节 (1/27-2/5) - 周二、五、日
		"2025-01-28", // 二
		"2025-01-31", // 五
		"2025-02-02", // 日
		"2025-02-04", // 二
		// 2025 国庆 (10/1-10/4) - 周二、五、日
		"2025-10-03", // 五
		// 2026 春节 (2/14-2/23) - 周二、五、日
		"2026-02-15", // 日
		"2026-02-17", // 二
		"2026-02-20", // 五
		"2026-02-22", // 日
		// 2026 国庆 (10/1-10/4) - 周二、五、日
		"2026-10-02", // 五
		"2026-10-04", // 日
		// 2027 春节 (预测: 2/6春节, 休市1/30-2/8) - 周二、五、日
		"2027-01-30", // 六 -> 2027-01-31 日
		"2027-02-02", // 二
		"2027-02-05", // 五
		"2027-02-07", // 日
		// 2027 国庆 (10/1-10/4) - 周二、五、日
		"2027-10-01", // 五
		"2027-10-03", // 日
		// 2028 春节 (预测: 1/26春节, 休市1/20-1/29) - 周二、五、日
		"2028-01-21", // 五
		"2028-01-23", // 日
		"2028-01-25", // 二
		"2028-01-28", // 五
		// 2028 国庆 (10/1-10/4) - 周二、五、日
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
	//周二、周五、周日 20:30
	th, tm := PrizeHour, PrizeMinute

	y, m, d := fromTime.Date()
	h, mm, _ := fromTime.Clock()
	week := fromTime.Weekday()

	var prizeAt time.Time

	switch week {
	case 1, 3, 4, 6:
		prizeAt = time.Date(y, m, d+1, th, tm, 0, 0, timed.LocAsiaShanghai)
	case 5, 0:
		if h < th && mm < tm {
			prizeAt = time.Date(y, m, d, th, tm, 0, 0, timed.LocAsiaShanghai)
		} else {
			prizeAt = time.Date(y, m, d+2, th, tm, 0, 0, timed.LocAsiaShanghai)
		}
	case 2:
		if h < th && mm < tm {
			prizeAt = time.Date(y, m, d, th, tm, 0, 0, timed.LocAsiaShanghai)
		} else {
			prizeAt = time.Date(y, m, d+3, th, tm, 0, 0, timed.LocAsiaShanghai)
		}
	}

	if helper.InSlice(prizeAt.Format(time.DateOnly), X7cExcludeDates) {
		return t.next(fromIndex, prizeAt)
	}

	// 跨年
	issueIndex := fmt.Sprintf("%d", conv.Int(fromIndex)+1)
	if prizeAt.Year() != fromTime.Year() {
		issueIndex = fmt.Sprintf("%d001", prizeAt.Year())
	}

	return issueIndex, prizeAt.Add(-2 * time.Hour), prizeAt
}
