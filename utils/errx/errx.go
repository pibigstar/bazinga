package errx

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/pibigstar/bazinga/config"
	"github.com/pibigstar/bazinga/internal/code"
)

type Coder interface {
	Code() int
}

type ErrX struct {
	code Coder
	msg  string
	args []interface{}
}

func (e ErrX) Error() string {
	msg := fmt.Sprintf(e.msg, e.args...)
	return msg
}

func (e ErrX) Code() int {
	return e.code.Code()
}

func NewWithCode(coder Coder, args ...interface{}) *ErrX {
	// 根据code进行翻译
	msg := fmt.Sprintf("code: %d", coder)
	if m, ok := config.CodeMsg[gconv.String(coder.Code())]; ok {
		msg = m
	}
	return &ErrX{
		code: coder,
		msg:  msg,
		args: args,
	}
}

func NewWithMsg(msg string, args ...interface{}) *ErrX {
	return &ErrX{
		code: code.Error_Internal,
		msg:  msg,
		args: args,
	}
}
