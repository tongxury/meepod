package pgd

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"strings"
)

func SetJSONField(q *pg.Query, field, jField string, jValue any) *pg.Query {

	switch t := jValue.(type) {
	case string:
		jValue = fmt.Sprintf("\"%v\"", t)
	case []string:
		var parts []string
		for _, x := range t {
			parts = append(parts, fmt.Sprintf("\"%v\"", x))
		}
		jValue = "[" + strings.Join(parts, ",") + "]"
	default:
		jValue = fmt.Sprintf("%v", t)
	}

	return q.Set(field+" = jsonb_set("+field+", ?, ?)", fmt.Sprintf("{%s}", jField), jValue)
}
