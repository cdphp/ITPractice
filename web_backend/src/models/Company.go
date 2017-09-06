package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Company 企业表
type Company struct {
	ID          int64
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Usable      int    `gorm:"not null"`
	Unusable    int    `gorm:"not null"`
	IsDelete    bool   `gorm:"not null;default:0"`
	CreatedAt   int64  `gorm:"not null"`
	UpdatedAt   int64  `gorm:"not null;default:0"`
}

// TransformedCompany 对外输出json
type TransformedCompany struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Usable      int    `json:"usable"`
	Unusable    int    `json:"unusable"`
	CreatedAt   int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (company *Company) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
