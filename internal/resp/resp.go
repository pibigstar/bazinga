package resp

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

type (
	Request struct {
		req  *ghttp.Request
		resp *Resp
	}

	Resp struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

func New(r *ghttp.Request) *Request {
	req := &Request{
		req:  r,
		resp: &Resp{},
	}
	return req
}

func (r *Request) Data(date interface{}) *Request {
	r.resp.Data = date
	return r
}

func (r *Request) Code(code int) *Request {
	r.resp.Code = code
	return r
}

func (r *Request) Msg(msg string) *Request {
	r.resp.Msg = msg
	return r
}

func (r *Request) Write() {
	r.req.Response.WriteJson(r.resp)
	r.req.Exit()
}

func Success(r *ghttp.Request) *Request {
	req := &Request{
		req: r,
		resp: &Resp{
			Msg:  "Ok",
			Code: http.StatusOK,
		},
	}
	return req
}

func WriteSuccess(r *ghttp.Request) {
	req := &Request{
		req: r,
		resp: &Resp{
			Msg:  "Ok",
			Code: http.StatusOK,
		},
	}
	req.Write()
}

func SuccessWithDate(r *ghttp.Request, date interface{}) {
	req := &Request{
		req: r,
		resp: &Resp{
			Msg:  "Ok",
			Data: date,
			Code: http.StatusOK,
		},
	}
	req.Write()
}

func Error(r *ghttp.Request, msg string) {
	req := &Request{
		req: r,
		resp: &Resp{
			Msg:  msg,
			Code: http.StatusInternalServerError,
		},
	}
	req.Write()
	r.Exit()
}

func ErrorWithCode(r *ghttp.Request, code int, msg string) {
	req := &Request{
		req: r,
		resp: &Resp{
			Msg:  msg,
			Code: code,
		},
	}
	req.Write()
}
