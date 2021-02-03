package controller

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ret       int         `json:"ret"`
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
}

func Error(ctx *gin.Context, ret int, code int, error string, data interface{}) {
	r := Response{
		Ret:       ret,
		ErrorCode: code,
		Msg:       error,
		Error:     "error",
		Data:      data,
	}
	ctx.JSON(200, r)
}

func Render(ctx *gin.Context, msg string, data interface{}) {
	r := Response{
		Ret:       0,
		Msg:       msg,
		ErrorCode: 200,
		Error:     "ok",
		Data:      data,
	}
	ctx.JSON(200, r)
}
