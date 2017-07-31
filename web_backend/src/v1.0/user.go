package main

import (
	"crypto/md5"
	"encoding/hex"
)

// User 用户信息
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Type      int    `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Users array
type Users []User

// Get 根据id获取数据
func (u *User) Get(id int) *User {

	helper := new(DbHelper)
	helper.Init()

	sql := "select id,username,email,password,type,created_at,updated_at from users where id=?"
	err := helper.db.QueryRow(sql, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		panic(err)
	}
	defer helper.db.Close()
	return u

}

// Add 添加
func (u *User) Add() {
	helper := new(DbHelper)
	helper.Init()

	//插入数据
	stmt, err := helper.db.Prepare("INSERT users SET username=?,email=?,password=?,type=?,created_at=?")

	res, err := stmt.Exec(u.Username, u.Email, u.Md5(u.Password), u.Type, u.CreatedAt)

	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	defer helper.db.Close()

	u.ID = id

}

// Auth 验证
func (u *User) Auth(username string, password string) (*User, *Error) {
	helper := new(DbHelper)
	helper.Init()

	sql := "select id,username,email,password,type,created_at,updated_at from users where username=?"
	err := helper.db.QueryRow(sql, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, &Error{101, "该账号不存在"}
	}

	if u.Md5(password) != u.Password {
		return nil, &Error{102, "密码错误，请重试"}
	}

	return u, nil

}

// Md5 加密
func (u *User) Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// HasName 判断用户名是否已存在
func (u *User) HasName(username string) bool {
	helper := new(DbHelper)
	helper.Init()

	sql := "select id from users where username=?"
	err := helper.db.QueryRow(sql, username).Scan(&u.ID)

	defer helper.db.Close()

	if err == nil {
		return true
	}
	return false
}

// HasEmail 判断邮箱是否已存在
func (u *User) HasEmail(email string) bool {
	helper := new(DbHelper)
	helper.Init()

	sql := "select id from users where email=?"
	err := helper.db.QueryRow(sql, email).Scan(&u.ID)

	defer helper.db.Close()

	if err == nil {
		return true
	}
	return false
}
