package okooo

import (
	"bytes"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"strings"
	"time"
)

type Adapter struct {
}

func (a Adapter) ListZjcMatches(index string) (types.Matches, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) ListZjcMatchResults(index string) (types.Matches, types.PrizeGrades, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) ListZ14Matches(index string) (types.Matches, error) {
	// 默认获取最新场次
	url := fmt.Sprintf("https://www.okooo.com/zucai/ren9/")
	if index != "" {
		url += index
	}

	resultBytes, err := new(helper.Http).Get(url, false)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	reader := transform.NewReader(bytes.NewReader(resultBytes), simplifiedchinese.GBK.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if doc.Find(".zucaitop").Length() == 0 {
		return nil, xerror.Wrapf("invalid html")
	}

	var matches types.Matches

	var closeAt string
	doc.Find(".overTime em").Each(func(i int, closeAtEm *goquery.Selection) {
		closeAt = strings.TrimSpace(closeAtEm.Text())
	})

	var issue string
	doc.Find(".dqrm_hostroy").Each(func(i int, issueA *goquery.Selection) {
		if val, exists := issueA.Attr("rel"); exists {
			issue = val
		}
	})

	if issue == "" {
		return nil, fmt.Errorf("issue parse err")
	}

	doc.Find(".jcmaintable tr").Each(func(i int, row *goquery.Selection) {

		var league string
		var startAt string
		//var shortStartAt string
		var home string
		var homeTag string
		var guest string
		var guestTag string

		var oddsVictory float64
		var oddsDogfall float64
		var oddsDefeat float64

		//var resultValue string
		//var resultGoals string

		row.Find("td").Each(func(i int, part *goquery.Selection) {

			if i == 1 {
				row.Find(".jsLeagueName").Each(func(i int, x *goquery.Selection) {
					league = x.Text()
				})
			}
			if i == 2 {
				row.Find(".MatchTime").Each(func(i int, x *goquery.Selection) {
					startAt = strings.TrimSpace(x.Text())
				})
			}
			if i == 3 {
				row.Find("a").Each(func(i int, x *goquery.Selection) {

					if i == 1 {
						x.Find(".homename").Each(func(i int, homeSpan *goquery.Selection) {
							home = homeSpan.Text()
						})

						x.Find(".pltxt").Each(func(i int, oddsEm *goquery.Selection) {
							oddsVictory = conv.Float64(oddsEm.Text())
						})

						x.Find(".paim_em i").Each(func(i int, tag *goquery.Selection) {
							homeTag += tag.Text()
						})
					}

					if i == 2 {
						x.Find(".pltxt").Each(func(i int, oddsEm *goquery.Selection) {
							oddsDogfall = conv.Float64(oddsEm.Text())
						})
					}

					if i == 3 {
						x.Find(".awayname").Each(func(i int, guestSpan *goquery.Selection) {
							guest = guestSpan.Text()
						})

						x.Find(".pltxt").Each(func(i int, oddsEm *goquery.Selection) {
							oddsDefeat = conv.Float64(oddsEm.Text())
						})

						x.Find(".paim_em i").Each(func(i int, tag *goquery.Selection) {
							guestTag += tag.Text()
						})
					}
				})
			}
		})

		odds := types.Odds{
			Items: types.OddsItems{
				{Name: "主胜", Result: "3", Value: oddsVictory},
				{Name: "平局", Result: "1", Value: oddsDogfall},
				{Name: "客胜", Result: "0", Value: oddsDefeat},
			},
		}

		match := types.Match{
			League:       league,
			HomeTeam:     home,
			HomeTeamTag:  homeTag,
			GuestTeam:    guest,
			GuestTeamTag: guestTag,
			Category:     enum.MatchCategory_Z14.Value,
			Issue:        issue,
			CloseAt:      a.parseTime(closeAt),
			StartAt:      a.parseTime(startAt),
			Status:       enum.MatchStatus_Pending.Value,
			Odds:         odds,
		}

		matches = append(matches, &match)
	})

	return matches, nil
}

func (a Adapter) ListZ14MatchResults(index string) (types.Matches, types.PrizeGrades, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) parseTime(val string) time.Time {

	//expect "03-02 18:00"

	if len(val) > 11 {
		val = val[5:]
	}

	ti := strings.TrimSpace(val)
	// 跨年
	year := time.Now().Year()
	if strings.HasPrefix(ti, "01") {
		year += 1
	}

	dt := fmt.Sprintf("%d-%s:00", year, val)

	timed, _ := time.ParseInLocation(time.DateTime, dt, time.Local)

	return timed
}
