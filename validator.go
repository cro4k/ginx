package ginx

import "github.com/gin-gonic/gin"

type Validator interface {
	Valid(c *gin.Context) error
}

type Empty struct{}

func (e Empty) Valid(c *gin.Context) error {
	return nil
}
