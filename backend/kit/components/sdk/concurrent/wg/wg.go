package wg

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"sync"
)

func Run[P any, T any](ctx context.Context, params []P, fun func(ctx context.Context, p P) ([]T, error)) ([]T, []error) {

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	var errs []error
	var rsp []T

	for _, param := range params {
		wg.Add(1)
		go func(ctx context.Context, p P) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			ts, err := fun(ctx, p)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp = append(rsp, ts...)
			lc.Unlock()

		}(ctx, param)
	}

	wg.Wait()

	return rsp, errs
}
