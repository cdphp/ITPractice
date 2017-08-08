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
	Username     string `json:"username"`
	Expire       int64  `json:"expire"`
	Auth         string `json:"auth"`
	LogoutAt     int64  `json:"lougout_at"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	vendor.Model `json:"-"`
}

// NewToken 初始化
func NewToken() *Token {
	t := new(Token)
	t.Init("tokens")

	return t
}

// Obtian 获取token
func (t *Token) Obtian(userID int64, username, auth string, expire int64) bool {
	t.Token = vendor.Md5(strconv.FormatInt(time.Now().UnixNano(), 10))
	t.UserID = userID
	t.Username = username
	t.Expire = expire
	t.Auth = auth
	t.CreatedAt = time.Now().Unix()

	return t.Add()
}

// Validate 验证
func (t *Token) Validate(token string) bool {
	sql := "select id,user_id, username,expire,auth,logout_at,created_at from " + t.Resource + " where token=? and is_delete=0"
	err := t.ModelManager.QueryRow(sql, token).Scan(&t.ID, &t.UserID, &t.Username, &t.Expire, &t.Auth, &t.LogoutAt, &t.CreatedAt)

	// 没找到token
	if err != nil {
		return false
	}

	// token过期
	if time.Now().Unix()-t.CreatedAt > t.Expire {
		return false
	}
	defer t.CloseDb()

	return true
}

// Add 添加
func (t *Token) Add() bool {

	//插入数据
	stmt, err := t.ModelManager.Prepare("INSERT tokens SET token=?,user_id=?,username=?,expire=?,auth=?,created_at=?")

	res, err := stmt.Exec(t.Token, t.UserID, t.Username, t.Expire, t.Auth, t.CreatedAt)

	id, err := res.LastInsertId()

	if err == nil {
		t.ID = id
		return true
	}

	return false

}
