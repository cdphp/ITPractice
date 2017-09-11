package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Question 记录问题
type Question struct {
	ID        int64
	Title     string `gorm:"not null"`
	User      User
	UserID    int64  `gorm:"not null"`
	Content   string `gorm:"not null" sql:"type:text"`
	Clicks    int    `gorm:"not null;default:0"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedQuestion 对外输出json
type TransformedQuestion struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	Author    string `json:"author"`
	Clicks    int    `json:"clicks"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (question *Question) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
