package ginx

import (
	"encoding/json"
	"net/http"
)

const (
	CodeFail = 0
	CodeOK   = 1
)

var (
	codes = map[int]int{
		CodeOK:   CodeOK,
		CodeFail: CodeFail,
	}
	messages = map[int]string{}
)

func SetCode(from, to int) {
	codes[from] = to
}

func SetMessage(code int, message string) {
	messages[code] = message
}

func SetCodeMap(m map[int]int) {
	for k := range m {
		codes[k] = m[k]
	}
}

func SetMessageMap(m map[int]string) {
	for k := range m {
		messages[k] = m[k]
	}
}

type response struct {
	RID     string      `json:"rid"`
	CID     string      `json:"cid"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *Context[T]) OK(data ...interface{}) {
	c.rsp(codes[CodeOK], "", data...)
}

func (c *Context[T]) Code(code int, data ...interface{}) {
	c.rsp(code, "", data...)
}

func (c *Context[T]) Fail(err ...string) {
	var msg string
	if len(err) > 0 {
		msg = err[0]
	}
	c.rsp(codes[CodeFail], msg)
}

func (c *Context[T]) FailError(err ...error) {
	var e string
	if len(err) > 0 {
		e = err[0].Error()
	}
	c.Fail(e)
}
func (c *Context[T]) CodeFail(code int, err ...string) {
	var msg string
	if len(err) > 0 {
		msg = err[0]
	}
	c.rsp(code, msg)
}

func (c *Context[T]) CodeFailError(code int, err ...error) {
	var e string
	if len(err) > 0 {
		e = err[0].Error()
	}
	c.CodeFail(code, e)
}

func (c *Context[T]) rsp(code int, message string, data ...interface{}) {
	if message == "" {
		message = messages[code]
	}
	rsp := &response{
		RID:     c.RID,
		CID:     c.CID,
		Code:    code,
		Message: message,
	}
	if len(data) > 0 {
		rsp.Data = data[0]
	}
	if c.signatureSecret != "" {
		s := NewSigner(c.Writer, c.signatureSecret)
		_ = json.NewEncoder(s).Encode(rsp)
		c.Writer.Header().Set("signature", s.Signature())
		c.Writer.Header().Set("Content-Type", "application/json")
	} else {
		c.JSON(http.StatusOK, rsp)
	}
}

func (c *Context[T]) Sign(secret string) *Context[T] {
	c.signatureSecret = secret
	return c
}
