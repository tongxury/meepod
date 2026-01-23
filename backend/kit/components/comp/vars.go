package comp

import "fmt"

var ins = Var{mp: map[string]any{}}

func Vars() *Var {
	return &ins
}

type Var struct {
	mp map[string]any
}

func (v *Var) Set(key string, val any) *Var {
	if v.mp == nil {
		v.mp = map[string]any{}
	}
	v.mp[key] = val
	return v
}

func (v *Var) Find(key string) (any, bool) {
	val, found := v.mp[key]
	return val, found
}

func (v *Var) RequireString(key string) string {
	val := v.GetString(key, "")
	if val != "" {
		return val
	}
	panic(fmt.Errorf("no vars found by keyt: %s", key))
}

func (v *Var) GetString(key string, defVal string) string {
	val, found := v.mp[key]
	if found {
		str, ok := val.(string)
		if ok {
			return str
		}
	}

	return defVal

}
