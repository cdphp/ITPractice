package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Comment 记录文章
type Comment struct {
	ID        int64
	User      User
	UserID    int64  `gorm:"not null"`
	Content   string `gorm:"not null"`
	Type      uint   `gorm:"not null;default:1"`
	TargetID  int64  `gorm:"not null"`
	RootID    int64  `gorm:"not null"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedComment struct
type TransformedComment struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	Content   string `json:"content"`
	Type      uint   `json:"type"`
	TargetID  int64  `json:"target_id"`
	RootID    int64  `json:"root_id"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (comment *Comment) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
