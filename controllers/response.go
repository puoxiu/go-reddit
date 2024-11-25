package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
构建响应对象
{
	"code": 1000,
	"msg": "请求成功",
	"data": {}
}

*/
type ResponseData struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Data interface{} `json:"data"`
}


func NewResponseData(code int, data interface{}) *ResponseData {
	return &ResponseData{
		Code: code,
		Msg: GetMsg(code),
		Data: data,
	}
}

func ResponseError(ctx *gin.Context, code int) {
	rd := NewResponseData(code, nil)
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code int, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg: msg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := NewResponseData(CodeSuccess, data)
	ctx.JSON(http.StatusOK, rd)
}