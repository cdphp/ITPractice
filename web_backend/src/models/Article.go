package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 记录文章
type Article struct {
	ID        int64
	Title     string `gorm:"not null"`
	User      User
	UserID    int64  `gorm:"not null"`
	Digest    string `gorm:"not null"`
	Content   string `gorm:"not null" sql:"type:text"`
	Labels    string `gorm:"not null"`
	Clicks    uint   `gorm:"not null;default:0"`
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedArticle 对外输出json
type TransformedArticle struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Digest    string `json:"digest"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	Author    string `json:"author"`
	Avatar    string `json:"avatar"`
	Labels    string `json:"labels"`
	Clicks    uint   `json:"clicks"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
