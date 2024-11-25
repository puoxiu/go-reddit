package mysql

import (
	"database/sql"
	"errors"
	"web-app/models"

	"go.uber.org/zap"
)

var (
	ErrorServerBusy = errors.New("服务繁忙")
)

// InsertUser 向数据库中插入用户信息
func InsertUser(u *models.User) error {
	sqlStr := "insert into users(user_id, username, password) values(?, ?, ?)"

	if _, err := db.Exec(sqlStr, u.UserID, u.UserName, u.Password); err != nil {
		zap.L().Error("InsertUser failed", zap.Error(err))
		return err
	}

	return nil
}

// CheckUserExistByName 根据用户名判断用户是否存在
func CheckUserExistByName(username string) (bool, error) {
	sqlStr := "select count(user_id) from users where username = ?"

	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		zap.L().Error("CheckUserExistByName failed", zap.Error(err))
		return false, ErrorServerBusy
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 检查用户名判断是否正确, 返回密码
func CheckUser(user *models.User) (error) {
	sqlStr := "select user_id, username, password from users where username = ?"
	err := db.Get(user, sqlStr, user.UserName) 

	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

// QueryUserByUserID 根据用户ID查询用户信息
func QueryUserByUserID()  {
	// 执行SQL
}
