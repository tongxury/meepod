package timed

import (
	"time"
)

var (
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	YYYY_MM_DD          = "2006-01-02"
	YYYYMMDDHH          = "2006010215"

	LocAsiaShanghai, _ = time.LoadLocation("Asia/Shanghai")
)

func LastDate(date string) string {

	t, _ := time.Parse(YYYY_MM_DD, date)

	return t.AddDate(0, 0, -1).Format(YYYY_MM_DD)
}

func DateTs(date string) int64 {
	t, _ := time.Parse(YYYY_MM_DD, date)
	return t.Unix()
}

func DateEndTs(date time.Time) int64 {
	year, month, day := date.Date()
	d := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return d.Unix()
}

func DateStartTs(date time.Time) int64 {
	year, month, day := date.Date()
	d := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return d.Unix()
}

func DateStart(t time.Time) time.Time {
	year, month, day := t.Date()
	d := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return d
}

func DateEnd(t time.Time) time.Time {
	year, month, day := t.Date()
	d := time.Date(year, month, day, 23, 59, 59, 0, t.Location())
	return d
}
