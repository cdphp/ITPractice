package model

import (
	"fmt"
	"time"

	"v1.0/vendor"
)

// Score 配置
type Score struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	Action       string `json:"action"`
	Num          int    `json:"num"`
	Type         int    `json:"type"`
	IsDelete     int    `json:"-"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	vendor.Model `json:"-"`
}

// Scores array
type Scores []Score

// NewScore 初始化
func NewScore() *Score {
	s := new(Score)
	s.Init("scores") //设置表名

	return s
}

// Add func
func (s *Score) Add() bool {
	s.CreatedAt = time.Now().Unix()
	if s.Type == 0 {
		s.Type = 1
	}
	//插入数据
	stmt, err := s.ModelManager.Prepare("INSERT " + s.Resource + " SET user_id=?,action=?,num=?,type=?,created_at=?")
	if err != nil {
		fmt.Println("err:", err)
	}

	res, err := stmt.Exec(s.UserID, s.Action, s.Num, s.Type, s.CreatedAt)
	if err != nil {
		fmt.Println("err:", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("err:", err)
	}

	defer stmt.Close()

	if err == nil {
		s.ID = id
		return true
	}
	return false
}
