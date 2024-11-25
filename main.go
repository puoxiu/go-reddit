package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/logger"
	"web-app/pkg/snowflake"

	// "web-app/pkg/snowflake"
	"web-app/routes"
	"web-app/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/*
	通用的go语言web开发框架
*/

func main() {
	// 1. 加载配置
	err := settings.Init()
	if err != nil {
		fmt.Printf("init settings error: %v\n", err)
	}

	// 2. 初始化日志
	err = logger.Init(settings.Conf.LogConfig, settings.Conf.AppConfig.Mode)
	if err != nil {
		fmt.Printf("init logger error: %v\n", err)
	}
	defer zap.L().Sync()	// 将缓冲区的日志追加到文件中
	zap.L().Debug("logger init success...")
	
	// 3. 初始化MySQL连接
	err = mysql.InitDB(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql error: %v\n", err)
		return
	}

	defer mysql.Close()

	// 4. 初始化Redis连接
	err = redis.Init(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Printf("init redis error: %v\n", err)
		return
	}
	defer redis.Close()	

	// 雪花算法
	err = snowflake.Init(settings.Conf.StartTime, settings.Conf.AppConfig.MachineID)
	if err != nil {
		fmt.Printf("init snowflake error: %v\n", err)
		return
	}

	// 5. 注册路由
	server := routes.Register()

	// 6. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: server,
	}

	go func()  {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Sugar().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)	// 捕获信号 此处不会阻塞
	<-quit	// 阻塞在此处，当接收到信号时才会往下执行

	zap.L().Info("Shutdown Server ...")

	// 优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Sugar().Fatal("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}	
