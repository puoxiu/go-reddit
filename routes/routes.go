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
	v1 := server.Group("/api/v1")
	v1.POST("/signup", contr.SignUpHandler)
	v1.POST("/login", contr.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post", controllers.GetPostListHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
	}

	// 测试接口
	server.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 登录后才能访问的接口
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "pong",
		})
	})

	return server
}
