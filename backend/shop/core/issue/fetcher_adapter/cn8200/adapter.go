package cn8200

import (
	"bytes"
	"fmt"
	"strings"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/PuerkitoBio/goquery"
)

type Adapter struct {
}

func (t *Adapter) FindX7cResultByIndex(index string) (*types.X7cResult, error) {

	url := fmt.Sprintf("https://www.8200.cn/kjh/qxc/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.X7cResult{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 7 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {
		result.Result = append(result.Result, strings.TrimSpace(selection.Text()))
	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	doc.Find(".kjTable tbody tr").EachWithBreak(func(idx int, selection *goquery.Selection) bool {

		var count string
		var amount string
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			if i == 2 {
				count = strings.TrimSpace(selection.Text())
				count = strings.ReplaceAll(count, " 注", "")
			}

			if i == 1 {
				amount = strings.TrimSpace(selection.Text())
				amount = strings.ReplaceAll(amount, "元", "")
				amount = strings.TrimSpace(amount)
				amount = strings.ReplaceAll(amount, ",", "")
			}

		})

		result.PrizeGrades = append(result.PrizeGrades, types.PrizeGrade{
			Grade:  idx + 1,
			Count:  count,
			Amount: amount,
		})

		return idx < 5
	})

	// 检查奖金结果
	if len(result.PrizeGrades) == 0 {
		return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
	}

	for _, grade := range result.PrizeGrades {
		if strings.Contains(grade.Amount, "-") || strings.Contains(grade.Count, "-") {
			return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
		}
	}

	return &result, nil
}

func (t *Adapter) FindF3dResultByIndex(index string) (*types.F3dResult, error) {

	url := fmt.Sprintf("https://www.8200.cn/kjh/3d/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.F3dResult{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 3 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {

		code := strings.TrimSpace(selection.Text())
		if !strings.Contains(code, "-") {
			result.Result = append(result.Result, strings.TrimSpace(selection.Text()))
		}

	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	if len(result.Result) != 3 {
		return nil, xerror.Wrapf("invalid result")
	}

	return &result, nil
}

func (t *Adapter) FindDltResultByIndex(index string) (*types.DltResult, error) {
	url := fmt.Sprintf("https://www.8200.cn/kjh/dlt/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.DltResult{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 7 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {

		if selection.HasClass("blue") {
			result.Blue = append(result.Blue, strings.TrimSpace(selection.Text()))
		} else {
			result.Red = append(result.Red, strings.TrimSpace(selection.Text()))
		}
	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	doc.Find(".kjTable tbody tr").EachWithBreak(func(idx int, selection *goquery.Selection) bool {

		var count string
		var amount string
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			if i == 2 {
				count = strings.TrimSpace(selection.Text())
				count = strings.ReplaceAll(count, " 注", "")
			}

			if i == 1 {
				amount = strings.TrimSpace(selection.Text())
				amount = strings.ReplaceAll(amount, "元", "")
				amount = strings.TrimSpace(amount)
				amount = strings.ReplaceAll(amount, ",", "")
			}

		})

		result.PrizeGrades = append(result.PrizeGrades, types.PrizeGrade{
			Grade:  idx + 1,
			Count:  count,
			Amount: amount,
		})

		return idx < 5
	})

	// 检查奖金结果
	if len(result.PrizeGrades) == 0 {
		return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
	}

	for _, grade := range result.PrizeGrades {
		if strings.Contains(grade.Amount, "-") || strings.Contains(grade.Count, "-") {
			return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
		}
	}

	return &result, nil
}

func (t *Adapter) FindSsqResultByIndex(index string) (*types.SsqResult, error) {
	url := fmt.Sprintf("https://www.8200.cn/kjh/ssq/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.SsqResult{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 7 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {
		if i == 6 {
			result.Blue = strings.TrimSpace(selection.Text())
		} else {
			result.Red = append(result.Red, strings.TrimSpace(selection.Text()))
		}
	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	doc.Find(".kjTable tbody tr").EachWithBreak(func(idx int, selection *goquery.Selection) bool {

		var count string
		var amount string
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			if i == 2 {
				count = strings.TrimSpace(selection.Text())
				count = strings.ReplaceAll(count, " 注", "")
			}

			if i == 1 {
				amount = strings.TrimSpace(selection.Text())
				amount = strings.ReplaceAll(amount, "元", "")
				amount = strings.TrimSpace(amount)
				amount = strings.ReplaceAll(amount, ",", "")
			}

		})

		result.PrizeGrades = append(result.PrizeGrades, types.PrizeGrade{
			Grade:  idx + 1,
			Count:  count,
			Amount: amount,
		})

		return idx < 5
	})

	// 检查奖金结果
	if len(result.PrizeGrades) == 0 {
		return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
	}

	for _, grade := range result.PrizeGrades {
		if strings.Contains(grade.Amount, "-") || strings.Contains(grade.Count, "-") {
			return nil, xerror.Wrapf("invalid result: %v", conv.S2J(result.PrizeGrades))
		}
	}

	return &result, nil
}

func (t *Adapter) FindPl3ResultByIndex(index string) (*types.Pl3Result, error) {

	url := fmt.Sprintf("https://www.8200.cn/kjh/p3/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.Pl3Result{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 3 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {
		result.Result = append(result.Result, strings.TrimSpace(selection.Text()))
	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	return &result, nil
}

func (t *Adapter) FindPl5ResultByIndex(index string) (*types.Pl5Result, error) {

	url := fmt.Sprintf("https://www.8200.cn/kjh/p5/%s.htm", index)

	repBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(repBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result := types.Pl5Result{}

	ballsSelector := doc.Find(".ballBox .ball")

	if ballsSelector.Length() != 5 {
		return nil, xerror.Wrapf("invalid result: %s", ballsSelector.Text())
	}

	ballsSelector.Each(func(i int, selection *goquery.Selection) {
		result.Result = append(result.Result, strings.TrimSpace(selection.Text()))
	})

	doc.Find(".mb-15").Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())

		if strings.Contains(text, "开奖时间") {
			result.BaseResult.PrizeDate = strings.ReplaceAll(text, "开奖时间：", "")
			result.BaseResult.PrizeDate = result.BaseResult.PrizeDate[:10]
		}
	})

	return &result, nil
}
