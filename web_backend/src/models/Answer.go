package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Answer 记录答案
type Answer struct {
	ID         int64
	Title      string `gorm:"not null"`
	User       User
	UserID     int64 `gorm:"not null"`
	QuestionID int64 `gorm:"not null"`
	Question   Question
	Content    string `gorm:"not null" sql:"type:text"`
	Approval   int    `gorm:"not null;default:0"`
	IsDelete   bool   `gorm:"not null;default:0"`
	CreatedAt  int64  `gorm:"not null"`
	UpdatedAt  int64  `gorm:"not null;default:0"`
}

// TransformedAnswer 对外输出json
type TransformedAnswer struct {
	ID         int64  `json:"id"`
	QuestionID int64  `json:"question_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Approval   int    `json:"approval"`
	UserID     int64  `json:"user_id"`
	Author     string `json:"author"`
	Avatar     string `json:"avatar"`
	CreatedAt  int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (answer *Answer) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
