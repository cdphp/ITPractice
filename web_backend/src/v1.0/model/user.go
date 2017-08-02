package model

import (
	"time"

	"v1.0/vendor"
)

// User 用户信息
type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Type         int    `json:"type"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	vendor.Model `json:"-"`
}

// Users array
type Users []User

// NewUser 初始化
func NewUser() *User {
	u := new(User)
	u.Init("user")

	return u
}

// Get 根据id获取数据
func (u *User) Get(id int) (*User, int) {

	sql := "select id,username,email,password,type,created_at,updated_at from users where id=?"
	err := u.ModelManager.QueryRow(sql, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, 101
	}

	return u, 0

}

// Add 添加
func (u *User) Add() bool {

	//插入数据
	stmt, err := u.ModelManager.Prepare("INSERT users SET username=?,email=?,password=?,type=?,created_at=?")

	res, err := stmt.Exec(u.Username, u.Email, u.Md5(u.Password), u.Type, u.CreatedAt)

	id, err := res.LastInsertId()

	if err == nil {
		u.ID = id
		return true
	}
	return false

}

// Auth 验证
func (u *User) Auth(username string, password string) (*User, int) {

	sql := "select id,username,email,password,type,created_at,updated_at from users where username=?"
	err := u.ModelManager.QueryRow(sql, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, 101
	}

	if u.Md5(password) != u.Password {
		return nil, 102
	}

	return u, 0

}

// Register 注册
func (u *User) Register(username, email, password string, userType int) (*User, int) {
	if u.HasName(username) {
		return nil, 103
	}

	if u.HasEmail(email) {
		return nil, 104
	}

	u.Username = username
	u.Email = email
	u.Password = password
	u.CreatedAt = time.Now().Unix()
	u.Type = userType

	if u.Add() {
		return u, 0
	}
	return nil, 23
}

// HasName 判断用户名是否已存在
func (u *User) HasName(username string) bool {

	sql := "select id from users where username=?"
	err := u.ModelManager.QueryRow(sql, username).Scan(&u.ID)

	if err == nil {
		return true
	}
	return false
}

// HasEmail 判断邮箱是否已存在
func (u *User) HasEmail(email string) bool {

	sql := "select id from users where email=?"
	err := u.ModelManager.QueryRow(sql, email).Scan(&u.ID)

	if err == nil {
		return true
	}
	return false
}

// GetAuthName 获取等级名称
func (u *User) GetAuthName(auth int) string {
	var name string
	if auth == 1 {
		name = "User"
	} else if auth == 2 {
		name = "Master"
	} else {
		name = "Admin"
	}
	return name
}
