package xinsh

type TradeParams struct {
	MerchantNo string
	OrderId    string
	Amount     float64
	//Type         string // ALIPAY
	Subject      string
	UserClientIp string
	TimeExpire   int // 分钟
	NotifyUrl    string
}

type MerchantParams struct {
	RequestId  string
	MerchantNm string // 自然人商户需与法人姓名一致
	//MerchantType string
	Province        string
	City            string
	District        string
	Address         string
	Username        string
	Email           string
	Phone           string
	IdCardNo        string
	IdCardFrom      string
	IdCardTo        string
	IdCardFront     string
	IdCardBack      string
	IdCardHandled   string
	StoreFront      string
	StoreInSide     string
	BankAccount     string
	BankAccountName string
	AccountType     string // 1 公 2 私
	AccountNature   string // 1-法人账户 2-非法人账户 对私必填
	CnapsCode       string
	BankId          string
	BankCardFront   string
	BankPhone       string
	BankName        string
	BankProvince    string
	BankCity        string
}

type Upload struct {
	PictureUrl  string `json:"pictureUrl"`
	PictureName string `json:"pictureName"`
}

type MerchantApply struct {
	MerchantNo string `json:"merchantNo"`
	RequestId  string `json:"requestId"`
}

type Resp[T any] struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	RespData *T     `json:"respData"`
}

type PayInfo struct {
	MerchantNo string `json:"merchantNo"`
	OrderNo    string `json:"orderNo"`
	OutOrderNo string `json:"outOrderNo"`
	PayUrl     string `json:"payUrl"`
}

type ApplyState struct {
	MerchantNo                string      `json:"merchantNo"`
	RequestId                 string      `json:"requestId"`
	OpStatus                  string      `json:"opStatus"`
	Suggestion                interface{} `json:"suggestion"`
	WechatPayRecordMerchantNo interface{} `json:"wechatPayRecordMerchantNo"`
	AliPayRecordMerchantNo    interface{} `json:"aliPayRecordMerchantNo"`
	UnionPayRecordMerchantNo  interface{} `json:"unionPayRecordMerchantNo"`
}

type TradeResult struct {
	OrderNo          string      `json:"orderNo"`
	OutOrderNo       string      `json:"outOrderNo"`
	ThirdPartyUuid   string      `json:"thirdPartyUuid"`
	TransactionId    string      `json:"transactionId"`
	MerchantNo       string      `json:"merchantNo"`
	Amt              string      `json:"amt"`
	FeeAmt           int         `json:"feeAmt"`
	PayType          string      `json:"payType"`
	PayWay           string      `json:"payWay"`
	DrType           interface{} `json:"drType"`
	TranSts          string      `json:"tranSts"`
	Scene            string      `json:"scene"`
	PayTime          string      `json:"payTime"`
	TranTime         string      `json:"tranTime"`
	TerminalId       interface{} `json:"terminalId"`
	DeviceNo         interface{} `json:"deviceNo"`
	Subject          string      `json:"subject"`
	BuyerId          interface{} `json:"buyerId"`
	Extend           interface{} `json:"extend"`
	CouponDetail     interface{} `json:"couponDetail"`
	BankType         string      `json:"bankType"`
	PromotionDetail  interface{} `json:"promotionDetail"`
	SettleAmt        int         `json:"settleAmt"`
	BuyerPayAmount   int         `json:"buyerPayAmount"`
	SettleWay        string      `json:"settleWay"`
	SettleStatus     interface{} `json:"settleStatus"`
	SettleFeeRate    interface{} `json:"settleFeeRate"`
	SettleFeeAmt     interface{} `json:"settleFeeAmt"`
	ActualArrivalAmt interface{} `json:"actualArrivalAmt"`
}

func (t *TradeResult) IsSuccess() bool {
	return t.TranSts == "SUCCESS"
}

func (t *TradeResult) Paying() bool {
	return t.TranSts == "PAYING" || t.TranSts == "NEEDPAY"
}
