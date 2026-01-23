package conv

func AnySlice[T any](values []T) (rsp []any) {
	for _, v := range values {
		rsp = append(rsp, v)
	}

	return
}
