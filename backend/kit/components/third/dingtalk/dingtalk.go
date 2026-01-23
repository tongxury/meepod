package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"io"
	"net/http"

	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
)

type ActionButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type ActionButtons []ActionButton

func (a ActionButtons) ToMap() int {
	return len(a)
}

type params struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

type ActionCard struct {
	Title          string        `json:"title"`
	Text           string        `json:"text"`
	BtnOrientation string        `json:"btnOrientation"`
	Btns           ActionButtons `json:"btns"`
}

// "msgtype": "actionCard",
// 		"actionCard": map[string]any{
// 			"title":          title,
// 			"text":           "【打新提醒】![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
// 			"btnOrientation": "0",
// 			"btns": []map[string]any{
// 				{
// 					"title":     "内容不错",
// 					"actionURL": "https://www.dingtalk.com/",
// 				},
// 				{
// 					"title":     "不感兴趣",
// 					"actionURL": "https://www.dingtalk.com/",
// 				},
// 			},
// 		},

func SendActionCardMessage(ctx context.Context, keyword string, message string, actions ActionButtons, accessToken string) error {

	// accessToken := "c2b846dd8dba094a4116164cec51fe68652ee3b96aa2a2c8c2d5b4ecd78bfdd3"
	//accessToken := "4450ef0ddf2eb36e7ead20a54739d61f44728f48e32f3d9d0890b514ed16464c"

	URL := "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken

	p := params{
		Msgtype: "actionCard",
		ActionCard: ActionCard{
			Title:          "",
			Text:           fmt.Sprintf("【%s】%s", keyword, message),
			BtnOrientation: "0",
			Btns:           actions,
		},
	}

	bodyBytes := bytes.NewBuffer([]byte(conv.S2J(p)))
	req, err := http.NewRequest("POST", URL, bodyBytes)
	if err != nil {
		return xerror.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return xerror.Wrap(err)
	}
	defer resp.Body.Close()

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return xerror.Wrap(err)
	}

	var mapResult map[string]any
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal(bytesBody, &mapResult); err != nil {
		return xerror.Wrap(err)
	}

	// {
	// 	"errcode": 0,
	// 	"errmsg": "ok"
	//   }
	if mapResult["errmsg"] != "ok" {
		slf.WithError(err).Errorln(mapResult)
		return xerror.Wrapf("response code %v", mapResult["errcode"])
	}

	return nil
}

type textParams struct {
	Text    textMsg `json:"text"`
	MsgType string  `json:"msgtype"`
}

type markdownParams struct {
	MsgType  string      `json:"msgtype"`
	Markdown markdownMsg `json:"markdown"`
}

type markdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type textMsg struct {
	Content string `json:"content"`
}

func SendMarkdownMessage(ctx context.Context, keyword string, message string, accessToken string) error {
	//c2b846dd8dba094a4116164cec51fe68652ee3b96aa2a2c8c2d5b4ecd78bfdd3
	URL := "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken

	p := markdownParams{
		MsgType: "markdown",
		Markdown: markdownMsg{
			Title: keyword,
			Text:  message,
		},
	}

	bytesBody, code, err := httpcli.Client().POST(ctx, URL, p)
	if err != nil {
		return xerror.Wrap(err)
	}

	if code != http.StatusOK {
		return xerror.Wrapf("response code %d", code)
	}

	var mapResult map[string]any
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal(bytesBody, &mapResult); err != nil {
		return xerror.Wrap(err)
	}

	// {
	// 	"errcode": 0,
	// 	"errmsg": "ok"
	//   }
	if mapResult["errmsg"] != "ok" {
		slf.WithError(err).Errorln(mapResult)
		return xerror.Wrapf("response code %v", mapResult["errcode"])
	}

	return nil
}

func SendTextMessage(ctx context.Context, keyword string, message string, accessToken string) error {

	// accessToken := "c2b846dd8dba094a4116164cec51fe68652ee3b96aa2a2c8c2d5b4ecd78bfdd3"
	//accessToken := "4450ef0ddf2eb36e7ead20a54739d61f44728f48e32f3d9d0890b514ed16464c"

	URL := "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken

	p := textParams{
		MsgType: "text",
		Text: textMsg{
			Content: fmt.Sprintf("【%s】%s", keyword, message),
		},
	}

	bodyBytes := bytes.NewBuffer([]byte(conv.S2J(p)))
	req, err := http.NewRequest("POST", URL, bodyBytes)
	if err != nil {
		return xerror.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return xerror.Wrap(err)
	}
	defer resp.Body.Close()

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return xerror.Wrap(err)
	}

	var mapResult map[string]any
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal(bytesBody, &mapResult); err != nil {
		return xerror.Wrap(err)
	}

	// {
	// 	"errcode": 0,
	// 	"errmsg": "ok"
	//   }
	if mapResult["errmsg"] != "ok" {
		slf.WithError(err).Errorln(mapResult)
		return xerror.Wrapf("response code %v", mapResult["errcode"])
	}

	return nil
}
