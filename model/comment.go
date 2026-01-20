package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID      uint   `json:"id"gorm:"primaryKey"`
	Content string `json:"content" gorm:"type:text;not null"`

	UserId    uint           `json:"user_id" gorm:"not null"`
	PostId    uint           `json:"post_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
}
