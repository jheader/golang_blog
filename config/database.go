package config

import (
	"fmt"
	"log"

	"github.com/jheader/golang_blog/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	var err error

	dbHost := viper.Get("DB_HOST")
	dbPort := viper.Get("DB_PORT")
	dbUser := viper.Get("DB_USER")
	dbPassword := viper.Get("DB_PASSWORD")
	dbName := viper.Get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {

		log.Fatal("Failed to connect to MySQL database:", err)
	}

	err = DB.AutoMigrate(
		&model.Comment{},
		&model.Post{},
		&model.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("MySQL database connected and migrated successfully")
}

func InitViper() {

	viper.SetConfigType("env")
	// 优先在当前工作目录找
	viper.AddConfigPath(".")
	// 再去上级目录找
	viper.AddConfigPath("..")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("警告：未找到 .env 配置文件，将使用默认值或环境变量：%v\n", err)
	}

	// 5. 重要：让 Viper 自动读取环境变量（优先级高于配置文件）
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "123456") // 开发环境默认密码，生产环境通过环境变量覆盖
	viper.SetDefault("DB_NAME", "golang_blog")

}
