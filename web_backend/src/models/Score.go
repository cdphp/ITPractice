package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Score 记录文章
type Score struct {
	ID        int64
	User      User
	UserID    int64  `gorm:"not null"`
	Action    string `gorm:"not null"`           //操作:article,comment
	Type      uint   `gorm:"not null;default:1"` //类型:1表示加,2:表示减
	Num       int    `gorm:"not null"`           //数量
	IsDelete  bool   `gorm:"not null;default:0"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"not null;default:0"`
}

// TransformedScore struct
type TransformedScore struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Action    string `json:"action"`
	Num       int    `json:"num"`
	Type      int    `json:"type"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (score *Score) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
