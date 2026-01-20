package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jheader/golang_blog/middleware"
)

func SetupRoutes() *gin.Engine {

	r := gin.New()
	r.Use(middleware.LoggerMiddleWare())
	r.Use(middleware.ErrorHandleMiddleWare())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				// 这里写注册逻辑
				c.JSON(200, gin.H{"message": "register success"})
			})
			auth.POST("/login", func(c *gin.Context) {
				// 这里写登录逻辑
				c.JSON(200, gin.H{"message": "login success"})
			})
		}
	}

	return r
}
