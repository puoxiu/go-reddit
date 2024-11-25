package mysql

import (
	"fmt"
	"web-app/settings"

	_ "github.com/go-sql-driver/mysql" // 显式导入 MySQL 驱动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func InitDB(cfg *settings.MySQLConfig) error {
	// 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return err
	}

	// 设置数据库连接池
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)



	return nil
}

func Close() {
	// 关闭数据库
    _ = db.Close()
}

func Judge() {
	if db == nil {
		fmt.Println("db is nil in u1111212213123123")
	} else {
		fmt.Println("db is not nil in 3132132132132132")
	}
}
