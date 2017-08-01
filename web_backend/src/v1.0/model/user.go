package model

import "v1.0/vendor"

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
func (u *User) Get(id int) (*User, *vendor.Error) {

	sql := "select id,username,email,password,type,created_at,updated_at from users where id=?"
	err := u.ModelManager.QueryRow(sql, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, &vendor.Error{101, "该账号不存在"}
	}

	return u, nil

}

// Add 添加
func (u *User) Add() {

	//插入数据
	stmt, err := u.ModelManager.Prepare("INSERT users SET username=?,email=?,password=?,type=?,created_at=?")

	res, err := stmt.Exec(u.Username, u.Email, u.Md5(u.Password), u.Type, u.CreatedAt)

	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	u.ID = id

}

// Auth 验证
func (u *User) Auth(username string, password string) (*User, *vendor.Error) {

	sql := "select id,username,email,password,type,created_at,updated_at from users where username=?"
	err := u.ModelManager.QueryRow(sql, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, &vendor.Error{101, "该账号不存在"}
	}

	if u.Md5(password) != u.Password {
		return nil, &vendor.Error{102, "密码错误，请重试"}
	}

	return u, nil

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
