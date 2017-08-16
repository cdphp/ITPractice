package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Config 记录token
type Config struct {
	ID          int64
	Name        string `gorm:"size:32;not null;unique"`
	Value       string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	IsDelete    bool   `gorm:"not null;default:0"`
	CreatedAt   int64  `gorm:"not null"`
	UpdatedAt   int64  `gorm:"not null;default:0"`
}

// TransformedConfig 对外输出json
type TransformedConfig struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (config *Config) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
