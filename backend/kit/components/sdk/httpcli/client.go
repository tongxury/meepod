package httpcli

import (
	"bytes"
	"context"
	"encoding/json"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"io"
	"net/http"
)

type httpClient struct {
	client *http.Client
}

func Client() *httpClient {
	return &httpClient{
		client: &http.Client{},
	}
}

type Options struct {
	Headers map[string]string
}

func (h *httpClient) GET(ctx context.Context, url string, options ...Options) ([]byte, int, error) {

	req, err := h.build(ctx, "GET", url, nil, options...)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return h.do(ctx, req)
}

func (h *httpClient) POST(ctx context.Context, url string, body any, options ...Options) ([]byte, int, error) {
	req, err := h.build(ctx, "POST", url, body, options...)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}
	return h.do(ctx, req)
}

func (h *httpClient) build(ctx context.Context, method, url string, body any, options ...Options) (*http.Request, error) {

	var bodyBuf io.Reader

	if body != nil {

		bodyBytes, err := json.Marshal(body)

		if err != nil {
			return nil, xerror.Wrap(err)
		}

		bodyBuf = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, url, bodyBuf)

	req.Header.Set("Content-Type", "application/json")
	if len(options) > 0 {
		for k, v := range options[0].Headers {
			req.Header.Set(k, v)
		}
	}

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return req, nil
}

func (h *httpClient) do(ctx context.Context, req *http.Request) ([]byte, int, error) {
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	defer resp.Body.Close()

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, xerror.Wrapf("%s", string(bytesBody))
	}

	//var rsp T
	//
	//err = json.Unmarshal(bytesBody, &rsp)
	//if err != nil {
	//	return nil, 0, xerror.Wrap(err)
	//}

	return bytesBody, resp.StatusCode, nil
}
