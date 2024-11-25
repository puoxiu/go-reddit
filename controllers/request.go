package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"
var (
	ErrorUserNotLogin = errors.New("用户未登录")
)

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	// 从请求的上下文c中获取当前请求的用户信息
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}