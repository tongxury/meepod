package types

type KeyValue[T comparable] struct {
	Key   string `json:"key"`
	Value T      `json:"value"`
}

type KeyValues[T comparable] []*KeyValue[T]

func (ks KeyValues[T]) AsMap() map[string]T {
	rsp := make(map[string]T, len(ks))

	for _, k := range ks {
		rsp[k.Key] = k.Value
	}
	return rsp
}

// KeyFloatValues float
//
//type KeyValueFloat64 = KeyValue[float64]
//
//type KeyValueFloat64s = KeyValues[float64]
