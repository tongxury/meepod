package cct

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"runtime"
)

func DeferFunc(f ...func()) {
	if r := recover(); r != nil {

		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}

		slf.WithError(err).Errorln("[PANIC]...\n" + string(buf))
	}

	if len(f) > 0 {
		f[0]()
	}
}
