package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"time"
)

type CoStore struct {
	Id           string    `json:"id"`
	Store        *Store    `json:"store"`
	CoStore      *Store    `json:"co_store"`
	Items        CoItems   `json:"items"`
	CreatedAt    string    `json:"created_at"`
	CreatedAtTs  int64     `json:"created_at_ts"`
	Status       enum.Enum `json:"status"`
	Tags         Tags      `json:"tags"`
	Updatable    bool      `json:"updatable"`
	TopUpable    bool      `json:"topupable"`
	Pausable     bool      `json:"pausable"`
	Resumable    bool      `json:"resumable"`
	EndApplyable bool      `json:"endapplyable"`
	Endable      bool      `json:"endable"`
	Recoverable  bool      `json:"recoverable"`
	Balance      float64   `json:"balance"`
}

type CoItem struct {
	Item *Item   `json:"item"`
	Rate float64 `json:"rate"`
}

type CoItems []*CoItem

type CoStores []*CoStore

func FromDbCoStore(x *db.CoStore, store *db.Store, storeOwner *db.User, coStore *db.Store, coStoreOwner *db.User, extra ...db.ExtraMaps) *CoStore {

	if x == nil {
		return nil
	}

	rsp := CoStore{
		Id:          x.Id,
		Store:       FromDbStore(store, storeOwner, extra...),
		CoStore:     FromDbStore(coStore, coStoreOwner, extra...),
		Items:       nil,
		CreatedAt:   x.CreatedAt.Format(time.DateOnly),
		CreatedAtTs: x.CreatedAt.Unix(),
		Status:      enum.CoStoreStatus(x.Status),
		//MStatus:      enum.ProxyStatus(x.MStatus),
		Tags:         nil,
		Updatable:    helper.InSlice(x.Status, enum.UpdatableCoStoreStatus),
		TopUpable:    helper.InSlice(x.Status, enum.TopupableCoStoreStatus),
		Pausable:     helper.InSlice(x.Status, enum.PausableCoStoreStatus),
		Resumable:    helper.InSlice(x.Status, enum.ResumableCoStoreStatus),
		EndApplyable: helper.InSlice(x.Status, enum.EndApplyableCoStoreStatus),
		Endable:      helper.InSlice(x.Status, enum.EndableCoStoreStatus),
		Recoverable:  helper.InSlice(x.Status, enum.RecoverableCoStoreStatus),
		Balance:      0,
	}

	if len(extra) > 0 {

		for k, v := range x.Items {

			rsp.Items = append(rsp.Items, &CoItem{
				Item: FromDbItem(extra[0].Items[k]),
				Rate: v,
			})
		}

	}

	return &rsp
}
