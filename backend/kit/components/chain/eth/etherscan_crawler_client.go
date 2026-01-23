package eth

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"github.com/gocolly/colly/v2"
	"strings"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

//const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

func NewEtherScanCrawlerClient() *EtherScanCrawlerClient {
	return &EtherScanCrawlerClient{}
}

type EtherScanCrawlerClient struct {
}

type EToken struct {
	Address        string
	TxHash         string
	Icon           string
	Name           string
	Symbol         string
	Price          float64
	Volume24h      float64
	Holders        int64
	CreatedAtBlock int64
}

type ETokens []*EToken

func (e *EtherScanCrawlerClient) FindAddressMeta(address string) (*EToken, error) {
	//https://etherscan.io/token/0xdac17f958d2ee523a2206206994597c13d831ec7
	url := "https://etherscan.io/address/" + address

	c := colly.NewCollector()
	c.UserAgent = UserAgent

	rsp := EToken{Address: address}

	c.OnHTML("div #ContentPlaceHolder1_trContract", func(e *colly.HTMLElement) {

		e.ForEach("a", func(i int, a *colly.HTMLElement) {
			if i == 1 {
				rsp.TxHash = a.Text
			}
		})

	})

	c.OnError(func(response *colly.Response, err error) {
	})

	c.OnRequest(func(r *colly.Request) {
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return &rsp, nil
}

type ETxMeta struct {
	TxHash         string
	CreatedAtBlock int64
}

type ETxMetas []*ETxMeta

func (e *EtherScanCrawlerClient) FindTxMeta(tx string) (*ETxMeta, error) {
	url := "https://etherscan.io/tx/" + tx

	c := colly.NewCollector()
	c.UserAgent = UserAgent

	rsp := ETxMeta{TxHash: tx}

	c.OnHTML("div #ContentPlaceHolder1_maintable", func(e *colly.HTMLElement) {

		e.ForEach("a", func(i int, a *colly.HTMLElement) {
			if i == 1 {
				rsp.CreatedAtBlock = conv.Int64(a.Text)
			}
		})

	})

	c.OnError(func(response *colly.Response, err error) {
	})

	c.OnRequest(func(r *colly.Request) {
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return &rsp, nil
}

func (e *EtherScanCrawlerClient) ListTopTokens(page int) (ETokens, error) {

	//ctx := context.Background()
	url := fmt.Sprintf("https://etherscan.io/tokens?&sort=24h_volume_usd&order=desc&ps=100&p=%d", page)

	partTokens, err := e.collect(url)
	if err != nil {
		slf.WithError(err).Errorw("collect err")
		return nil, err
	}

	return partTokens, nil
}

func (e *EtherScanCrawlerClient) collect(url string) (ETokens, error) {

	//url := "https://www.baidu.com/"

	c := colly.NewCollector()
	c.UserAgent = UserAgent

	var rsp ETokens

	c.OnHTML("table tbody", func(tb *colly.HTMLElement) {

		tb.ForEach("tr", func(i int, tr *colly.HTMLElement) {

			var address, icon, name, symbol, price, volume24h, holders string

			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				if i == 1 {
					uri := td.ChildAttr("a", "href")
					address = strings.ReplaceAll(uri, "/token/", "")
					address = strings.ToLower(address)

					icon = td.ChildAttr("a img", "src")

					name = td.ChildText("a div div")
					symbol = td.ChildText("a div span")
					symbol = strings.ReplaceAll(symbol, ")", "")
					symbol = strings.ReplaceAll(symbol, "(", "")
				}

				if i == 2 {
					td.ForEach("div", func(i int, d *colly.HTMLElement) {
						if i == 0 {
							price = strings.Trim(d.Text, "\n")
							price = strings.ReplaceAll(price, "$", "")
						}
					})
				}

				if i == 4 {
					volume24h = strings.Trim(td.Text, "\n")
					volume24h = strings.ReplaceAll(volume24h, "$", "")
					volume24h = strings.ReplaceAll(volume24h, ",", "")
				}

				if i == 7 {
					td.ForEach("div", func(i int, d *colly.HTMLElement) {
						if i == 0 {
							holders = strings.ReplaceAll(d.Text, ",", "")
						}
					})
				}
			})

			rsp = append(rsp, &EToken{
				Address:   address,
				Icon:      icon,
				Name:      name,
				Symbol:    symbol,
				Price:     conv.Float64(price),
				Volume24h: conv.Float64(volume24h),
				Holders:   conv.Int64(holders),
			})
		})

	})

	c.OnError(func(response *colly.Response, err error) {
	})

	c.OnRequest(func(r *colly.Request) {
	})

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	return rsp, nil
}
