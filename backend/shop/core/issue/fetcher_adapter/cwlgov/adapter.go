package cwlgov

import (
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type Adapter struct {
}

func (a Adapter) FindX7cResultByIndex(index string) (*types.X7cResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindF3dResultByIndex(index string) (*types.F3dResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindDltResultByIndex(index string) (*types.DltResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindSsqResultByIndex(index string) (*types.SsqResult, error) {

	//url := "http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=%s&issueEnd=%s&dayStart=&dayEnd=&pageNo=1&pageSize=30&week=&systemType=PC"
	//url = fmt.Sprintf(url, index, index)
	//
	//resultBytes, err := new(helper.Http).Get(url, true)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}
	//
	//var ssq SSQResult
	//err = conv.B2S[SSQResult](resultBytes, &ssq)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}
	//
	//if len(ssq.Result) == 0 {
	//	return nil, nil
	//}
	//
	//x := ssq.Result[0]
	//
	//if len(x.Prizegrades) == 0 {
	//	return nil, nil
	//}
	//
	//var result = SingleSsqTicket{
	//	Red:       strings.Split(x.Red, ","),
	//	Blue:      []string{x.Blue},
	//	Sales:     x.Sales,
	//	PoolMoney: x.Poolmoney,
	//}
	//
	//for _, xx := range x.Prizegrades {
	//	if xx.Typemoney == "" {
	//		continue
	//	}
	//
	//	result.PrizeGrades = append(result.PrizeGrades, types.PrizeGrade{
	//		Grade:  xx.Type,
	//		Count:  xx.Typenum,
	//		Amount: xx.Typemoney,
	//	})
	//}
	//
	//date, _ := time.ParseInLocation(time.DateOnly, x.Date[:10], timed.LocAsiaShanghai)
	//y, m, d := date.Date()
	//
	//prizedAt := time.Date(y, m, d, PrizeHour, PrizeMinute, 0, 0, timed.LocAsiaShanghai)
	//
	//return &types.PrizeResult{
	//	Issue:   x.Code,
	//	Result:  conv.S2J(SingleSsqTickets{&result}),
	//	PrizeAt: prizedAt,
	//}, nil

	//TODO implement me
	panic("implement me")
}

//type SSQResult struct {
//	//State    int    `json:"state"`
//	//Message  string `json:"message"`
//	//Total    int    `json:"total"`
//	//PageNum  int    `json:"pageNum"`
//	//PageNo   int    `json:"pageNo"`
//	//PageSize int    `json:"pageSize"`
//	//Tflag    int    `json:"Tflag"`
//	Result []struct {
//		//Name string `json:"name"`
//		Code string `json:"code"`
//		//DetailsLink string `json:"detailsLink"`
//		//VideoLink   string `json:"videoLink"`
//		Date string `json:"date"` // 2023-05-09(二)
//		//Week        string `json:"week"`
//		Red  string `json:"red"`
//		Blue string `json:"blue"`
//		//Blue2       string `json:"blue2"`
//		Sales     string `json:"sales"`
//		Poolmoney string `json:"poolmoney"`
//		//Content     string `json:"content"`
//		//Addmoney    string `json:"addmoney"`
//		//Addmoney2   string `json:"addmoney2"`
//		//Msg         string `json:"msg"`
//		//Z2Add       string `json:"z2add"`
//		//M2Add       string `json:"m2add"`
//		Prizegrades []struct {
//			Type      int    `json:"type"`
//			Typenum   string `json:"typenum"`
//			Typemoney string `json:"typemoney"`
//		} `json:"prizegrades"`
//	} `json:"result"`
//}
//

func (a Adapter) FindPl3ResultByIndex(index string) (*types.Pl3Result, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) FindPl5ResultByIndex(index string) (*types.Pl5Result, error) {
	//TODO implement me
	panic("implement me")
}
