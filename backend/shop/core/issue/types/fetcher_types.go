package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"time"
)

type BaseResult struct {
	PrizeDate string
}

func (t *BaseResult) PrizeAt(hh, mm int) time.Time {
	date, _ := time.ParseInLocation(time.DateOnly, t.PrizeDate, timed.LocAsiaShanghai)
	y, m, d := date.Date()
	return time.Date(y, m, d, hh, mm, 0, 0, timed.LocAsiaShanghai)
}

type Pl3Result struct {
	BaseResult
	Result      []string
	PrizeGrades PrizeGrades
}

type F3dResult struct {
	BaseResult
	Result      []string
	PrizeGrades PrizeGrades
}

type Pl5Result struct {
	BaseResult
	Result      []string
	PrizeGrades PrizeGrades
}

type DltResult struct {
	BaseResult
	Red         []string
	Blue        []string
	PrizeGrades PrizeGrades
}

type SsqResult struct {
	BaseResult
	Red         []string
	Blue        string
	PrizeGrades PrizeGrades
}

type X7cResult struct {
	BaseResult
	Result      []string
	PrizeGrades PrizeGrades
}
