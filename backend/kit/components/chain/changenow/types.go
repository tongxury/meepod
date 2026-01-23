package changenow

import "time"

type ChOrder struct {
	FromAmount       float64 `json:"fromAmount,omitempty"`
	ToAmount         float64 `json:"toAmount,omitempty"`
	Flow             string  `json:"flow,omitempty"`
	Type             string  `json:"type,omitempty"`
	PayinAddress     string  `json:"payinAddress,omitempty"`
	PayoutAddress    string  `json:"payoutAddress,omitempty"`
	PayoutExtraId    string  `json:"payoutExtraId,omitempty"`
	FromCurrency     string  `json:"fromCurrency,omitempty"`
	ToCurrency       string  `json:"toCurrency,omitempty"`
	RefundAddress    string  `json:"refundAddress,omitempty"`
	RefundExtraId    string  `json:"refundExtraId,omitempty"`
	Id               string  `json:"id,omitempty"`
	PayinExtraIdName string  `json:"payinExtraIdName,omitempty"`
	FromNetwork      string  `json:"fromNetwork,omitempty"`
	ToNetwork        string  `json:"toNetwork,omitempty"`
}

type reqBody struct {
	FromCurrency  string `json:"fromCurrency,omitempty"`
	ToCurrency    string `json:"toCurrency,omitempty"`
	FromNetwork   string `json:"fromNetwork,omitempty"`
	ToNetwork     string `json:"toNetwork,omitempty"`
	FromAmount    string `json:"fromAmount,omitempty"`
	ToAmount      string `json:"toAmount,omitempty"`
	Address       string `json:"address,omitempty"`
	ExtraId       string `json:"extraId,omitempty"`
	RefundAddress string `json:"refundAddress,omitempty"`
	RefundExtraId string `json:"refundExtraId,omitempty"`
	UserId        string `json:"userId,omitempty"`
	Payload       string `json:"payload,omitempty"`
	ContactEmail  string `json:"contactEmail,omitempty"`
	Source        string `json:"source,omitempty"`
	Flow          string `json:"flow,omitempty"`
	Type          string `json:"type,omitempty"`
	RateId        string `json:"rateId,omitempty"`
}

type OrderStatus struct {
	Id                 string      `json:"id,omitempty"`
	Status             string      `json:"status,omitempty"`
	ActionsAvailable   bool        `json:"actionsAvailable,omitempty"`
	FromCurrency       string      `json:"fromCurrency,omitempty"`
	FromNetwork        string      `json:"fromNetwork,omitempty"`
	ToCurrency         string      `json:"toCurrency,omitempty"`
	ToNetwork          string      `json:"toNetwork,omitempty"`
	ExpectedAmountFrom interface{} `json:"expectedAmountFrom,omitempty"`
	ExpectedAmountTo   float64     `json:"expectedAmountTo,omitempty"`
	AmountFrom         int         `json:"amountFrom,omitempty"`
	AmountTo           float64     `json:"amountTo,omitempty"`
	PayinAddress       string      `json:"payinAddress,omitempty"`
	PayoutAddress      string      `json:"payoutAddress,omitempty"`
	PayinExtraId       interface{} `json:"payinExtraId,omitempty"`
	PayoutExtraId      interface{} `json:"payoutExtraId,omitempty"`
	RefundAddress      interface{} `json:"refundAddress,omitempty"`
	RefundExtraId      interface{} `json:"refundExtraId,omitempty"`
	CreatedAt          time.Time   `json:"createdAt,omitempty"`
	UpdatedAt          time.Time   `json:"updatedAt,omitempty"`
	ValidUntil         interface{} `json:"validUntil,omitempty"`
	DepositReceivedAt  time.Time   `json:"depositReceivedAt,omitempty"`
	PayinHash          string      `json:"payinHash,omitempty"`
	PayoutHash         string      `json:"payoutHash,omitempty"`
	FromLegacyTicker   string      `json:"fromLegacyTicker,omitempty"`
	ToLegacyTicker     string      `json:"toLegacyTicker,omitempty"`
	RefundHash         interface{} `json:"refundHash,omitempty"`
	RefundAmount       interface{} `json:"refundAmount,omitempty"`
}

type EstimateFeeResp struct {
	FromCurrency             string      `json:"fromCurrency,omitempty"`
	FromNetwork              string      `json:"fromNetwork,omitempty"`
	ToCurrency               string      `json:"toCurrency,omitempty"`
	ToNetwork                string      `json:"toNetwork,omitempty"`
	Flow                     string      `json:"flow,omitempty"`
	Type                     string      `json:"type,omitempty"`
	RateId                   interface{} `json:"rateId,omitempty"`
	ValidUntil               interface{} `json:"validUntil,omitempty"`
	TransactionSpeedForecast string      `json:"transactionSpeedForecast,omitempty"`
	WarningMessage           interface{} `json:"warningMessage,omitempty"`
	DepositFee               float64     `json:"depositFee,omitempty"`
	WithdrawalFee            float64     `json:"withdrawalFee,omitempty"`
	UserId                   interface{} `json:"userId,omitempty"`
	FromAmount               int         `json:"fromAmount,omitempty"`
	ToAmount                 float64     `json:"toAmount,omitempty"`
}
