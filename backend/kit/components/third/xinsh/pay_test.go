package xinsh

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_GenerateTradeQrCode(t *testing.T) {

	ctx := context.Background()

	c := NewXinShPayClient(Config{})

	//c.uploadPicture(ctx, "1", "https://eimg.oss-cn-beijing.aliyuncs.com/pl3.png")

	info, _, err := c.GenerateTradeQrCode(ctx, TradeParams{
		//MerchantNo:   "HPX2023091100023",
		MerchantNo:   "HPX2023091600005",
		OrderId:      "211595",
		Amount:       1,
		Subject:      "商品",
		UserClientIp: "35.23.123.21",
		TimeExpire:   15,
	})

	fmt.Println(info, err)
}
