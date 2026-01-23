package cct

func Go(f func()) {
	go func() {
		defer DeferFunc()
		f()
	}()
}
