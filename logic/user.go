package logic

import (
	"database/sql"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/jwt"
	"web-app/pkg/snowflake"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)



func SignUp(param *models.ParamSignUp) error {	
	// 保存进数据库->dao层
	// 1. 判断用户是否存在
	_,err := mysql.GetUserByByName(param.UserName)
	if err == nil {
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

// Login 登录
func Login(param *models.ParamLogin) (user *models.User, err error) {
	user, err = mysql.GetUserByByName(param.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNoUser
		}
		zap.L().Error("mysql.GetUserByByName() failed", zap.Error(err))
		return nil, ErrorServerBusy
	}
	// 判断密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password)); err != nil {
		return nil, ErrorUserNameOrPassword
	}

	// 生成JWT
	var token string
	if token, err = jwt.GetToken(user.UserID, user.UserName); err != nil {
		zap.L().Error("GenToken failed", zap.Error(err))
		return nil, ErrorServerBusy
	}
	user.Token = token
	return user, nil
}
