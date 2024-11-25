package controllers

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	
	CodeNoAuth
	CodeInvalidToken
)

var codeMsgMap = map[int]string{
	CodeSuccess:       "请求成功",
	CodeInvalidParam:  "请求参数错误, 请检查输入参数",
	CodeUserExist:     "用户已存在",
	CodeUserNotExist:  "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:    "服务繁忙",
	CodeNoAuth:        "无权限访问, 需要登录",
	CodeInvalidToken:   "无效的Token",
}

func GetMsg(code int) string {
	msg, ok := codeMsgMap[code]
	if ok {
		return msg
	}
	return ""
}