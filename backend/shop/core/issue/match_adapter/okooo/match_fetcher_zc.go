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
	"strings"
	"time"
)

type ZcMatchFetcher struct {
}

func NewZcMatchFetcher() *ZcMatchFetcher {
	return &ZcMatchFetcher{}
}

func (t ZcMatchFetcher) FetchZcMatches() {

	ctx := context.Background()

	var matches db.Matches
	//
	matchesWithOdds, err := t.fetchZcMatchesWithOdds(ctx, "")
	if err != nil {
		slf.WithError(err).Errorw("fetchZcMatchesWithOdds err")
	}
	matches = append(matches, matchesWithOdds...)
	//
	matchesWithGoalsOdds, err := t.fetchZcMatchesWithGoalsOdds(ctx, "")
	if err != nil {
		slf.WithError(err).Errorw("fetchZcMatchesWithGoalsOdds err")
	}
	matches = append(matches, matchesWithGoalsOdds...)
	//
	matchesWithScoreOdds, err := t.fetchZcMatchesWithScoreOdds(ctx, "")
	if err != nil {
		slf.WithError(err).Errorw("fetchZcMatchesWithScoreOdds err")
	}
	matches = append(matches, matchesWithScoreOdds...)
	//
	matchesWithHalfFullOdds, err := t.fetchZcMatchesWithHalfFullOdds(ctx, "")
	if err != nil {
		slf.WithError(err).Errorw("fetchZcMatchesWithHalfFullOdds err")
	}
	matches = append(matches, matchesWithHalfFullOdds...)
	//
	for _, match := range matches {
		err := match.UpsertOdds(ctx)
		if err != nil {
			slf.WithError(err).Errorw("UpsertOdds err")
		}
	}
}

// 总比分
func (t ZcMatchFetcher) fetchZcMatchesWithScoreOdds(ctx context.Context, date string) (db.Matches, error) {
	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}
	url := fmt.Sprintf("https://www.310win.com/buy/jingcai.aspx?typeID=102&oddstype=2&date=%s", date)

	resultBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resultBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	// 为了通过id对应数据
	matchesTap := map[string]*db.Match{}

	doc.Find("table tbody tr").Each(func(idx int, selection *goquery.Selection) {

		if matchId, found := selection.Attr("matchid"); found {
			match := db.Match{
				Category: enum.MatchCategory_Zjc.Value,
				Status:   enum.MatchStatus_UnStart.Value,
			}

			match.League, _ = selection.Attr("gamename")

			selection.Find("td").Each(func(i int, selection *goquery.Selection) {

				if i == 2 {
					match.StartAt = t.parseTime(selection.Text())
				}

				//if i == 3 {
				//	title, _ := selection.Attr("title")
				//	parts := strings.Split(title, "：")
				//
				//	if len(parts) == 2 {
				//		match.CloseAt = t.parseTime(parts[1][5:])
				//	}
				//}

				if i == 4 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.HomeTeam = strings.TrimSpace(selection.Text())
					})
				}
				if i == 6 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.GuestTeam = strings.TrimSpace(selection.Text())
					})
				}
			})

			matchesTap[matchId] = &match
		}

		if id, ex := selection.Attr("id"); ex && strings.HasPrefix(id, "tr_") {
			matchId := strings.ReplaceAll(id, "tr_", "")

			//matchesTap[matchId].

			selection.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {

				if i == 0 {
					selection.Find("td").Each(func(i int, selection *goquery.Selection) {
						if i == 0 {
							return
						}
						item := types.OddsItem{}

						selection.Find("b").Each(func(i int, selection *goquery.Selection) {
							item.Name = strings.TrimSpace(selection.Text())
							item.Result = strings.TrimSpace(selection.Text())
						})

						selection.Find("span").Each(func(i int, selection *goquery.Selection) {
							item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
						})

						matchesTap[matchId].Odds.ScoreVictoryItems = append(matchesTap[matchId].Odds.ScoreVictoryItems, item)

					})
				}

				if i == 1 {
					selection.Find("td").Each(func(i int, selection *goquery.Selection) {
						if i == 0 {
							return
						}

						if _, f := selection.Attr("id"); f {

							item := types.OddsItem{}
							selection.Find("b").Each(func(i int, selection *goquery.Selection) {
								item.Name = strings.TrimSpace(selection.Text())
								item.Result = strings.TrimSpace(selection.Text())
							})

							selection.Find("span").Each(func(i int, selection *goquery.Selection) {
								item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
							})

							matchesTap[matchId].Odds.ScoreDogfallItems = append(matchesTap[matchId].Odds.ScoreDogfallItems, item)
						}
					})
				}

				if i == 2 {
					selection.Find("td").Each(func(i int, selection *goquery.Selection) {
						if i == 0 {
							return
						}
						item := types.OddsItem{}
						selection.Find("b").Each(func(i int, selection *goquery.Selection) {
							item.Name = strings.TrimSpace(selection.Text())
							item.Result = strings.TrimSpace(selection.Text())
						})

						selection.Find("span").Each(func(i int, selection *goquery.Selection) {
							item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
						})

						matchesTap[matchId].Odds.ScoreDefeatItems = append(matchesTap[matchId].Odds.ScoreDefeatItems, item)

					})
				}
			})
		}

	})

	var matches db.Matches
	for _, match := range matchesTap {
		matches = append(matches, match)
	}

	return matches, nil
}

// 半全场
func (t ZcMatchFetcher) fetchZcMatchesWithHalfFullOdds(ctx context.Context, date string) (db.Matches, error) {
	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}
	url := fmt.Sprintf("https://www.310win.com/buy/jingcai.aspx?typeID=104&oddstype=2&date=%s", date)

	resultBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resultBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var matches db.Matches

	doc.Find("table tbody tr").Each(func(idx int, selection *goquery.Selection) {

		if _, found := selection.Attr("matchid"); found {

			match := db.Match{
				Category: enum.MatchCategory_Zjc.Value,
				Status:   enum.MatchStatus_UnStart.Value,
			}

			match.League, _ = selection.Attr("gamename")

			selection.Find("td").Each(func(i int, selection *goquery.Selection) {

				if i == 2 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.StartAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 3 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.CloseAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 4 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.HomeTeam = strings.TrimSpace(selection.Text())
					})
				}
				if i == 6 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.GuestTeam = strings.TrimSpace(selection.Text())
					})
				}

				if id, ex := selection.Attr("id"); ex && strings.HasPrefix(id, "cell_") {

					name, _ := selection.Attr("title")

					var value float64
					selection.Find("span").Each(func(i int, selection *goquery.Selection) {
						value = conv.Float64(strings.TrimSpace(selection.Text()))
					})

					match.Odds.HalfFullItems = append(match.Odds.HalfFullItems, types.OddsItem{
						Name:   name,
						Result: strings.ReplaceAll(name, "球", ""),
						Value:  value,
					})
				}
			})

			matches = append(matches, &match)

		}

	})

	return matches, nil
}

// 进球数
func (t ZcMatchFetcher) fetchZcMatchesWithGoalsOdds(ctx context.Context, date string) (db.Matches, error) {
	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}
	url := fmt.Sprintf("https://www.310win.com/buy/jingcai.aspx?typeID=103&oddstype=2&date=%s", date)

	resultBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	//utf8Body, err := iconv.NewReader(bytes.NewReader(resultBytes))

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resultBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var matches db.Matches

	doc.Find("table tbody tr").Each(func(idx int, selection *goquery.Selection) {

		if _, found := selection.Attr("matchid"); found {

			match := db.Match{
				Category: enum.MatchCategory_Zjc.Value,
				Status:   enum.MatchStatus_UnStart.Value,
			}

			match.League, _ = selection.Attr("gamename")

			selection.Find("td").Each(func(i int, selection *goquery.Selection) {

				if i == 2 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.StartAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 3 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.CloseAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 4 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.HomeTeam = strings.TrimSpace(selection.Text())
					})
				}
				if i == 6 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.GuestTeam = strings.TrimSpace(selection.Text())
					})
				}

				if id, ex := selection.Attr("id"); ex && strings.HasPrefix(id, "cell_") {

					name, _ := selection.Attr("title")

					var value float64
					selection.Find("span").Each(func(i int, selection *goquery.Selection) {
						value = conv.Float64(strings.TrimSpace(selection.Text()))
					})

					match.Odds.GoalsItems = append(match.Odds.GoalsItems, types.OddsItem{
						Name:   name,
						Result: strings.ReplaceAll(name, "球", ""),
						Value:  value,
					})
				}
			})

			matches = append(matches, &match)

		}

	})

	return matches, nil
}

// 胜平负
func (t ZcMatchFetcher) fetchZcMatchesWithOdds(ctx context.Context, date string) (db.Matches, error) {

	// date 指指定日期可买的场次
	if date == "" {
		date = time.Now().Format(time.DateOnly)
	}

	url := fmt.Sprintf("https://www.310win.com/buy/jingcai.aspx?typeID=105&oddstype=2&date=%s", date)

	resultBytes, err := new(helper.Http).Get(url, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resultBytes))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var matches db.Matches

	doc.Find("table tbody tr").Each(func(idx int, selection *goquery.Selection) {

		if _, found := selection.Attr("matchid"); found {

			match := db.Match{
				Category: enum.MatchCategory_Zjc.Value,
				Status:   enum.MatchStatus_UnStart.Value,
			}

			odds := types.Odds{}
			match.League, _ = selection.Attr("gamename")

			selection.Find("td").Each(func(i int, selection *goquery.Selection) {

				if i == 2 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.StartAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 3 {
					title, _ := selection.Attr("title")
					parts := strings.Split(title, "：")

					if len(parts) == 2 {
						match.CloseAt = t.parseTime(parts[1][5:])
					}
				}

				if i == 4 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.HomeTeam = strings.TrimSpace(selection.Text())
					})
				}
				if i == 7 {
					selection.Find("a").Each(func(i int, selection *goquery.Selection) {
						match.GuestTeam = strings.TrimSpace(selection.Text())
					})
				}

				if i == 12 {
					selection.Find("table tbody tr").Each(func(i int, selection *goquery.Selection) {
						if i == 0 {
							selection.Find("td").Each(func(i int, selection *goquery.Selection) {
								if i == 1 {

									item := types.OddsItem{}

									item.Name, _ = selection.Attr("title")
									item.Result = "3"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})

									odds.Items = append(odds.Items, item)
								}
								if i == 2 {
									item := types.OddsItem{}
									item.Name, _ = selection.Attr("title")
									item.Result = "1"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})

									odds.Items = append(odds.Items, item)

								}
								if i == 3 {
									item := types.OddsItem{}
									item.Name, _ = selection.Attr("title")
									item.Result = "0"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										item.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})

									odds.Items = append(odds.Items, item)
								}
							})
						}

						if i == 1 {
							selection.Find("td").Each(func(i int, selection *goquery.Selection) {
								if i == 0 {
									selection.Find("font").Each(func(i int, selection *goquery.Selection) {
										match.RCount = conv.Int(strings.TrimSpace(selection.Text()))
									})
								}
								if i == 1 {
									rItem := types.OddsItem{}
									rItem.Name = "让球主胜"
									rItem.Result = "3"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										rItem.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})

									odds.RItems = append(odds.RItems, rItem)
								}
								if i == 2 {
									rItem := types.OddsItem{}
									rItem.Name = "让球平"
									rItem.Result = "1"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										rItem.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})
									odds.RItems = append(odds.RItems, rItem)
								}
								if i == 3 {
									rItem := types.OddsItem{}
									rItem.Name = "让球客胜"
									rItem.Result = "0"

									selection.Find("span").Each(func(i int, selection *goquery.Selection) {
										rItem.Value = conv.Float64(strings.TrimSpace(selection.Text()))
									})
									odds.RItems = append(odds.RItems, rItem)
								}
							})
						}
					})
				}

			})

			match.Odds = odds

			matches = append(matches, &match)

		}

	})

	return matches, nil
}

func (t ZcMatchFetcher) parseTime(val string) time.Time {

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
