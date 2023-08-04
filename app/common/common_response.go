package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Context interface{} `json:"context"`
}

func newResponse() *Response {
	return &Response{
		Code:    200,
		Msg:     "success",
		Context: nil,
	}
}

func Success(ctx *gin.Context, data interface{}) {
	resp := newResponse()
	resp.Context = data
	ctx.JSON(http.StatusOK, resp)
}

func Error(ctx *gin.Context, statusCode int, errorMsg string) {
	resp := newResponse()
	resp.Msg = errorMsg
	resp.Code = statusCode
	ctx.JSON(statusCode, resp)
}
