package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Cheat 关系表
type Cheat struct {
	ID        int64
	User      User
	UserID    int64  `gorm:"not null"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null" sql:"type:text"`
	Labels    string `gorm:"not null"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedCheat 对外输出json
type TransformedCheat struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Labels    string `json:"labels"`
	Type      int    `json:"type"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (cheat *Cheat) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
