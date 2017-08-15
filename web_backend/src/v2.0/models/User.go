package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户基本数据，对应数据表格
type User struct {
	ID        uint
	Username  string `gorm:"size:50;not null;unique"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:32;not null"`
	Type      uint   `gorm:"not null;default:1"`
	State     uint   `gorm:"not null;default:0"`
	Profile   Profile
	IsDelete  bool  `gorm:"not null;default:0"`
	CreatedAt int64 `gorm:"not null"`
	UpdatedAt int64 `gorm:"not null;default:0"`
}

// TransformedUser 对外输出json
type TransformedUser struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Type      uint   `json:"type"`
	State     uint   `json:"state"`
	Avatar    string `json:"avatar"`
	About     string `json:"about"`
	Labels    string `json:"labels"`
	Score     uint   `json:"score"`
	CreatedAt int64  `json:"created_at"`
}

// Profile 用户信息
type Profile struct {
	ID        uint
	UserID    uint   `gorm:"not null"`
	Avatar    string `gorm:"not null"`
	About     string `gorm:"not null"`
	Labels    string `gorm:"not null"`
	Score     uint   `gorm:"not null;default:0"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// BeforeCreate 创建数据前的初始化
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now().Unix())
	return err
}
