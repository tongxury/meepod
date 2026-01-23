package db

import "gitee.com/meepo/backend/kit/components/sdk/conv"

func FindInExtra(extra string, key string) (string, error) {
	if extra == "" || extra == "{}" {
		return "", nil
	}

	var mp map[string]any

	err := conv.J2M(extra, &mp)
	if err != nil {
		return "", err
	}

	if val, found := mp[key]; found {
		return conv.String(val), nil
	}

	return "", nil
}

type ExtraMaps struct {
	Items map[string]*Item
}
