package types

import "time"

type PrizeResult struct {
	Issue   string
	Result  string
	PrizeAt time.Time
}

type Issue struct {
	Index   string
	StartAt time.Time
	CloseAt time.Time
	PrizeAt time.Time
	Prized  bool
}
