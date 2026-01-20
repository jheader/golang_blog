package model

import (
	"time"

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
