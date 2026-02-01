package lottery

import (
	"fmt"
	"strings"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type Adapter struct {
}

// sportteryHeaders 返回访问 sporttery.cn 所需的请求头，避免被 EdgeOne 安全策略拦截
func (a Adapter) sportteryHeaders() map[string]string {
	return map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Referer":         "https://www.sporttery.cn/",
		"Origin":          "https://www.sporttery.cn",
	}
}

func (a Adapter) ListZjcMatches(index string) (types.Matches, error) {
	//url := "https://webapi.sporttery.cn/gateway/jc/football/getMatchListV1.qry?clientCode=3001"

	//url := "https://webapi.sporttery.cn/gateway/jc/football/getMatchCalculatorV1.qry?poolCode=&channel=c"
	url := "https://webapi.sporttery.cn/gateway/uniform/football/getMatchListV1.qry?clientCode=3001"

	resultBytes, err := new(helper.Http).Get(url, true, a.sportteryHeaders())

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var result ZcMatchesResponse
	err = conv.B2S[ZcMatchesResponse](resultBytes, &result)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if !result.Success || len(result.Value.MatchInfoList) == 0 {
		return nil, xerror.Wrapf("err result %s", string(resultBytes))
	}

	var matches types.Matches

	for _, xx := range result.Value.MatchInfoList {
		for _, x := range xx.SubMatchList {

			// 新 API 的 matchTime 格式是 HH:mm，需要补上秒数
			matchTimeStr := x.MatchTime
			if len(matchTimeStr) == 5 { // HH:mm 格式
				matchTimeStr = matchTimeStr + ":00"
			}
			startAt, _ := time.ParseInLocation(time.DateTime, x.MatchDate+" "+matchTimeStr, timed.LocAsiaShanghai)

			y := types.Match{
				League:       x.LeagueAllName,
				HomeTeam:     x.HomeTeamAllName,
				HomeTeamTag:  "",
				GuestTeam:    x.AwayTeamAllName,
				GuestTeamTag: "",
				Category:     enum.MatchCategory_Zjc.Value,
				Issue:        x.MatchDate,
				StartAt:      startAt,
				CloseAt:      startAt,
				Status:       enum.MatchStatus_UnStart.Value,
			}

			// 从 oddsList 中获取赔率
			var hadH, hadD, hadA string
			var hhadH, hhadD, hhadA, hhadGoalLine string
			for _, odds := range x.OddsList {
				switch odds.PoolCode {
				case "HAD":
					hadH, hadD, hadA = odds.H, odds.D, odds.A
				case "HHAD":
					hhadH, hhadD, hhadA, hhadGoalLine = odds.H, odds.D, odds.A, odds.GoalLine
				}
			}

			y.Odds.Items = append(y.Odds.Items,
				types.OddsItem{Name: "主胜", Result: "3", Value: conv.Float64(hadH)},
				types.OddsItem{Name: "平局", Result: "1", Value: conv.Float64(hadD)},
				types.OddsItem{Name: "客胜", Result: "0", Value: conv.Float64(hadA)},
			)

			if hhadGoalLine != "" && len(hhadGoalLine) > 1 {
				y.RCount = conv.Int(hhadGoalLine[1:])
			}
			y.Odds.RItems = append(y.Odds.RItems,
				types.OddsItem{Name: "让球主胜", Result: "3", Value: conv.Float64(hhadH)},
				types.OddsItem{Name: "让球平", Result: "1", Value: conv.Float64(hhadD)},
				types.OddsItem{Name: "让球客胜", Result: "0", Value: conv.Float64(hhadA)},
			)

			// 新 API 不再在此接口提供 HAFU/CRS/TTG 的详细赔率，这些字段留空

			matches = append(matches, &y)
		}

	}

	return matches, nil
}

func (a Adapter) ListZjcMatchResults(index string) (types.Matches, types.PrizeGrades, error) {

	url := "https://webapi.sporttery.cn/gateway/jc/football/getMatchResultV1.qry?matchPage=1&matchBeginDate=%s&matchEndDate=%s&leagueId=&pageSize=30&pageNo=1&isFix=0&pcOrWap=1"
	url = fmt.Sprintf(url, index, index)

	resultBytes, err := new(helper.Http).Get(url, true, a.sportteryHeaders())

	if err != nil {
		return nil, nil, xerror.Wrap(err)
	}

	var result ZcMatchResultsResponse
	err = conv.B2S[ZcMatchResultsResponse](resultBytes, &result)
	if err != nil {
		return nil, nil, xerror.Wrap(err)
	}

	if !result.Success {
		return nil, nil, xerror.Wrapf("err result %s", string(resultBytes))
	}

	if len(result.Value.MatchResult) == 0 {
		return nil, nil, nil
	}

	var matches types.Matches

	//issueInfo := result.Value

	//closeAt, err := time.ParseInLocation(time.DateTime, issueInfo.LotterySaleEndtime, timed.LocAsiaShanghai)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}

	for _, x := range result.Value.MatchResult {

		if x.SectionsNo999 == "" {
			continue
		}

		//startAt, err := time.ParseInLocation(time.DateOnly, x.StartTime, timed.LocAsiaShanghai)
		//if err != nil {
		//	return nil, xerror.Wrap(err)
		//}

		y := types.Match{
			//League:       x.MatchName,
			HomeTeam: strings.ReplaceAll(x.AllHomeTeam, " ", ""),
			//HomeTeamTag:  "",
			GuestTeam: strings.ReplaceAll(x.AllAwayTeam, " ", ""),
			//GuestTeamTag: "",
			Category: enum.MatchCategory_Zjc.Value,
			Issue:    index,
			//StartAt:      startAt,
			//CloseAt:      closeAt,
			Status: enum.MatchStatus_End.Value,
			//Odds: types.Odds{
			//	Items: types.OddsItems{
			//		{Name: "主队胜", Result: "3", Value: conv.Float64(x.H)},
			//		{Name: "平局", Result: "1", Value: conv.Float64(x.D)},
			//		{Name: "主队负", Result: "0", Value: conv.Float64(x.A)},
			//	},
			//},
			RealOdds: types.Odds{
				Items: types.OddsItems{
					{Name: "主胜", Result: x.H, Value: conv.Float64(x.H)},
					{Name: "平局", Result: x.D, Value: conv.Float64(x.D)},
					{Name: "客胜", Result: x.A, Value: conv.Float64(x.A)},
				},
			},
			Result: types.Result{
				Value:     a.goalsToValue(x.SectionsNo999),
				HalfValue: a.goalsToValue(x.SectionsNo1),
				Goals:     x.SectionsNo999,
				HalfGoals: x.SectionsNo1,
			},
		}

		matches = append(matches, &y)
	}

	return matches, nil, nil

}

func (a Adapter) goalsToValue(goals string) string {

	parts := strings.Split(goals, ":")
	if len(parts) != 2 {
		return ""
	}

	goalHome := conv.Int(parts[0])
	goalGuest := conv.Int(parts[1])

	if goalHome > goalGuest {
		return "3"
	} else if goalHome == goalGuest {
		return "1"
	} else {
		return "0"
	}

}

func (a Adapter) ListZ14Matches(index string) (types.Matches, error) {
	url := "https://webapi.sporttery.cn/gateway/lottery/getFootBallMatchV1.qry?param=90,0&lotteryDrawNum=%s&sellStatus=0&termLimits=10"
	//https://webapi.sporttery.cn/gateway/lottery/getFootBallMatchV1.qry?param=90,0&lotteryDrawNum=&sellStatus=1&termLimits=10

	url = fmt.Sprintf(url, index)

	resultBytes, err := new(helper.Http).Get(url, true, a.sportteryHeaders())

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var result Z14MatchesResponse
	err = conv.B2S[Z14MatchesResponse](resultBytes, &result)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if !result.Success {
		return nil, xerror.Wrapf("err result %s", string(resultBytes))
	}

	if len(result.Value.SfcMatch.MatchList) == 0 {
		return nil, nil
	}

	var matches types.Matches

	issueInfo := result.Value.SfcMatch

	closeAt, err := time.ParseInLocation(time.DateTime, issueInfo.LotterySaleEndtime, timed.LocAsiaShanghai)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, x := range result.Value.SfcMatch.MatchList {

		startAt, err := time.ParseInLocation(time.DateOnly, x.StartTime, timed.LocAsiaShanghai)
		if err != nil {
			return nil, xerror.Wrap(err)
		}

		y := types.Match{
			League:       x.MatchName,
			HomeTeam:     strings.ReplaceAll(x.MasterTeamAllName, " ", ""),
			HomeTeamTag:  "",
			GuestTeam:    strings.ReplaceAll(x.GuestTeamAllName, " ", ""),
			GuestTeamTag: "",
			Category:     enum.MatchCategory_Z14.Value,
			Issue:        issueInfo.LotteryDrawNum,
			StartAt:      startAt,
			CloseAt:      closeAt,
			Status:       enum.MatchStatus_Pending.Value,
			Odds: types.Odds{
				Items: types.OddsItems{
					{Name: "主队胜", Result: "3", Value: conv.Float64(x.H)},
					{Name: "平局", Result: "1", Value: conv.Float64(x.D)},
					{Name: "主队负", Result: "0", Value: conv.Float64(x.A)},
				},
			},
		}

		matches = append(matches, &y)
	}

	return matches, nil
}

func (a Adapter) ListZ14MatchResults(index string) (types.Matches, types.PrizeGrades, error) {
	//url := "https://webapi.sporttery.cn/gateway/lottery/getFootBallMatchV1.qry?param=90,0&lotteryDrawNum=&sellStatus=0&termLimits=10"
	url := "https://webapi.sporttery.cn/gateway/lottery/getFootBallDrawInfoByDrawNumV1.qry?isVerify=1&lotteryGameNum=90&lotteryDrawNum=%s"
	url = fmt.Sprintf(url, index)

	resultBytes, err := new(helper.Http).Get(url, true, a.sportteryHeaders())

	if err != nil {
		return nil, nil, xerror.Wrap(err)
	}

	var result Z14MatchResultsResponse
	err = conv.B2S[Z14MatchResultsResponse](resultBytes, &result)
	if err != nil {
		return nil, nil, xerror.Wrap(err)
	}

	if !result.Success {
		return nil, nil, xerror.Wrapf("err result %s", string(resultBytes))
	}

	if len(result.Value.MatchList) == 0 {
		return nil, nil, nil
	}

	var matches types.Matches

	for _, x := range result.Value.MatchList {

		if x.Result == "" {
			// 14个中任意一个没有结果都直接返回nil
			return nil, nil, nil
		}

		y := types.Match{
			//League:       x.MatchName,
			HomeTeam: strings.ReplaceAll(x.MasterTeamAllName, " ", ""),
			//HomeTeamTag:  "",
			GuestTeam: strings.ReplaceAll(x.GuestTeamAllName, " ", ""),
			//GuestTeamTag: "",
			Category: enum.MatchCategory_Z14.Value,
			Issue:    result.Value.LotteryDrawNum,
			//StartAt:      startAt,
			//CloseAt:      closeAt,
			Status: enum.MatchStatus_End.Value,
			//Odds: types.Odds{
			//	Items: types.OddsItems{
			//		{Name: "主队胜", Result: "3", Value: conv.Float64(x.H)},
			//		{Name: "平局", Result: "1", Value: conv.Float64(x.D)},
			//		{Name: "主队负", Result: "0", Value: conv.Float64(x.A)},
			//	},
			//},
			//RealOdds: types.Odds{},
			Result: types.Result{
				Value:     x.Result,
				HalfValue: "",
				Goals:     x.CzScore,
				HalfGoals: "",
			},
		}

		matches = append(matches, &y)
	}

	var prizeGrades types.PrizeGrades
	for _, x := range result.Value.PrizeLevelList {

		if x.PrizeLevel == "一等奖" {
			prizeGrades = append(prizeGrades, types.PrizeGrade{
				ItemId: enum.ItemId_sfc,
				Grade:  1,
				Count:  x.StakeCount,
				Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
			})
		} else if x.PrizeLevel == "二等奖" {
			prizeGrades = append(prizeGrades, types.PrizeGrade{
				ItemId: enum.ItemId_sfc,
				Grade:  2,
				Count:  x.StakeCount,
				Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
			})
		} else if x.PrizeLevel == "任选9场" {
			prizeGrades = append(prizeGrades, types.PrizeGrade{
				ItemId: enum.ItemId_rx9,
				Grade:  1,
				Count:  x.StakeCount,
				Amount: strings.ReplaceAll(x.StakeAmount, ",", ""),
			})
		}

	}

	if len(prizeGrades) == 0 {
		return nil, nil, nil
	}

	return matches, prizeGrades, nil
}
