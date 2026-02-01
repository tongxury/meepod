package lottery

type ZcMatchResultsResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Value        struct {
		ResultCount int `json:"resultCount"`
		Total       int `json:"total"`
		Pages       int `json:"pages"`
		LeagueList  []struct {
			LeagueAbbName string `json:"leagueAbbName"`
			LeagueId      int    `json:"leagueId"`
			LeagueAllName string `json:"leagueAllName"`
		} `json:"leagueList"`
		PageNo      int `json:"pageNo"`
		MatchResult []struct {
			A                 string `json:"a"`
			AllAwayTeam       string `json:"allAwayTeam"`
			AllHomeTeam       string `json:"allHomeTeam"`
			AwayTeam          string `json:"awayTeam"`
			AwayTeamId        int    `json:"awayTeamId"`
			BettingSingle     int    `json:"bettingSingle"`
			D                 string `json:"d"`
			GoalLine          string `json:"goalLine"`
			H                 string `json:"h"`
			HomeTeam          string `json:"homeTeam"`
			HomeTeamId        int    `json:"homeTeamId"`
			LeagueBackColor   string `json:"leagueBackColor"`
			LeagueId          int    `json:"leagueId"`
			LeagueName        string `json:"leagueName"`
			LeagueNameAbbr    string `json:"leagueNameAbbr"`
			MatchDate         string `json:"matchDate"`
			MatchId           int    `json:"matchId"`
			MatchNum          string `json:"matchNum"`
			MatchNumStr       string `json:"matchNumStr"`
			MatchResultStatus string `json:"matchResultStatus"`
			PoolStatus        string `json:"poolStatus"`
			SectionsNo1       string `json:"sectionsNo1"`
			SectionsNo999     string `json:"sectionsNo999"`
			WinFlag           string `json:"winFlag"`
		} `json:"matchResult"`
		PageSize       int    `json:"pageSize"`
		LastUpdateTime string `json:"lastUpdateTime"`
	} `json:"value"`
	EmptyFlag bool   `json:"emptyFlag"`
	DataFrom  string `json:"dataFrom"`
	Success   bool   `json:"success"`
}

type ZcMatchesResponse struct {
	DataFrom     string `json:"dataFrom"`
	EmptyFlag    bool   `json:"emptyFlag"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		MatchInfoList []struct {
			BusinessDate string `json:"businessDate"`
			SubMatchList []struct {
				AwayTeamAbbName string `json:"awayTeamAbbName"`
				AwayTeamAllName string `json:"awayTeamAllName"`
				AwayTeamCode    string `json:"awayTeamCode"`
				AwayTeamId      int    `json:"awayTeamId"`
				BackColor       string `json:"backColor"`
				BusinessDate    string `json:"businessDate"`
				HomeTeamAbbName string `json:"homeTeamAbbName"`
				HomeTeamAllName string `json:"homeTeamAllName"`
				HomeTeamCode    string `json:"homeTeamCode"`
				HomeTeamId      int    `json:"homeTeamId"`
				LeagueAbbName   string `json:"leagueAbbName"`
				LeagueAllName   string `json:"leagueAllName"`
				LeagueId        string `json:"leagueId"`
				LineNum         string `json:"lineNum"`
				MatchDate       string `json:"matchDate"`
				MatchId         int    `json:"matchId"`
				MatchName       string `json:"matchName"`
				MatchNum        int    `json:"matchNum"`
				MatchNumDate    string `json:"matchNumDate"`
				MatchNumStr     string `json:"matchNumStr"`
				MatchStatus     string `json:"matchStatus"`
				MatchTime       string `json:"matchTime"`
				MatchWeek       string `json:"matchWeek"`
				Weekday         string `json:"weekday"`
				Remark          string `json:"remark"`
				SellStatus      string `json:"sellStatus"`
				PoolStatus      string `json:"poolStatus"`
				OddsList        []struct {
					A             string `json:"a"`
					D             string `json:"d"`
					H             string `json:"h"`
					GoalLine      string `json:"goalLine"`
					GoalLineF     string `json:"goalLineF"`
					GoalLineValue string `json:"goalLineValue"`
					MatchId       int    `json:"matchId"`
					MatchNum      int    `json:"matchNum"`
					Odds          string `json:"odds"`
					PoolCode      string `json:"poolCode"`
					PoolId        int    `json:"poolId"`
					UpdateDate    string `json:"updateDate"`
					UpdateTime    string `json:"updateTime"`
				} `json:"oddsList"`
				PoolList []struct {
					CbtAllUp      int    `json:"cbtAllUp"`
					CbtSingle     int    `json:"cbtSingle"`
					CbtValue      int    `json:"cbtValue"`
					IntAllUp      int    `json:"intAllUp"`
					IntSingle     int    `json:"intSingle"`
					IntValue      int    `json:"intValue"`
					PoolCode      string `json:"poolCode"`
					PoolPrimaryId int    `json:"poolPrimaryId"`
					PoolStatus    string `json:"poolStatus"`
				} `json:"poolList"`
			} `json:"subMatchList"`
			Weekday      string `json:"weekday"`
			MatchCount   int    `json:"matchCount"`
			MatchNumDate string `json:"matchNumDate"`
		} `json:"matchInfoList"`
		MatchDateList []struct {
			BusinessDate   string `json:"businessDate"`
			BusinessDateCn string `json:"businessDateCn"`
			MatchDate      string `json:"matchDate"`
			MatchDateCn    string `json:"matchDateCn"`
		} `json:"matchDateList"`
		LeagueList []struct {
			LeagueId       string `json:"leagueId"`
			LeagueName     string `json:"leagueName"`
			LeagueNameAbbr string `json:"leagueNameAbbr"`
		} `json:"leagueList"`
		TotalCount     int    `json:"totalCount"`
		LastUpdateTime string `json:"lastUpdateTime"`
	} `json:"value"`
}

type Z14MatchResultsResponse struct {
	DataFrom     string `json:"dataFrom"`
	EmptyFlag    bool   `json:"emptyFlag"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		DelayRemark            string `json:"delayRemark"`
		DrawFlowFund           string `json:"drawFlowFund"`
		DrawFlowFundRj         string `json:"drawFlowFundRj"`
		DrawPdfUrl             string `json:"drawPdfUrl"`
		DrawPdfUrlRj           string `json:"drawPdfUrlRj"`
		EstimateDrawTime       string `json:"estimateDrawTime"`
		IsDelay                int    `json:"isDelay"`
		IsGetKjpdf             int    `json:"isGetKjpdf"`
		IsGetXlpdf             int    `json:"isGetXlpdf"`
		LotteryDrawNum         string `json:"lotteryDrawNum"`
		LotteryDrawResult      string `json:"lotteryDrawResult"`
		LotteryDrawTime        string `json:"lotteryDrawTime"`
		LotteryGameName        string `json:"lotteryGameName"`
		LotteryGameNum         string `json:"lotteryGameNum"`
		LotteryPaidBeginTime   string `json:"lotteryPaidBeginTime"`
		LotteryPaidEndTime     string `json:"lotteryPaidEndTime"`
		LotteryPromotionFlag   int    `json:"lotteryPromotionFlag"`
		LotteryPromotionFlagRj int    `json:"lotteryPromotionFlagRj"`
		LotterySaleBeginTime   string `json:"lotterySaleBeginTime"`
		LotterySaleEndtime     string `json:"lotterySaleEndtime"`
		MatchList              []struct {
			A                 string `json:"a"`
			CzHalfScore       string `json:"czHalfScore"`
			CzScore           string `json:"czScore"`
			D                 string `json:"d"`
			GuestTeamAllName  string `json:"guestTeamAllName"`
			GuestTeamName     string `json:"guestTeamName"`
			H                 string `json:"h"`
			InfohubMatchId    int    `json:"infohubMatchId"`
			LeagueId          int    `json:"leagueId"`
			MasterTeamAllName string `json:"masterTeamAllName"`
			MasterTeamName    string `json:"masterTeamName"`
			MatchName         string `json:"matchName"`
			MatchNum          int    `json:"matchNum"`
			Result            string `json:"result"`
			StartTime         string `json:"startTime"`
		} `json:"matchList"`
		PoolBalanceAfterdraw   string `json:"poolBalanceAfterdraw"`
		PoolBalanceAfterdrawRj string `json:"poolBalanceAfterdrawRj"`
		PrizeLevelList         []struct {
			AwardType        int    `json:"awardType"`
			Group            string `json:"group"`
			LotteryCondition string `json:"lotteryCondition"`
			PrizeLevel       string `json:"prizeLevel"`
			Sort             int    `json:"sort"`
			StakeAmount      string `json:"stakeAmount"`
			StakeCount       string `json:"stakeCount"`
			TotalPrizeamount string `json:"totalPrizeamount"`
		} `json:"prizeLevelList"`
		SalePdfUrl        string `json:"salePdfUrl"`
		SalePdfUrlRj      string `json:"salePdfUrlRj"`
		SurplusAmount     string `json:"surplusAmount"`
		SurplusAmountRj   string `json:"surplusAmountRj"`
		TotalSaleAmount   string `json:"totalSaleAmount"`
		TotalSaleAmountRj string `json:"totalSaleAmountRj"`
		Verify            int    `json:"verify"`
	} `json:"value"`
}

type Z14MatchesResponse struct {
	DataFrom     string `json:"dataFrom"`
	EmptyFlag    bool   `json:"emptyFlag"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		BqcMatch struct {
		} `json:"bqcMatch"`
		Bqclist  []interface{} `json:"bqclist"`
		JqcMatch struct {
		} `json:"jqcMatch"`
		Jqclist  []interface{} `json:"jqclist"`
		SfcMatch struct {
			EstimateDrawTime     string `json:"estimateDrawTime"`
			LotteryDrawNum       string `json:"lotteryDrawNum"`
			LotteryDrawTime      string `json:"lotteryDrawTime"`
			LotteryGameName      string `json:"lotteryGameName"`
			LotteryGameNum       string `json:"lotteryGameNum"`
			LotterySaleBegintime string `json:"lotterySaleBegintime"`
			LotterySaleEndtime   string `json:"lotterySaleEndtime"`
			MatchList            []struct {
				A                 string `json:"a"`
				CzHalfScore       string `json:"czHalfScore"`
				CzScore           string `json:"czScore"`
				D                 string `json:"d"`
				GuestTeamAllName  string `json:"guestTeamAllName"`
				GuestTeamName     string `json:"guestTeamName"`
				H                 string `json:"h"`
				InfohubMatchId    int    `json:"infohubMatchId"`
				LeagueId          int    `json:"leagueId"`
				MasterTeamAllName string `json:"masterTeamAllName"`
				MasterTeamName    string `json:"masterTeamName"`
				MatchName         string `json:"matchName"`
				MatchNum          int    `json:"matchNum"`
				Result            string `json:"result"`
				StartTime         string `json:"startTime"`
			} `json:"matchList"`
		} `json:"sfcMatch"`
		Sfclist []string `json:"sfclist"`
	} `json:"value"`
}
