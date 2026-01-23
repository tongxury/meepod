package xerror

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	LcauseBy = 1 << iota
	LdateTime
	Llongfile
)

type serviceErr struct {
	Code    int
	Message string
}

type Error struct {
	Current     error
	serviceErrs []serviceErr
	StackTraces []string
	Errors      []error
	Flag        int
}

func (e Error) ServiceErr() (int, string) {
	l := len(e.serviceErrs)
	if l == 0 {
		return 0, ""
	}

	s := e.serviceErrs[l-1]
	return s.Code, s.Message
}

func (e Error) Error() string {
	var header = ""
	//if len(e.Header) != 0 {
	//	buf, _ := json.MarshalIndent(e.Header, "  ", "  ")
	//	header += fmt.Sprintf("header:\n%s\n", buf)
	//}
	rs := fmt.Sprintf("%s%s", header, e.StackTraceValue())

	return rs
}

func (e Error) StackTrace() string {
	rs := ""
	for _, v := range e.StackTraces {
		rs = rs + v + ""
	}
	return rs
}

func (e Error) StackTraceValue() string {
	//header := make([]string, 0, 10)
	//if e.Flag&LdateTime > 0 {
	//	header = append(header, "HappenAt")
	//}
	//if e.Flag&Llongfile > 0 {
	//	header = append(header, "StackTrace")
	//}
	//if e.Flag&LcauseBy > 0 {
	//	header = append(header, "CauseBy")
	//}
	//headerStr := strings.Join(header, " | ")
	rs := make([]string, 0, len(e.StackTraces)+1)
	// rs = append(rs, headerStr)
	rs = append(rs, e.StackTraces...)
	return strings.Join(rs, "   ==   ")
}

func Wrapf(format string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	return wrap(fmt.Errorf(format, args...), 0, "", file, line)
}

func WrapSCf(code int, message string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)

	err := fmt.Errorf(message, args...)
	return wrap(err, code, message, file, line)

}

func WrapSf(err error, code int, message string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	return wrap(err, code, fmt.Sprintf(message, args...), file, line)
}

func WrapS(err error, code int, message string) error {
	_, file, line, _ := runtime.Caller(1)
	return wrap(err, code, message, file, line)
}
func WrapSC(code int, message string) error {
	_, file, line, _ := runtime.Caller(1)
	err := fmt.Errorf(message)
	return wrap(err, code, message, file, line)
}

func Wrap(e error) error {
	_, file, line, _ := runtime.Caller(1)
	return wrap(e, 0, "", file, line)
}

func wrap(e error, code int, message string, file string, line int) error {
	if code > 0 {
		e = fmt.Errorf(message)
	}

	if e == nil {
		return nil
	}

	switch v := e.(type) {

	case Error:

		trace := PrintStackFormat(v.Flag, file, line, "")
		v.StackTraces = append(v.StackTraces, trace)
		v.Errors = append(v.Errors, v)
		v.Current = e

		if code > 0 {
			v.serviceErrs = append(v.serviceErrs, serviceErr{
				Code: code, Message: message,
			})
		}

		return v
	case error:
		rsp := Error{
			Current:     v,
			StackTraces: make([]string, 0, 30),
			Errors:      make([]error, 0, 30),
			serviceErrs: make([]serviceErr, 0, 30),
			Flag:        Llongfile | LcauseBy | LdateTime,
		}

		trace := PrintStackFormat(rsp.Flag, file, line, e.Error())
		rsp.StackTraces = append(rsp.StackTraces, trace)
		rsp.Errors = append(rsp.Errors, rsp)

		if code > 0 {
			rsp.serviceErrs = append(rsp.serviceErrs, serviceErr{
				Code: code, Message: message,
			})
		}

		return rsp

	}

	panic("")
}

func getFilePath(file string, line int) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return fmt.Sprintf("%s:%d", file[idx+1:], line)

}

func PrintStackFormat(flag int, file string, line int, cause string) string {

	var formatGroup = make([]string, 0, 3)
	var formatArgs = make([]interface{}, 0, 3)

	//if flag&Llongfile > 0 {
	formatGroup = append(formatGroup, "%s")
	//trace := fmt.Sprintf("%s:%d", file, line)
	trace := getFilePath(file, line)

	formatArgs = append(formatArgs, trace)

	formatGroup = append(formatGroup, "%s")

	if cause == "" {
		formatArgs = append(formatArgs, "")
	} else {
		formatArgs = append(formatArgs, "::"+cause)
	}

	//}
	//if flag&LcauseBy > 0 && cause != "" {
	//	formatGroup = append(formatGroup, "%s")
	//	formatArgs = append(formatArgs, cause)
	//}
	return fmt.Sprintf(strings.Join(formatGroup, ""), formatArgs...)
}
