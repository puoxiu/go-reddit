package mysql

import (
	"errors"
	"web-app/models"
)

var (
	ErrorServerBusy = errors.New("服务繁忙")
)

// InsertUser 向数据库中插入用户信息
func InsertUser(u *models.User) error {
	sqlStr := "insert into users(user_id, username, password) values(?, ?, ?)"
	_, err := db.Exec(sqlStr, u.UserID, u.UserName, u.Password)

	return err
}

// GetUserByByName 根据用户名获取用户信息
func GetUserByByName(username string) (user *models.User, err error) {
	sqlStr := "select user_id, username, password from users where username = ?"
	user = new(models.User)
	err = db.Get(user, sqlStr, username)
	
	return 
}


// QueryUserByUserID 根据用户ID查询用户信息
func GetUserByID(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := "select user_id, username from users where user_id = ?"
	err = db.Get(user, sqlStr, uid)
	return
}
