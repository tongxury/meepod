package okooo

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"strings"
)

func (t ZcMatchFetcher) FetchZ14Matches() {

	ctx := context.Background()

	matches, err := t.fetchZ14Matches(ctx, "", false)

	if err != nil {
		slf.WithError(err).Errorw("fetchZ14Matches err")
		return
	}

	err = matches.UpsertBatch(ctx)
	if err != nil {
		slf.WithError(err).Errorw("UpsertBatch err")
		return
	}
}

/*
*
尝试获取最新周期的场次。
只有成功获取到了最新数据 才会触发下一个周期。
*/
func (t ZcMatchFetcher) fetchLatestZ14Matches(ctx context.Context) (db.Matches, error) {
	return t.fetchZ14Matches(ctx, "", false)
}

func (t ZcMatchFetcher) fetchZ14Matches(ctx context.Context, index string, end bool) (db.Matches, error) {

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

	var matches db.Matches

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

		var resultValue string
		var resultGoals string

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

					if end {
						if i == 7 {
							resultGoals = strings.TrimSpace(x.Text())
						}
						if i == 8 {
							resultValue = strings.TrimSpace(x.Text())
						}
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

		match := db.Match{
			League:       league,
			HomeTeam:     home,
			HomeTeamTag:  homeTag,
			GuestTeam:    guest,
			GuestTeamTag: guestTag,
			Category:     enum.MatchCategory_Z14.Value,
			Issue:        issue,
			CloseAt:      t.parseTime(closeAt),
			StartAt:      t.parseTime(startAt),
			Status:       enum.MatchStatus_Pending.Value,
		}

		if resultGoals != "" && end {
			match.RealOdds = odds
			match.Result = types.Result{
				Value: resultValue,
				Goals: resultGoals,
			}
		} else {
			match.Odds = odds
		}

		matches = append(matches, &match)
	})

	return matches, nil
}
