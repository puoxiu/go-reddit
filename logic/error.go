package logic

import "errors"


var (
	ErrorUserExist = errors.New("用户已经存在")
	ErrorNoUser = errors.New("用户不存在")
	ErrorUserNameOrPassword = errors.New("用户名或者密码错误")
	ErrorServerBusy = errors.New("服务繁忙")
	ErrorNoData = errors.New("没有数据")
)