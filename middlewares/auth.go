package middlewares

import (
	"strings"
	"time"
	"web-app/controllers"
	"web-app/pkg/jwt"
	"web-app/pkg/jwtV2"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header中获取jwt
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNoAuth)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 校验jwt
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userid信息保存到请求的上下文c上
		c.Set(controllers.CtxUserIDKey, mc.UserID) 

		c.Next()
	}
}

// JWTAuthMiddlewareV2 基于JWT的认证中间件，支持refresh token
func JWTAuthMiddlewareV2() gin.HandlerFunc {
// JWTAuthMiddleware 基于JWT的认证中间件
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNoAuth)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}

		// 检查 token 是否快过期
		if time.Until(time.Unix(mc.ExpiresAt, 0)) < 10 * time.Minute {
			// 快过期，自动刷新
			newAccessToken, _, err := jwtV2.GetToken(mc.UserID, mc.Username)
			if err != nil {
				controllers.ResponseError(c, controllers.CodeServerBusy)
				c.Abort()
				return
			}
			// 把新 token 添加到响应头
			c.Writer.Header().Set("Authorization", "Bearer "+newAccessToken)
		}

		// 把用户信息存储到上下文
		c.Set(controllers.CtxUserIDKey, mc.UserID)

		c.Next()
	}
}
