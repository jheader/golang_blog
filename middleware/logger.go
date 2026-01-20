package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleWare() gin.HandlerFunc {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		logger.WithFields(logrus.Fields{
			"status_code": param.StatusCode,
			"latency":     param.Latency,
			"client_ip":   param.ClientIP,
			"method":      param.Method,
			"path":        param.Path,
			"user_agent":  param.Request.UserAgent(),
			"error":       param.ErrorMessage,
			"timestamp":   param.TimeStamp.Format(time.RFC3339),
		}).Info("HTTP Request")
		return ""
	})

}
