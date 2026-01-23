package alipay

type Options struct {
	Entrypoint      string
	AppId           string
	PrivateKey      string
	AlipayPublicKey string
	Income          IncomeOptions
}

type IncomeOptions struct {
	Pid           string
	Name          string
	Memo          string
	LoginName     string
	BindLoginName string
	Rate          float64
}
