package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;size 100"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

// 如果没有错误信息 就是保存更新成，更新是有ID加自己文章才能更新
func (p *Post) SaveOrUpdate(db *gorm.DB) error {

	return nil
}
