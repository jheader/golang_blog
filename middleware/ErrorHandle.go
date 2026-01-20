package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandleMiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				logrus.WithFields(logrus.Fields{
					"error":  err,
					"path":   ctx.Request.URL.Path,
					"method": ctx.Request.Method,
				}).Error("Panic recovered")

				ctx.JSON(500, gin.H{
					"code":    500,
					"message": "Internal server error",
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}

}
