package logic

import (
	"errors"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/jwt"
	"web-app/pkg/snowflake"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorUserExist = errors.New("用户已经存在")
	ErrorUserNameOrPassword = errors.New("用户名或者密码错误")
	ErrorServerBusy = errors.New("服务繁忙")
)


func SignUp(param *models.ParamSignUp) error {	
	// 保存进数据库->dao层
	// 1. 判断用户是否存在
	exit,err := mysql.CheckUserExistByName(param.UserName)
	if err != nil {
		return ErrorServerBusy
	}
	
	if exit {
		return ErrorUserExist
	}

	// 2. 保存用户信息
	uid := snowflake.GetID()
	u := &models.User{
		UserID: uid,
		UserName: param.UserName,
		Password: param.Password,
	}
	if err = mysql.InsertUser(u); err != nil {
		return ErrorServerBusy
	}
	return nil
}


func Login(param *models.ParamLogin) (string, error) {
	user := &models.User{
		UserName: param.UserName,
		Password: param.Password,
	}
	err := mysql.CheckUser(user)
	if err != nil {
		return "", ErrorServerBusy
	}
	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password)); err != nil {
		return "", ErrorUserNameOrPassword
	}

	// 生成JWT
	var token string
	if token, err = jwt.GenToken(user.UserID, user.UserName); err != nil {
		zap.L().Error("GenToken failed", zap.Error(err))
		return "", ErrorServerBusy
	}
	return token, nil
}
