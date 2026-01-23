package enum

var (
	PayMethod_Account = Enum{Value: "account", Name: "账本"}
	PayMethod_Alipay  = Enum{Value: "alipay", Name: "支付宝", Color: "#02a6e7"}
	PayMethod_Wechat  = Enum{Value: "wechat", Name: "微信", Color: "#31a606"}
	//PayMethod_QrCode  = Enum{Value: "qrCode", Name: "扫码"}
	//PayMethod_AlipayQrCode = Enum{Value: "alipayQrCode", Name: "支付宝扫码"}
	//PayMethod_WechatQrCode = Enum{Value: "wechatQrCode", Name: "微信扫码"}

	AllPayMethods = []Enum{PayMethod_Account, PayMethod_Alipay, PayMethod_Wechat}
)

func PayMethod(value string) Enum {
	for _, x := range AllPayMethods {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
