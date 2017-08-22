package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Relation 关系表
type Relation struct {
	ID        int64
	Master    User
	MasterID  int64 `gorm:"not null"`
	Pupil     User
	PupilID   int64 `gorm:"not null"`
	Type      int   `gorm:"not null;default:1"`
	IsDelete  bool  `gorm:"not null;default:0"`
	CreatedAt int64 `gorm:"not null"`
	UpdatedAt int64 `gorm:"not null;default:0"`
}

// TransformedRelation 对外输出json
type TransformedRelation struct {
	ID        int64  `json:"id"`
	MasterID  int64  `json:"master_id"`
	PupilID   int64  `json:"pupil_id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Labels    string `json:"labels"`
	Type      int    `json:"type"`
	CreatedAt int64  `json:"created_at"`
}

// BeforeCreate 创建数据前的初始化
func (relation *Relation) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedAt", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
