package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"strings"
)

func NewEtherScanClient(url string) *EtherScanClient {
	if url == "" {
		url = "https://api.etherscan.io/api?apikey=MV5PGK79G5PEEZM17RGS29791E2XMPD3NS"
	}

	return &EtherScanClient{url: url}
}

type EtherScanClient struct {
	url     string
	options httpcli.Options
}

type ECreation struct {
	Address string `json:"contractAddress"`
	Creator string `json:"contractCreator"`
	TxHash  string `json:"txHash"`
}

type ECreations []*ECreation

type resp struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  ECreations `json:"result"`
}

func (e ECreations) AsMap() map[string]ECreation {
	m := map[string]ECreation{}
	for _, creation := range e {
		m[creation.Address] = *creation
	}

	return m
}

func (e ECreations) TxHashes() []string {
	var rsp []string
	for _, creation := range e {
		rsp = append(rsp, creation.TxHash)
	}

	return rsp
}

func (t *EtherScanClient) RequireContractCreation(ctx context.Context, address string) (*ECreation, error) {
	creations, err := t.FindCreatorAndCreationHash(ctx, []string{address})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if len(creations) == 0 {
		return nil, xerror.Wrapf("no creation found by address: %s", address)
	}

	return creations[0], nil
}

func (t *EtherScanClient) FindCreatorAndCreationHash(ctx context.Context, addresses []string) (ECreations, error) {

	url := t.url + fmt.Sprintf("&module=contract&action=getcontractcreation&contractaddresses=%s",
		strings.Join(addresses, ","))

	tokenBytes, code, err := httpcli.Client().GET(ctx, url, t.options)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("code != 200: %d", code)
	}

	var rsp resp
	err = json.Unmarshal(tokenBytes, &rsp)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), string(tokenBytes))
	}

	return rsp.Result, nil
}
