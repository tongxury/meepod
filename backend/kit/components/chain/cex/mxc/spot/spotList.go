package spotList

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/chain/cex/mxc/config"
	"gitee.com/meepo/backend/kit/components/chain/cex/mxc/utils"
)

// 具体请求配置
// 深度信息
func SpotMarketDepth(jsonParams string) interface{} {
	caseUrl := "/depth"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, _ := utils.PublicGet(requestUrl, jsonParams)
	return response
}

// 现货账户信息
func SpotAccountInfo(jsonParams string) interface{} {
	caseUrl := "/account"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, _ := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// 账户成交历史
func SpotmyTrade(jsonParams string) interface{} {
	caseUrl := "/myTrades"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, _ := utils.PrivateGet(requestUrl, jsonParams)
	return response
}

// 创建子账户
func CreateSub(jsonParams string) interface{} {
	caseUrl := "/sub-account/virtualSubAccount"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, _ := utils.PrivatePost(requestUrl, jsonParams)
	return response
}
