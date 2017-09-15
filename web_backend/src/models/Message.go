package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Message 消息
type Message struct {
	ID        int64
	User      User
	UserID    int64  `gorm:"not null"`
	Content   string `gorm:"not null"`
	Target    User
	TargetID  int64 `gorm:"not null"`
	IsDelete  bool  `gorm:"not null;default:0"`
	CreatedAt int64 `gorm:"not null"`
	UpdatedAt int64 `gorm:"not null;default:0"`
}

// TransformedMessage struct
type TransformedMessage struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
	Content      string `json:"content"`
	TargetID     int64  `json:"target_id"`
	TargetName   string `json:"target_name"`
	TargetAvatar string `json:"target_avatar"`
	CreatedAt    int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (message *Message) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
