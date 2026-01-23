package mathd

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/types"
	mathRand "math/rand"
	"strconv"
	"time"
)

func RandNumber(start, end int) int {
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	return r.Intn(end-start+1) + start
}

func Max[T types.Number](ts ...T) T {

	if len(ts) == 0 {
		return 0
	}

	var rsp T
	for i, t := range ts {
		if i == 0 {
			rsp = t
			continue
		}

		if rsp < t {
			rsp = t
		}
	}

	return rsp
}

func Min[T types.Number](ts ...T) T {

	if len(ts) == 0 {
		return 0
	}

	var rsp T
	for i, t := range ts {
		if i == 0 {
			rsp = t
			continue
		}

		if rsp > t {
			rsp = t
		}
	}

	return rsp
}

func Sum[T types.Number](ts ...T) T {
	var rsp T
	for _, t := range ts {
		rsp += t
	}

	return rsp
}

func Avg[T types.Float](ts ...T) T {

	if len(ts) == 0 {
		return 0
	}

	sum := Sum(ts...)

	return sum / T(len(ts))
}

func Factorial[T types.Int](num T) T {
	if num > 0 {
		return num * Factorial(num-1)
	} else {
		return 1
	}
}

func Cmn(total, target int) int {

	if target > total {
		return 0
	}

	return Factorial(total) / (Factorial(total-target) * Factorial(target))
}

func ToFixed4(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

func ToFixed2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
