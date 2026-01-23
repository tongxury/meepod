package errorx

import (
	"fmt"
)

type Error struct {
	Message string
	Err     string
	Code    int
}

func (t Error) Error() string {
	return fmt.Sprintf("illegal error: %s", t.Err)
}

func ParamErrorf(format string, a ...any) Error {
	return Error{Err: fmt.Sprintf(format, a...), Code: 10400}
}
func ParamError(err error) Error {
	if IsMyErr(err) {
		return err.(Error)
	}
	return Error{Err: err.Error(), Code: 10400}
}

func ServiceErrorf(format string, a ...any) Error {
	return Error{Message: fmt.Sprintf(format, a...), Code: 10501}
}

func ServerError(err error) Error {

	if IsMyErr(err) {
		return err.(Error)
	}

	return Error{Err: err.Error(), Code: 10500}
}

func UserMessage(message string) Error {
	return Error{Message: message, Code: 10000, Err: message}
}

func IsMyErr(err error) bool {

	switch err.(type) {
	case Error:
		return true
	default:
		return false
	}
}
