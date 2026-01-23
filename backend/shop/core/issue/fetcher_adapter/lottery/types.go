package lottery

type HttpResult struct {
	//DataFrom     string `json:"dataFrom"`
	//EmptyFlag    bool   `json:"emptyFlag"`
	//ErrorCode    string `json:"errorCode"`
	//ErrorMessage string `json:"errorMessage"`
	Success bool `json:"success"`
	Value   struct {
		//LastPoolDraw struct {
		//	LotteryDrawNum       string `json:"lotteryDrawNum"`
		//	LotteryDrawResult    string `json:"lotteryDrawResult"`
		//	LotteryDrawTime      string `json:"lotteryDrawTime"`
		//	LotteryGameName      string `json:"lotteryGameName"`
		//	LotteryGameNum       string `json:"lotteryGameNum"`
		//	PoolBalanceAfterdraw string `json:"poolBalanceAfterdraw"`
		//	PrizeLevelList       []struct {
		//		//AwardType        int    `json:"awardType"`
		//		Group string `json:"group"`
		//		//LotteryCondition string `json:"lotteryCondition"`
		//		PrizeLevel       string `json:"prizeLevel"`
		//		Sort             int    `json:"sort"`
		//		StakeAmount      string `json:"stakeAmount"`
		//		StakeCount       string `json:"stakeCount"`
		//		TotalPrizeamount string `json:"totalPrizeamount"`
		//	} `json:"prizeLevelList"`
		//} `json:"lastPoolDraw"`
		List []Result `json:"list"`
		//PageNo   int `json:"pageNo"`
		//PageSize int `json:"pageSize"`
		//Pages    int `json:"pages"`
		//Total    int `json:"total"`
	} `json:"value"`
}

type Result struct {
	//DrawFlowFund           string `json:"drawFlowFund"`
	//DrawFlowFundRj         string `json:"drawFlowFundRj"`
	//DrawPdfUrl             string `json:"drawPdfUrl"`
	//EstimateDrawTime       string `json:"estimateDrawTime"`
	//IsGetKjpdf             int    `json:"isGetKjpdf"`
	//IsGetXlpdf             int    `json:"isGetXlpdf"`
	LotteryDrawNum    string `json:"lotteryDrawNum"`
	LotteryDrawResult string `json:"lotteryDrawResult"`
	LotteryDrawStatus int    `json:"lotteryDrawStatus"`
	//LotteryDrawStatusNo    string `json:"lotteryDrawStatusNo"`
	LotteryDrawTime string `json:"lotteryDrawTime"`
	//LotteryEquipmentCount  int    `json:"lotteryEquipmentCount"`
	LotteryGameName string `json:"lotteryGameName"`
	//LotteryGameNum         string `json:"lotteryGameNum"`
	//LotteryGamePronum      int    `json:"lotteryGamePronum"`
	//LotteryNotice          int    `json:"lotteryNotice"`
	//LotteryNoticeShowFlag  int    `json:"lotteryNoticeShowFlag"`
	//LotteryPaidBeginTime   string `json:"lotteryPaidBeginTime"`
	//LotteryPaidEndTime     string `json:"lotteryPaidEndTime"`
	//LotteryPromotionFlag   int    `json:"lotteryPromotionFlag"`
	//LotteryPromotionFlagRj int    `json:"lotteryPromotionFlagRj"`
	//LotterySaleBeginTime   string `json:"lotterySaleBeginTime"`
	//LotterySaleEndTimeUnix  string        `json:"lotterySaleEndTimeUnix"`
	//LotterySaleEndtime      string        `json:"lotterySaleEndtime"`
	//LotterySuspendedFlag    int           `json:"lotterySuspendedFlag"`
	//LotteryUnsortDrawresult string        `json:"lotteryUnsortDrawresult"`
	//MatchList               []interface{} `json:"matchList"`
	//PdfType                 int           `json:"pdfType"`
	//PoolBalanceAfterdraw    string        `json:"poolBalanceAfterdraw"`
	//PoolBalanceAfterdrawRj  string        `json:"poolBalanceAfterdrawRj"`
	PrizeLevelList []struct {
		//AwardType        int    `json:"awardType"`
		Group string `json:"group"`
		//LotteryCondition string `json:"lotteryCondition"`
		PrizeLevel       string `json:"prizeLevel"`
		Sort             int    `json:"sort"`
		StakeAmount      string `json:"stakeAmount"`
		StakeCount       string `json:"stakeCount"`
		TotalPrizeamount string `json:"totalPrizeamount"`
	} `json:"prizeLevelList"`
	//PrizeLevelListRj  []interface{} `json:"prizeLevelListRj"`
	//RuleType          int           `json:"ruleType"`
	//SurplusAmount     string        `json:"surplusAmount"`
	//SurplusAmountRj   string        `json:"surplusAmountRj"`
	//TermList          []interface{} `json:"termList"`
	//TermResultList    []interface{} `json:"termResultList"`
	//TotalSaleAmount   string        `json:"totalSaleAmount"`
	//TotalSaleAmountRj string        `json:"totalSaleAmountRj"`
	//Verify            int           `json:"verify"`
	//VtoolsConfig      struct {} `json:"vtoolsConfig"`
}
