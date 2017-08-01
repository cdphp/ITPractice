package model

import (
	"strconv"
	"time"

	"v1.0/vendor"
)

// Token struct
type Token struct {
	ID           int64  `json:"id"`
	Token        string `json:"token"`
	UserID       int64  `json:"user_id"`
	Expire       int    `json:"expire"`
	Auth         string `json:"auth"`
	LogoutAt     int64  `json:"lougout_at"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	vendor.Model `json:"-"`
}

// NewToken 初始化
func NewToken() *Token {
	t := new(Token)
	t.Init("token")

	return t
}

// Obtian 获取token
func (t *Token) Obtian(userID int64, auth string, expire int) (*Token, int) {
	t.Token = t.Md5(strconv.FormatInt(time.Now().UnixNano(), 10))
	t.UserID = userID
	t.Expire = expire
	t.Auth = auth
	t.CreatedAt = time.Now().Unix()

	if t.Add() {
		return t, 0
	}
	return nil, 23
}

// Add 添加
func (t *Token) Add() bool {

	//插入数据
	stmt, err := t.ModelManager.Prepare("INSERT tokens SET token=?,user_id=?,expire=?,auth=?,created_at=?")

	res, err := stmt.Exec(t.Token, t.UserID, t.Expire, t.Auth, t.CreatedAt)

	id, err := res.LastInsertId()

	if err == nil {
		t.ID = id
		return true
	}

	return false

}
