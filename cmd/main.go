package main

import (
	"log"

	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// 初始化日志
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting blog application...")
	//加载配置
	config.InitViper()
	// 初始化数据库
	config.InitDB()

	// 设置路由
	r := routes.SetupRoutes()

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
