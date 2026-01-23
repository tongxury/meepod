package nodereal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"io"
	"net/http"
)

func NewClient() *Client {
	return &Client{
		Entrypoint: "https://open-platform.nodereal.io/26aa3d8814b64e869003388314286d25/pancakeswap-free/graphql",
	}
}

type Client struct {
	Entrypoint string
}

func (t *Client) ListPairs(ctx context.Context, block int64, page, size int64) {

	params := map[string]interface{}{
		//"createdAtTimestampStart": timestampStart,
		"block": conv.String(block),
		"first": size,
		"skip":  page * size,
	}

	sql := `
		 query pairs($block: Block_height, $first: Int!, $skip: Int!) {
				pairs(first: $first, skip: $skip, block: $block) {
					  id
					  token0 {
						id
						symbol
						name
						decimals
					  }
					  token1 {
						id
						symbol
						name
						decimals
					  }
					createdAtTimestamp
				}
			  }
	`

	var dst struct {
		Data struct {
			Pairs APairs
		}
	}

	err := t.execute(ctx, t.Entrypoint, sql, params, &dst)
	if err != nil {
		return
	}
}

func (t *Client) execute(ctx context.Context, url, sql string, params map[string]interface{}, dst any) error {

	reqBody := map[string]interface{}{
		//"operationName": "pairs",
		"variables": params,
		"query":     sql,
	}

	bodyBytes := bytes.NewBuffer([]byte(conv.M2J(reqBody)))
	req, err := http.NewRequest("POST", url, bodyBytes)
	if err != nil {
		return xerror.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return xerror.Wrap(err)
	}

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return xerror.Wrap(err)
	}

	if resp.StatusCode != http.StatusOK {
		return xerror.Wrapf(string(bytesBody))
	}

	var errResp errorResponse
	if err := json.Unmarshal(bytesBody, &errResp); err != nil {
		return err
	}

	if len(errResp.Errors) > 0 {
		return errors.New(errResp.Errors[0].Message)
	}

	//defer resp.Body.Close()

	if err := json.Unmarshal(bytesBody, dst); err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

type errorResponse struct {
	Errors []struct {
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Message string `json:"message"`
	} `json:"errors"`
}
