package main

import (
	"fmt"

	"github.com/jheader/golang_blog/config"
	"github.com/jheader/golang_blog/model"
)

func main() {

	// 初始化日志
	fmt.Println("Starting blog application...")
	//加载配置
	config.InitViper()
	// 初始化数据库
	config.InitDB()

	var existingUser model.User
	if err := config.DB.Where("username = ?", "").First(&existingUser).Error; err == nil {
		fmt.Println("err=", err)
		return
	}

	// 执行生成，复制输出作为你的 JWT_SECRET
	fmt.Println("JWT_SECRET=", existingUser)
}
