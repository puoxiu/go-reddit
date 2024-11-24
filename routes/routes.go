package routes

import (
	"net/http"
	"web-app/logger"

	"github.com/gin-gonic/gin"
)

func Register() (*gin.Engine){
	server := gin.Default()
	
	server.Use(logger.GinLogger(), logger.GinRecovery(true))

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	return server
}