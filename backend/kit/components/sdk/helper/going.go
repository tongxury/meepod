package helper

import (
	"context"
	"sync"
)

func Going(ctx context.Context, fs ...func() error) []error {

	var errs []error

	wg := sync.WaitGroup{}
	l := sync.Mutex{}

	for _, f := range fs {

		wg.Add(1)

		fn := f
		go func(ctx context.Context) {
			defer DeferFunc(func() {
				wg.Done()
			})

			if err := fn(); err != nil {
				l.Lock()
				errs = append(errs, err)
				l.Unlock()
			}
		}(ctx)
	}

	wg.Wait()

	return errs
}
