package rest

import (
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/dlt"
	"gitee.com/meepo/backend/shop/core/issue/adapter/fc/ssq"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func getFilters() gin.H {

	var month types.OptionItems

	max := 3

	month = append(month, types.OptionItem{
		Label: fmt.Sprintf("全部(近%d月)", max),
		Value: "",
	})

	y, m, _ := time.Now().In(timed.LocAsiaShanghai).Date()
	for i := 0; i < max; i++ {

		xy := y
		xm := int(m) - i
		if xm <= 0 {
			xy -= 1
			xm = 12 - xm
		}

		month = append(month, types.OptionItem{
			Label: fmt.Sprintf("%d", xm) + "月",
			Value: fmt.Sprintf("%d-%s", xy, fmt.Sprintf(helper.Choose(xm >= 10, "%d", "0%d"), xm)),
		})
	}

	coStorePaymentCats := types.OptionItems{
		{Label: "全部类型", Value: ""},
	}

	for _, x := range enum.AllCoStorePaymentCategories {
		coStorePaymentCats = append(coStorePaymentCats, types.OptionItem{
			Label: x.Name, Value: x.Value,
		})
	}

	proxyRewardStatus := types.OptionItems{
		{Label: "全部类型", Value: ""},
	}

	for _, x := range enum.AllProxyRewardStatus {
		proxyRewardStatus = append(proxyRewardStatus, types.OptionItem{
			Label: x.Name, Value: x.Value,
		})
	}

	return gin.H{
		"month":              month,
		"proxyRewardCats":    proxyRewardStatus,
		"coStorePaymentCats": coStorePaymentCats,
	}
}

//
//func GetUpdates(c *gin.Context) {
//	uVersion := c.GetHeader("U-Version")
//
//	gind.OK(c, gin.H{
//		"forceUpdate": uVersion != "1.3.2",
//	})
//}

func GetSettings(c *gin.Context) {

	client := c.GetHeader("Client")

	var settings any

	switch client {
	case "keeper":

		settings = map[string]any{
			"showCoStore": enum.ShowConStore,
			"rejectReasons": []struct {
				Id   string `json:"id"`
				Text string `json:"text"`
			}{
				{"1", "联系不上客户"},
				{"2", "不支持此彩种"},
				{"3", "过期无法出票"},
				{"4", "其他原因"},
			},
			"rejectWithdrawReasons": []string{"用户异常", "资金来源不明"},
			"maxUploadTicket":       3,
			"service": map[string]any{
				"wechat": "https://eimg.oss-cn-beijing.aliyuncs.com/mmqrcode1691251775557.png",
				"email":  "lottery@example.com",
			},
			"filters": getFilters(),
			"proxy": gin.H{
				"maxRewardRate": 0.07,
				"rewardRateItems": types.OptionItems{
					{Label: "1%", Value: "0.01"},
					{Label: "2%", Value: "0.02"},
					{Label: "3%", Value: "0.03"},
					{Label: "3.5%", Value: "0.035"},
					{Label: "4%", Value: "0.04"},
					{Label: "4.5%", Value: "0.045"},
					{Label: "5%", Value: "0.05"},
					{Label: "5.5%", Value: "0.055"},
					{Label: "6%", Value: "0.06"},
					{Label: "6.5%", Value: "0.065"},
					{Label: "7%", Value: "0.07"},
				},
			},
		}

	case "user":
		settings = map[string]any{
			"ssqBlueLimit":   ssq.MaxBluePerTicket,
			"ssqRedLimit":    ssq.MaxRedPerTicket,
			"ssqDRedLimit":   ssq.MaxDRedPerTicket,
			"dltBlueLimit":   dlt.MaxBluePerTicket,
			"dltRedLimit":    dlt.MaxRedPerTicket,
			"dltDRedLimit":   dlt.MaxDRedPerTicket,
			"rx9DanLimit":    5,
			"minUnionAmount": 30,
			"maxVolume":      150,
			"maxMultiple":    50000,
		}

	}

	gind.OK(c, settings)
}
