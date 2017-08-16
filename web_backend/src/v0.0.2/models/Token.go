package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Token 记录token
type Token struct {
	ID        int64
	Token     string `gorm:"size:32;not null;unique"`
	User      User
	UserID    int64  `gorm:"not null"`
	Expire    uint   `gorm:"not null;default:0"`
	Auth      string `gorm:"size:20;"`
	LogoutAt  int64  `gorm:"not null;default:0"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedToken 对外输出json
type TransformedToken struct {
	ID        int64  `json:"id"`
	Token     string `json:"token"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Expire    uint   `json:"expire"`
	Auth      string `json:"auth"`
	LogoutAt  string `json:"logout_at"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (token *Token) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	if err := scope.SetColumn("Expire", 3600*24); err != nil {
		return err
	}

	return nil
}
