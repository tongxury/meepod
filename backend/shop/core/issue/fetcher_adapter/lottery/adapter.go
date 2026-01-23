package lottery

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"strings"
)

type Adapter struct {
}

func (a Adapter) FindX7cResultByIndex(index string) (*types.X7cResult, error) {
	result, err := client().FindByIndex("04", index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	parts := strings.Split(result.LotteryDrawResult, " ")

	if len(parts) != 7 {
		slf.Errorw("[Main] invalid result", slf.String("r", result.LotteryDrawResult))
		return nil, err
	}

	var prizeGrades types.PrizeGrades

	for _, x := range result.PrizeLevelList {

		if x.StakeAmount == "" || x.StakeAmount == "0" {
			continue
		}

		prizeGrades = append(prizeGrades, types.PrizeGrade{
			Grade:  conv.Int(x.Group) / 10,
			Count:  strings.ReplaceAll(x.StakeCount, ",", ""),
			Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
		})
	}

	return &types.X7cResult{
		Result: parts,
		BaseResult: types.BaseResult{
			PrizeDate: result.LotteryDrawTime,
		},
		PrizeGrades: prizeGrades,
	}, nil
}

func (a Adapter) FindF3dResultByIndex(index string) (*types.F3dResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindSsqResultByIndex(index string) (*types.SsqResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindDltResultByIndex(index string) (*types.DltResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindPl3ResultByIndex(index string) (*types.Pl3Result, error) {

	result, err := client().FindByIndex("35", index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	parts := strings.Split(result.LotteryDrawResult, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid result: %s", conv.S2J(result))
	}

	var prizeGrades types.PrizeGrades

	for _, x := range result.PrizeLevelList {

		if x.StakeAmount == "" || x.StakeAmount == "0" {
			continue
		}

		prizeGrades = append(prizeGrades, types.PrizeGrade{
			Grade:  conv.Int(x.Group) / 10,
			Count:  strings.ReplaceAll(x.StakeCount, ",", ""),
			Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
		})

	}
	//date, _ := time.ParseInLocation(time.DateOnly, result.LotteryDrawTime, timed.LocAsiaShanghai)
	//y, m, d := date.Date()
	//prizedAt := time.Date(y, m, d, 21, 15, 0, 0, timed.LocAsiaShanghai)

	return &types.Pl3Result{
		BaseResult: types.BaseResult{
			PrizeDate: result.LotteryDrawTime,
		},
		Result:      parts,
		PrizeGrades: prizeGrades,
	}, nil

}

func (a Adapter) FindPl5ResultByIndex(index string) (*types.Pl5Result, error) {
	result, err := client().FindByIndex("350133", index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	parts := strings.Split(result.LotteryDrawResult, " ")

	if len(parts) != 5 {
		return nil, fmt.Errorf("invalid result: %s", conv.S2J(result))
	}

	var prizeGrades types.PrizeGrades

	for _, x := range result.PrizeLevelList {

		if x.StakeAmount == "" || x.StakeAmount == "0" {
			continue
		}

		prizeGrades = append(prizeGrades, types.PrizeGrade{
			Grade:  conv.Int(x.Group) / 10,
			Count:  strings.ReplaceAll(x.StakeCount, ",", ""),
			Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
		})

	}
	//date, _ := time.ParseInLocation(time.DateOnly, result.LotteryDrawTime, timed.LocAsiaShanghai)
	//y, m, d := date.Date()
	//prizedAt := time.Date(y, m, d, 21, 15, 0, 0, timed.LocAsiaShanghai)

	return &types.Pl5Result{
		Result: parts,
		BaseResult: types.BaseResult{
			PrizeDate: result.LotteryDrawTime,
		},
		PrizeGrades: prizeGrades,
	}, nil
}
