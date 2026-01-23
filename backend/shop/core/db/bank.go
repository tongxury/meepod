package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
)

type Bank struct {
	tableName    struct{} `pg:"t_banks"`
	BankId       string   `json:"bank_id"`
	BankCode     string   `json:"bank_code"`
	BankName     string   `json:"bank_name"`
	BranchName   string   `json:"branch_name"`
	BranchNo     string   `json:"branch_no"`
	ProvinceName string   `json:"province_name"`
	CityName     string   `json:"city_name"`
	ProvinceId   string   `json:"province_id"`
	CityId       string   `json:"city_id"`
}

type Banks []*Bank

type ListBankParams struct {
	BankId         string
	BranchNo       string
	BranchNameLike string
	Page, Size     int64
}

func (t *Bank) RequireByBranchNo(ctx context.Context, branchNo string) (*Bank, error) {

	var tmp Banks
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("branch_no = ?", branchNo).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no branchNo found: %s", branchNo)
	}

	return tmp[0], nil
}

func (t *Bank) List(ctx context.Context, params ListBankParams) (Banks, int64, error) {

	var tmp Banks
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.BankId != "" {
		q = q.Where("bank_id = ?", params.BankId)
	}
	if params.BranchNo != "" {
		q = q.Where("branch_no = ?", params.BranchNo)
	}

	if params.BranchNameLike != "" {
		q = q.Where("branch_name like ?", "%"+params.BranchNameLike+"%")
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	count, err := q.SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}
