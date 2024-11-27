package controllers

import (
	"errors"
	"strconv"
	"web-app/logic"
	"web-app/models"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	emailRegexExp *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		// 编译期间就确定了，提高性能
		emailRegexExp: regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, regexp.None),
		passwordRegexExp: regexp.MustCompile("", regexp.None),
	}
}	

/*
{
    "username":"123456@qq.com",
    "password":"12345",
    "confirm_password":"12345"
}
*/
func (u *UserHandler)SignUpHandler(ctx *gin.Context) {
	// 1. 获取参数和参数校验
	var param models.ParamSignUp
	if err := ctx.ShouldBind(&param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return	
	}

	ok, err := u.emailRegexExp.MatchString(param.UserName)
	if err != nil {
		zap.L().Error("emailRegexExp.MatchString failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	if !ok {
		// 用户名格式错误
		ResponseErrorWithMsg(ctx, CodeInvalidParam, "用户名格式错误（必须符合邮箱格式）")
		return
	}

	ok, err = u.passwordRegexExp.MatchString(param.Password)
	if err != nil {
		zap.L().Error("passwordRegexExp.MatchString failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	if !ok {
		ResponseErrorWithMsg(ctx, CodeInvalidParam, "密码格式错误")
		return
	}

	if param.Password != param.ConfirmPassword {
		ResponseErrorWithMsg(ctx, CodeInvalidParam, "两次密码不一致")
		return
	}
	
	
	// 加密
	hash,err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("bcrypt.GenerateFromPassword failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	param.Password = string(hash)

	// 2. 业务处理--在logic层处理
	if err := logic.SignUp(&param); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		//
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
			return
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}

	// 返回正确响应
	ResponseSuccess(ctx, "注册成功, 这是注册成功的data")
}


/*
{
    "username":"123456@qq.com",
    "password":"12345"
}
*/
func (u *UserHandler)LoginHandler(ctx *gin.Context) {
	// 1. 获取参数和参数校验
	var param models.ParamLogin
	if err := ctx.ShouldBind(&param); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return	
	}

	ok, err := u.emailRegexExp.MatchString(param.UserName)
	if err != nil {
		zap.L().Error("emailRegexExp.MatchString failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	if !ok {
		ResponseErrorWithMsg(ctx, CodeInvalidParam, "用户名格式错误（必须符合邮箱格式）")
		return
	}

	// 2. 业务处理--在logic层处理
	
	user, err := logic.Login(&param)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, logic.ErrorServerBusy) {
			ResponseError(ctx, CodeServerBusy)
			return
		}
		ResponseErrorWithMsg(ctx, CodeInvalidPassword, "用户名或密码错误")
		return
	}

	// 3. 返回响应
	// ResponseSuccess(ctx, "登录成功了,这是登录成功的data")
	ResponseSuccess(ctx, gin.H{
		// int64-->string
		"user_id": strconv.FormatInt(user.UserID, 10),
		"user_name": user.UserName,
		"token": user.Token,
	})
}