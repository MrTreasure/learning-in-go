package error

import (
	"fmt"
	"runtime/debug"
)

// MyError 自定义错误类型
type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, msgf string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(msgf, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{})}
}

func (err MyError) error() string {
	return err.Message
}
