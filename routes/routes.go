package routes

import (
	"fmt"
	"net/http"
	"web-app/controllers"
	"web-app/logger"
	"web-app/middlewares"

	"github.com/gin-gonic/gin"
)

func Register() (*gin.Engine){
	server := gin.Default()
	
	server.Use(logger.GinLogger(), logger.GinRecovery(true))

	fmt.Println("routes.Register()")

	server.GET("/", func(c *gin.Context) {
		fmt.Println("Hello World!")
		c.String(http.StatusOK, "Hello World!")
	})

	contr := controllers.NewUserHandler()
	server.POST("/signup", contr.SignUpHandler)

	server.POST("/login", contr.LoginHandler)

	server.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 登录后才能访问的接口
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "pong",
		})
	})

	return server
}
