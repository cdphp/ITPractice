package main

import (
	"crypto/md5"
	"encoding/hex"
)

// User 用户信息
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Type      int    `json:"type"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
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
func (u *User) Add() int64 {
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

	return id

}

// Md5 加密
func (u *User) Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
