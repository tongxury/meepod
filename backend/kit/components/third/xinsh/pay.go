package xinsh

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/encryptor"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"github.com/go-resty/resty/v2"
	"github.com/lithammer/shortuuid/v4"
	"net/url"
	"time"
)

type Client struct {
	conf Config
}

//测试环境域名：
//进件：https://gateway-hpxtest1.hnapay.com/merchant
//支付：https://gateway-hpxtest1.hnapay.com/order
//
//生产环境域名：
//进件：https://gateway-hpx.hnapay.com/merchant
//支付：https://gateway-hpx.hnapay.com/order

func NewXinShPayClient(conf Config) *Client {
	//conf.MerchantEntrypoint = "https://gateway-hpxtest1.hnapay.com/merchant"
	//conf.OrderEntrypoint = "https://gateway-hpxtest1.hnapay.com/order"

	conf.MerchantEntrypoint = "https://gateway-hpx.hnapay.com/merchant"
	conf.OrderEntrypoint = "https://gateway-hpx.hnapay.com/order"
	conf.AppId = "202309110002"

	conf.PrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCxFMsXcEzsBcNbVNK3jFuQSWKAIEDKIqNoAqElZ+2SLgzenBMPp5HUhO6QulxZZVlseUcD7+z+nhYNOJ2jn/K76zsj38nBl/+vEc4mrJ4ePwzhGIK00+455J+X4nrhQpd3BtEEtxuUNGJi8ZJOvHedyRgo637vKnfqGRO/1sxCAgcuG3hh/I9Cf9yrS9JGbyw0GL3PovlUe4uM8V72R+NrkfO+MCsx2ma7hjcHSni1qROgExXnqynlZNmbzRvDSoQwD/Y22LjrE9ErJjaJkOtbKXGz+sFEQR5kv/5NfP06oWY115T7DemE3j56bAVvtn29+o/DSLKEOGHgb4RsrbaHAgMBAAECggEARKDRYOEV/rbxEllaXOZZqh62vP9C/NPVzn6OY4fWq7uLI08LmBqSSvHF33NKTR8ZPA/4dM8sYzVzSGZzubFIionM3sdhUtUMs5XL9nMZyJEY9bfhGCG2httS+rM8ekarmuaaZSBt4M84fi4vTLKADTedVGaISDqGRMS5QbUGbmHhR1JGJvz1UdeCx4pQvolb49e0WW/xRCydRP7SD2sB1Co3m1I46b6e27RmVG72Ei8RTONWjklo1p3MTWjt93UO0QAU19TZ7mV97FmogXoTkR/vNmXYpY3Oarpp2MFGeQloDxopM1u418GIgBpkls1+wP0qiymHqb2QQpiJQcXa+QKBgQD9fsERpLuFkg+xa1Hn+LzFrOpWwn5soNfBreLDa/es1pyYvLfdAB3mhg8BYq8wCHIVvjzbPwuZW9vxAkG9F/FNOmVEkeCJv0c4ylhfjjyHJUMldcUs04AXFKYUnqi3zGlmxq4+SZs29ktVowFB0dqW7Es9cCS9RSBt3kkVHAijVQKBgQCy1L28jOc/8n0LN8fre/gaf8Rgz5i8HOCU2VI1Mb9Nj4FREoaLfvURpcPKXMJyQv1IplAm6EiR/5sAX9bJxsN0w6pK1o1EZS9/oqzppZ7yRiHNOO1fWX8zt1fxsbNUjmHqfBIVQRECKs6I/hd+T34rjc17qoQIugleXzZ+DC+qawKBgEoyBXSSkMhhnfJCBTEuXqJFIDnQp4xH756itJKaUV3nWuJhqjcnR5knd9Dh/4DBmBLBIbLSWyTB/DgofvFHxrrh8q4FPIFU2RXIM+GUEidEQsj+FX4vUXhv9MRoQ924GMmaMXnNtX56zEX+dem78IzoEIWlAzvatckynJVvJSAZAoGBAJYfUUg/YMGl5qLMKN/eKeDU8R30J9uCwmUyKnjNUjLiDUTpsFjaMxClfz8Zr4IMCmQ6eX6v8Hvff1dJHR415U7YEtTQV7ba2ozjGxTpHA12IloNN/ebQdVATGtxKYIKJNibXvLItGaFWOxXoZE/WkNlvuHZuw04XzUB6NTXlgwRAoGBAJZBAa7MwERQy+DF/0xLZ9ENK/Zq8fHxqToEWnep4dxB2RjiCek5LxXOiYBT/idgrkiymrBn3iBXjyy4zB2Omqt/XzRoI51i5ZnP565h8Hs81gBdWqqYR+VgSoPYH4ZFPFtwT9yktXZamudn+NyCXatRf2eiEdakQ7Vescr+yFcJ"
	// public MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsRTLF3BM7AXDW1TSt4xbkEligCBAyiKjaAKhJWftki4M3pwTD6eR1ITukLpcWWVZbHlHA+/s/p4WDTido5/yu+s7I9/JwZf/rxHOJqyeHj8M4RiCtNPuOeSfl+J64UKXdwbRBLcblDRiYvGSTrx3nckYKOt+7yp36hkTv9bMQgIHLht4YfyPQn/cq0vSRm8sNBi9z6L5VHuLjPFe9kfja5HzvjArMdpmu4Y3B0p4takToBMV56sp5WTZm80bw0qEMA/2Nti46xPRKyY2iZDrWylxs/rBREEeZL/+TXz9OqFmNdeU+w3phN4+emwFb7Z9vfqPw0iyhDhh4G+EbK22hwIDAQAB
	conf.OriginPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmw7/GAmnsqrumfIJs9sW5ProNDewuRNSmqyQUShafZ1risfmaxvVEjmZoohJHbGLH8KY56N+MBAEjhnbk/s2mNLG6L9+ag+ymQn6KFJwTwCwtOSlVnOEW5nx6KniSkZJygjZ3Hu2rfDtm+CyT577AkY6iDFzCtWuPmoT/jFSTYQuObbSmPoEvHVdMVBcS3X+uD6yB3zEGddzJKjucaViMeG57ACqtgeBsfQK2hKrlzkpCylAOagpvQfUAlaLLqC0WIfNb5enZ/JXUNcZbDMBR5Ai3X7iiv6i53tdMwiUhuk3zOcyDdU61D1JyICWMlFTA+jwlTY6ZJcE+g4BnG0ZrQIDAQAB"

	return &Client{
		conf: conf,
	}
}

func (t *Client) genRequest(reqData map[string]any) map[string]any {

	reqId := conv.String(time.Now().UnixNano())
	ts := conv.String(time.Now().UnixMilli())

	req := map[string]any{
		"reqId":     reqId,
		"orgNo":     t.conf.AppId,
		"reqData":   reqData,
		"signType":  "RSA",
		"timestamp": conv.String(time.Now().UnixMilli()),
		"version":   "1.0",
	}

	toSignString := fmt.Sprintf("orgNo=%s&reqData=%s&reqId=%s&signType=RSA&timestamp=%s&version=1.0",
		t.conf.AppId, conv.M2J(reqData), reqId, ts)

	sign, _ := encryptor.RSASignWithRSA1(toSignString, t.conf.PrivateKey)

	req["sign"] = sign

	return req
}

func (t *Client) SaveMerchant(ctx context.Context, params MerchantParams, isUpdate bool) (*MerchantApply, error) {

	storeFront, err := t.uploadPicture(ctx, "8", params.StoreFront)
	if err != nil {
		return nil, err
	}
	storeInSide, err := t.uploadPicture(ctx, "9", params.StoreInSide)
	if err != nil {
		return nil, err
	}

	idCardFront, err := t.uploadPicture(ctx, "1", params.IdCardFront)
	if err != nil {
		return nil, err
	}
	idCardBack, err := t.uploadPicture(ctx, "2", params.IdCardBack)
	if err != nil {
		return nil, err
	}

	idCardFrontPay, err := t.uploadPicture(ctx, "21", params.IdCardFront)
	if err != nil {
		return nil, err
	}
	idCardBackPay, err := t.uploadPicture(ctx, "22", params.IdCardBack)
	if err != nil {
		return nil, err
	}

	idCardHandled, err := t.uploadPicture(ctx, "23", params.IdCardHandled)
	if err != nil {
		return nil, err
	}
	bankCardFront, err := t.uploadPicture(ctx, "16", params.BankCardFront)
	if err != nil {
		return nil, err
	}
	protocol, err := t.uploadPicture(ctx, "14", "https://eimg.oss-cn-beijing.aliyuncs.com/mmexport1694435706158.jpg")
	if err != nil {
		return nil, err
	}

	//p, c, d := t.mappingLoc(params.Province, params.City, params.District)
	//bp, bc, _ := t.mappingLoc(params.BankProvince, params.BankCity, "")

	reqData := t.genRequest(map[string]any{
		"requestId":                 params.RequestId,
		"merchantNm":                params.MerchantNm,
		"merchantShortNm":           params.MerchantNm,
		"merchantType":              "01",   // 01 自然人 02 个体户 03 企业
		"mcc":                       "5331", // 5331	生活百货	百货商城	杂货店
		"province":                  params.Province,
		"city":                      params.City,
		"district":                  params.District,
		"address":                   params.Address,
		"linkman":                   params.MerchantNm,
		"phone":                     params.Phone,
		"email":                     params.Email,
		"customerPhone":             params.Phone,
		"principal":                 params.Username,
		"principalIdcodeType":       "0",
		"principalIdcode":           params.IdCardNo,
		"legalPersonCertificateStt": params.IdCardFrom,
		"legalPersonCertificateEnt": params.IdCardTo,
		"protocolPhoto":             protocol.PictureName,
		"lawyerCertFrontPhoto":      idCardFront.PictureName,
		"lawyerCertBackPhoto":       idCardBack.PictureName,
		"mainPhoto":                 storeFront.PictureName,
		"storeHallPhoto":            storeInSide.PictureName,
		"accountNo":                 params.BankAccount,
		"accountNm":                 params.BankAccountName,
		"accountType":               params.AccountType,
		"accountNature":             params.AccountNature,
		"idcardType":                "1",
		"idcardNo":                  params.IdCardNo,
		"authorizedCertFrontPhoto":  idCardFrontPay.PictureName,
		"authorizedCertBackPhoto":   idCardBackPay.PictureName,
		"holdIdentityPic":           idCardHandled.PictureName,
		"bankCardFrontPhoto":        bankCardFront.PictureName,
		"validateDateStart":         params.IdCardFrom,
		"validateDateExpired":       params.IdCardTo,
		"cnapsCode":                 params.CnapsCode,
		"bankId":                    params.BankId,
		"bankName":                  params.BankName,
		"identityPhone":             params.BankPhone,
		"bankProvince":              params.BankProvince,
		"bankCity":                  params.BankCity,
		"settleWay":                 "02", // 01：T1 02：D1 03：D0
		"profitConf": []map[string]any{
			{
				"rateTypeId": "0101", //支付方式
				"channel":    "01",   //支付渠道（01-支付宝，02-微信，03-银联二维码，04-银联刷卡
				"openFlag":   "1",
				"feeRate":    0.0022, //费率
			},
			{
				"rateTypeId": "0201", //支付方式
				"channel":    "02",   //支付渠道（01-支付宝，02-微信，03-银联二维码，04-银联刷卡
				"openFlag":   "1",
				"feeRate":    0.0022, //费率
			},
			{
				"rateTypeId": "0307",
				"channel":    "03",
				"openFlag":   "0",
				"feeRate":    0.0065,
			},
			{
				"rateTypeId": "0308",
				"channel":    "03",
				"openFlag":   "0",
				"feeRate":    0.0065,
			},
			{
				"rateTypeId":  "0401",
				"channel":     "04",
				"openFlag":    "0",
				"feeRate":     0.0065,
				"priceCapped": 2500, //封顶金额，刷卡必填
			},
			{
				"rateTypeId":  "0402",
				"channel":     "04",
				"openFlag":    "0",
				"feeRate":     0.0065,
				"priceCapped": 2500, //封顶金额，刷卡必填
			},
		},
	})

	req := conv.M2J(reqData)

	fmt.Println(req)

	response, err := resty.New().R().SetContext(ctx).
		SetBody(reqData).
		Post(t.conf.MerchantEntrypoint + helper.Choose(isUpdate, "/merchant/patch", "/merchant/apply"))

	var resp Resp[MerchantApply]

	if err := conv.B2S(response.Body(), &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0000" {
		return nil, errors.New(resp.Msg)
	}

	return resp.RespData, nil
}

func (t *Client) GetApplyState(ctx context.Context, requestId string) (*ApplyState, error) {

	reqData := t.genRequest(map[string]any{
		"requestId": requestId,
	})

	response, err := resty.New().R().SetContext(ctx).SetBody(reqData).Post(t.conf.MerchantEntrypoint + "/merchant/queryApplyResult")
	if err != nil {
		return nil, err
	}

	var resp Resp[ApplyState]

	if err := conv.B2S(response.Body(), &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0000" {
		return nil, errors.New(resp.Msg)
	}

	return resp.RespData, nil
}

func (t *Client) uploadPicture(ctx context.Context, pictureType, url string) (*Upload, error) {

	imageResp, err := resty.New().R().SetContext(ctx).Get(url)
	if err != nil {
		return nil, err
	}

	response, err := resty.New().R().SetContext(ctx).
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"orgNo":       t.conf.AppId,
			"pictureType": pictureType,
		}).
		//SetFile("file", "/Users/tongxu/Desktop/favicon.png").
		SetFileReader("file", shortuuid.New()+".png", bytes.NewReader(imageResp.Body())).
		Post(t.conf.MerchantEntrypoint + "/merchant/uploadPicture")

	var resp Resp[Upload]

	if err := conv.B2S(response.Body(), &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0000" {
		return nil, errors.New(resp.Msg)
	}

	return resp.RespData, nil
}

// 文档地址 https://www.yuque.com/zhaohanying/egbhoy/wca6r5ueiaiyga3c
func (t *Client) GenerateTradeQrCode(ctx context.Context, params TradeParams) (*PayInfo, any, error) {

	reqData := t.genRequest(map[string]any{
		"merchantNo": params.MerchantNo,
		"orderNo":    params.OrderId,
		"amt":        params.Amount * 100, // 分
		"payType":    "ALIPAY",
		//"ledgerAccountFlag": "00", // 分账
		"subject":    params.Subject,
		"trmIp":      params.UserClientIp,
		"timeExpire": conv.String(params.TimeExpire),
		"notifyUrl":  url.QueryEscape(params.NotifyUrl),
	})

	response, err := resty.New().R().SetContext(ctx).SetBody(reqData).Post(t.conf.OrderEntrypoint + "/trade/activeScan")
	if err != nil {
		return nil, reqData, err
	}

	var resp Resp[PayInfo]

	if err := conv.B2S(response.Body(), &resp); err != nil {
		return nil, reqData, err
	}

	if resp.Code != "0000" {
		return nil, reqData, errors.New(resp.Msg)
	}

	return resp.RespData, reqData, nil
}

func (t *Client) GetTradeResult(ctx context.Context, merchantNo, orderId string) (*TradeResult, error) {

	reqData := t.genRequest(map[string]any{
		"merchantNo": merchantNo,
		"orderNo":    orderId,
	})

	response, err := resty.New().R().SetContext(ctx).SetBody(reqData).Post(t.conf.OrderEntrypoint + "/trade/tradeQuery")
	if err != nil {
		return nil, err
	}

	var resp Resp[TradeResult]

	if err := conv.B2S(response.Body(), &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0000" {
		return nil, errors.New(resp.Msg)
	}

	return resp.RespData, nil

}
