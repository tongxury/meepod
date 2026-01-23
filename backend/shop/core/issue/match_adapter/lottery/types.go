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
		VtoolsConfig struct {
			OffLineSaleStatus  int    `json:"offLineSaleStatus"`
			OffLineStopMessage string `json:"offLineStopMessage"`
			OnLineSaleStatus   int    `json:"onLineSaleStatus"`
			OnLineStopMessage  string `json:"onLineStopMessage"`
		} `json:"vtoolsConfig"`
		MatchInfoList []struct {
			BusinessDate string `json:"businessDate"`
			SubMatchList []struct {
				AwayRank          string `json:"awayRank"`
				AwayTeamAbbEnName string `json:"awayTeamAbbEnName"`
				AwayTeamAbbName   string `json:"awayTeamAbbName"`
				AwayTeamAllName   string `json:"awayTeamAllName"`
				AwayTeamCode      string `json:"awayTeamCode"`
				AwayTeamId        int    `json:"awayTeamId"`
				BackColor         string `json:"backColor"`
				BaseAwayTeamId    int    `json:"baseAwayTeamId"`
				BaseHomeTeamId    int    `json:"baseHomeTeamId"`
				BettingAllUp      int    `json:"bettingAllUp"`
				BettingSingle     int    `json:"bettingSingle"`
				BusinessDate      string `json:"businessDate"`
				Crs               struct {
					GoalLine   string `json:"goalLine"`
					S00S00     string `json:"s00s00"`
					S00S00F    string `json:"s00s00f"`
					S00S01     string `json:"s00s01"`
					S00S01F    string `json:"s00s01f"`
					S00S02     string `json:"s00s02"`
					S00S02F    string `json:"s00s02f"`
					S00S03     string `json:"s00s03"`
					S00S03F    string `json:"s00s03f"`
					S00S04     string `json:"s00s04"`
					S00S04F    string `json:"s00s04f"`
					S00S05     string `json:"s00s05"`
					S00S05F    string `json:"s00s05f"`
					S01S00     string `json:"s01s00"`
					S01S00F    string `json:"s01s00f"`
					S01S01     string `json:"s01s01"`
					S01S01F    string `json:"s01s01f"`
					S01S02     string `json:"s01s02"`
					S01S02F    string `json:"s01s02f"`
					S01S03     string `json:"s01s03"`
					S01S03F    string `json:"s01s03f"`
					S01S04     string `json:"s01s04"`
					S01S04F    string `json:"s01s04f"`
					S01S05     string `json:"s01s05"`
					S01S05F    string `json:"s01s05f"`
					S02S00     string `json:"s02s00"`
					S02S00F    string `json:"s02s00f"`
					S02S01     string `json:"s02s01"`
					S02S01F    string `json:"s02s01f"`
					S02S02     string `json:"s02s02"`
					S02S02F    string `json:"s02s02f"`
					S02S03     string `json:"s02s03"`
					S02S03F    string `json:"s02s03f"`
					S02S04     string `json:"s02s04"`
					S02S04F    string `json:"s02s04f"`
					S02S05     string `json:"s02s05"`
					S02S05F    string `json:"s02s05f"`
					S03S00     string `json:"s03s00"`
					S03S00F    string `json:"s03s00f"`
					S03S01     string `json:"s03s01"`
					S03S01F    string `json:"s03s01f"`
					S03S02     string `json:"s03s02"`
					S03S02F    string `json:"s03s02f"`
					S03S03     string `json:"s03s03"`
					S03S03F    string `json:"s03s03f"`
					S04S00     string `json:"s04s00"`
					S04S00F    string `json:"s04s00f"`
					S04S01     string `json:"s04s01"`
					S04S01F    string `json:"s04s01f"`
					S04S02     string `json:"s04s02"`
					S04S02F    string `json:"s04s02f"`
					S05S00     string `json:"s05s00"`
					S05S00F    string `json:"s05s00f"`
					S05S01     string `json:"s05s01"`
					S05S01F    string `json:"s05s01f"`
					S05S02     string `json:"s05s02"`
					S05S02F    string `json:"s05s02f"`
					S1Sa       string `json:"s1sa"`
					S1Saf      string `json:"s1saf"`
					S1Sd       string `json:"s1sd"`
					S1Sdf      string `json:"s1sdf"`
					S1Sh       string `json:"s1sh"`
					S1Shf      string `json:"s1shf"`
					UpdateDate string `json:"updateDate"`
					UpdateTime string `json:"updateTime"`
				} `json:"crs"`
				GroupName string `json:"groupName"`
				Had       struct {
					A          string `json:"a,omitempty"`
					Af         string `json:"af,omitempty"`
					D          string `json:"d,omitempty"`
					Df         string `json:"df,omitempty"`
					GoalLine   string `json:"goalLine,omitempty"`
					H          string `json:"h,omitempty"`
					Hf         string `json:"hf,omitempty"`
					UpdateDate string `json:"updateDate,omitempty"`
					UpdateTime string `json:"updateTime,omitempty"`
				} `json:"had"`
				Hafu struct {
					Aa         string `json:"aa"`
					Aaf        string `json:"aaf"`
					Ad         string `json:"ad"`
					Adf        string `json:"adf"`
					Ah         string `json:"ah"`
					Ahf        string `json:"ahf"`
					Da         string `json:"da"`
					Daf        string `json:"daf"`
					Dd         string `json:"dd"`
					Ddf        string `json:"ddf"`
					Dh         string `json:"dh"`
					Dhf        string `json:"dhf"`
					GoalLine   string `json:"goalLine"`
					Ha         string `json:"ha"`
					Haf        string `json:"haf"`
					Hd         string `json:"hd"`
					Hdf        string `json:"hdf"`
					Hh         string `json:"hh"`
					Hhf        string `json:"hhf"`
					Id         int    `json:"id"`
					UpdateDate string `json:"updateDate"`
					UpdateTime string `json:"updateTime"`
				} `json:"hafu"`
				Hhad struct {
					A          string `json:"a"`
					Af         string `json:"af"`
					D          string `json:"d"`
					Df         string `json:"df"`
					GoalLine   string `json:"goalLine"`
					H          string `json:"h"`
					Hf         string `json:"hf"`
					UpdateDate string `json:"updateDate"`
					UpdateTime string `json:"updateTime"`
				} `json:"hhad"`
				HomeRank          string `json:"homeRank"`
				HomeTeamAbbEnName string `json:"homeTeamAbbEnName"`
				HomeTeamAbbName   string `json:"homeTeamAbbName"`
				HomeTeamAllName   string `json:"homeTeamAllName"`
				HomeTeamCode      string `json:"homeTeamCode"`
				HomeTeamId        int    `json:"homeTeamId"`
				IsHide            int    `json:"isHide"`
				IsHot             int    `json:"isHot"`
				LeagueAbbName     string `json:"leagueAbbName"`
				LeagueAllName     string `json:"leagueAllName"`
				LeagueCode        string `json:"leagueCode"`
				LeagueId          int    `json:"leagueId"`
				LineNum           string `json:"lineNum"`
				MatchDate         string `json:"matchDate"`
				MatchId           int    `json:"matchId"`
				MatchName         string `json:"matchName"`
				MatchNum          int    `json:"matchNum"`
				MatchNumStr       string `json:"matchNumStr"`
				MatchStatus       string `json:"matchStatus"`
				MatchTime         string `json:"matchTime"`
				MatchWeek         string `json:"matchWeek"`
				OddsList          []struct {
					A          string `json:"a"`
					D          string `json:"d"`
					GoalLine   string `json:"goalLine"`
					GoalLineF  string `json:"goalLineF"`
					H          string `json:"h"`
					MatchId    int    `json:"matchId"`
					MatchNum   int    `json:"matchNum"`
					Odds       string `json:"odds"`
					PoolCode   string `json:"poolCode"`
					PoolId     int    `json:"poolId"`
					UpdateDate string `json:"updateDate"`
					UpdateTime string `json:"updateTime"`
				} `json:"oddsList"`
				PoolList []struct {
					AllUp             int    `json:"allUp"`
					BettingAllup      int    `json:"bettingAllup"`
					BettingSingle     int    `json:"bettingSingle"`
					CbtAllUp          int    `json:"cbtAllUp"`
					CbtSingle         int    `json:"cbtSingle"`
					CbtValue          int    `json:"cbtValue"`
					FixedOddsgoalLine string `json:"fixedOddsgoalLine"`
					IntAllUp          int    `json:"intAllUp"`
					IntSingle         int    `json:"intSingle"`
					IntValue          int    `json:"intValue"`
					MatchId           int    `json:"matchId"`
					MatchNum          int    `json:"matchNum"`
					PoolCloseDate     string `json:"poolCloseDate"`
					PoolCloseTime     string `json:"poolCloseTime"`
					PoolCode          string `json:"poolCode"`
					PoolId            int    `json:"poolId"`
					PoolOddsType      string `json:"poolOddsType"`
					PoolStatus        string `json:"poolStatus"`
					SellInitialDate   string `json:"sellInitialDate"`
					SellInitialTime   string `json:"sellInitialTime"`
					Single            int    `json:"single"`
					UpdateDate        string `json:"updateDate"`
					UpdateTime        string `json:"updateTime"`
					VbtAllUp          int    `json:"vbtAllUp"`
					VbtSingle         int    `json:"vbtSingle"`
					VbtValue          int    `json:"vbtValue"`
				} `json:"poolList"`
				Remark     string `json:"remark"`
				SellStatus int    `json:"sellStatus"`
				Ttg        struct {
					GoalLine   string `json:"goalLine"`
					S0         string `json:"s0"`
					S0F        string `json:"s0f"`
					S1         string `json:"s1"`
					S1F        string `json:"s1f"`
					S2         string `json:"s2"`
					S2F        string `json:"s2f"`
					S3         string `json:"s3"`
					S3F        string `json:"s3f"`
					S4         string `json:"s4"`
					S4F        string `json:"s4f"`
					S5         string `json:"s5"`
					S5F        string `json:"s5f"`
					S6         string `json:"s6"`
					S6F        string `json:"s6f"`
					S7         string `json:"s7"`
					S7F        string `json:"s7f"`
					UpdateDate string `json:"updateDate"`
					UpdateTime string `json:"updateTime"`
				} `json:"ttg"`
				Vote struct {
				} `json:"vote"`
			} `json:"subMatchList"`
			Weekday    string `json:"weekday"`
			MatchCount int    `json:"matchCount"`
		} `json:"matchInfoList"`
		MatchDateList []struct {
			BusinessDate   string `json:"businessDate"`
			BusinessDateCn string `json:"businessDateCn"`
			MatchDate      string `json:"matchDate"`
			MatchDateCn    string `json:"matchDateCn"`
		} `json:"matchDateList"`
		AllUpList struct {
			HHAD []struct {
				FValue        int    `json:"fValue"`
				Formula       string `json:"formula"`
				FormulaType   int    `json:"formulaType"`
				MaxMatchCount int    `json:"maxMatchCount"`
				PoolCode      string `json:"poolCode"`
			} `json:"HHAD"`
			CRS []struct {
				FValue        int    `json:"fValue"`
				Formula       string `json:"formula"`
				FormulaType   int    `json:"formulaType"`
				MaxMatchCount int    `json:"maxMatchCount"`
				PoolCode      string `json:"poolCode"`
			} `json:"CRS"`
			TTG []struct {
				FValue        int    `json:"fValue"`
				Formula       string `json:"formula"`
				FormulaType   int    `json:"formulaType"`
				MaxMatchCount int    `json:"maxMatchCount"`
				PoolCode      string `json:"poolCode"`
			} `json:"TTG"`
			HAFU []struct {
				FValue        int    `json:"fValue"`
				Formula       string `json:"formula"`
				FormulaType   int    `json:"formulaType"`
				MaxMatchCount int    `json:"maxMatchCount"`
				PoolCode      string `json:"poolCode"`
			} `json:"HAFU"`
			HAD []struct {
				FValue        int    `json:"fValue"`
				Formula       string `json:"formula"`
				FormulaType   int    `json:"formulaType"`
				MaxMatchCount int    `json:"maxMatchCount"`
				PoolCode      string `json:"poolCode"`
			} `json:"HAD"`
		} `json:"allUpList"`
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
