package helper

func AppendMap[K comparable, V any](dst map[K]V, src ...map[K]V) map[K]V {
	if dst == nil {
		dst = map[K]V{}
	}

	for _, m := range src {
		for k, v := range m {
			dst[k] = v
		}
	}

	return dst
}

func CreateStringMap(params ...string) map[string]string {
	rsp := make(map[string]string)

	for index := 0; index < len(params)-1; index = index + 2 {
		if index%2 == 0 {
			rsp[params[index]] = params[index+1]
		}
	}

	return rsp
}

func HasKey(m map[string]interface{}) bool {
	return len(Keys(m)) > 0
}

func ContainsKey(m map[string]interface{}, key string) bool {
	if m == nil {
		return false
	}

	keys := Keys(m)

	return Contains(keys, key)
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// func Keys(m map[string]interface{}) []string {
// 	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
// 	keys := make([]string, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }

func MapValues[T any](m map[string]T) []T {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	values := make([]T, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
