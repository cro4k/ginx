### gin wrapper

Usage:
```go
package main

import (
    "github.com/cro4k/ginx"
)

const(
    CodeInvalidParams = 1000
)

func main() {
    ginx.SetCode(ginx.CodeOK, 200)
    ginx.SetCode(CodeInvalidParams, CodeInvalidParams)
    ginx.SetMessage(CodeInvalidParams,"invalid params")
    
    eng := gin.Default()
    eng.POST("/hello", sayHello)
    http.ListenAndServe(":8080",eng)
}

type Request struct{
    ginx.Empty
    Name string `json:"name"`
}

type Response struct {
    Message string `json:"message"`
}

func sayHello(c *gin.Context) {
    ctx, err := ginx.With[Request](c)
    if err!=nil{
        ctx.FailError(err)
        return
    }
    msg := "Hello, " + ctx.Body.Name + "!"
    ctx.OK(&Response{Message:msg})
}
```