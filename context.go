package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Context[T Validator] struct {
	RID string // request id
	CID string // client id
	UID string // user id
	*gin.Context
	Body T
}

func With[T Validator](ctx *gin.Context) (*Context[T], error) {
	c := with[T](ctx)
	if err := ctx.ShouldBind(&c.Body); err != nil {
		return c, err
	}
	if err := c.Body.Valid(ctx); err != nil {
		return c, err
	}
	return c, nil
}

func WithJSON[T Validator](ctx *gin.Context) (*Context[T], error) {
	c := with[T](ctx)
	if err := ctx.ShouldBindJSON(&c.Body); err != nil {
		return c, err
	}
	if err := c.Body.Valid(ctx); err != nil {
		return c, err
	}
	return c, nil
}

func Ctx(ctx *gin.Context) *Context[Empty] {
	return with[Empty](ctx)
}

func with[T Validator](ctx *gin.Context) *Context[T] {
	c := &Context[T]{}
	c.Context = ctx
	c.RID = ctx.GetString("rid")
	c.CID = ctx.GetString("cid")
	c.UID = ctx.GetString("uid")
	body := new(T)
	c.Body = *body
	return c
}

func (c *Context[T]) Logger() *logrus.Entry {
	return logrus.WithContext(c)
}
