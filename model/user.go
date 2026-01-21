package model

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex; not null; size:50"`
	Password  string         `json:"password,omitempty" gorm:"not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Posts    []Post    `json:"posts,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

func GetUserByMap(mp map[string]interface{}, db *gorm.DB) ([]User, error) {

	var users []User
	result := db.Find(&users, mp)
	if result.Error != nil {
		return nil, fmt.Errorf("查询失败：%w", result.Error)

	}
	return users, nil

}

// result.Error == nil 只能说明「查询执行成功（无数据库异常）」，
// 但不能 100% 等同于「查询到了数据」—— 仅当使用 First/Take/Last 这类「查询单条记录」的方法时，
// result.Error == nil 才意味着「查询到了数据」；若使用 Find（查询多条）或 Count（统计数量），则需要额外判断结果长度 / 数量
func GetUserByUsername(username string, db *gorm.DB) (*User, error) {

	var u User
	result := db.Where("username = ?", username).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在(username: %s):%w", username, result.Error)
		}
		return nil, fmt.Errorf("查询用户失败(username: %s):%w", username, result.Error)
	}
	return &u, nil
}

func (u *User) Save(db *gorm.DB) error {

	result := db.Create(u)
	if result.Error != nil {
		return fmt.Errorf("创建(sername: %s):%w", (*u).Username, result.Error)
	}

	return nil

}

func (u *User) HashPassword() error {
	// 1. 检查明文密码是否为空（避免哈希空字符串）
	if u.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	// 2. 生成 bcrypt 盐值（cost 越大越安全，但耗时更长，推荐 10-14）
	salt, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err) // 包装错误，保留原始错误信息
	}

	// 3. 将哈希后的密码（字节切片）转为字符串，覆盖原明文密码
	u.Password = string(salt)
	return nil
}

// BeforeCreate GORM钩子，在创建用户前自动哈希密码
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.HashPassword()
}
