package model

import (
	"fmt"
	"strconv"
	"time"

	"v1.0/vendor"
)

// User 用户信息
type User struct {
	ID           int64    `json:"id"`
	UniqueKey    string   `json:"unique_key"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Password     string   `json:"-"`
	Type         int      `json:"type"`
	CreatedAt    int64    `json:"created_at"`
	UpdatedAt    int64    `json:"updated_at"`
	Info         Userinfo `json:"info"`
	vendor.Model `json:"-"`
}

// Userinfo 用户信息
type Userinfo struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"-"`
	Avatar    string `json:"avatar"`
	Bg        string `json:"bg"`
	About     string `json:"about"`
	Labels    string `json:"labels"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// Users array
type Users []*User

// NewUser 初始化
func NewUser() *User {
	u := new(User)
	u.Init("users")

	return u
}

// ListData user list
func (u *User) ListData(page, row int) Users {
	start := (page - 1) * row

	sql := "SELECT id,username,email FROM " + u.Resource + " where is_delete=0 and type=1 limit " + strconv.Itoa(start) + "," + strconv.Itoa(row)

	rows, err := u.ModelManager.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	var result Users
	for rows.Next() {
		var user = NewUser()
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err == nil && user.GetInfo() {
			result = append(result, user)
		}

	}
	return result
}

// Get 根据id获取数据
func (u *User) Get(id int) (*User, int) {

	sql := "select id,username,email,type,created_at,updated_at from " + u.Resource + " where id=?"
	err := u.ModelManager.QueryRow(sql, id).Scan(&u.ID, &u.Username, &u.Email, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err == nil && u.GetInfo() {
		return u, 0
	}
	return nil, 101

}

// GetInfo 根据id获取用户信息
func (u *User) GetInfo() bool {
	sql := "select id,avatar,bg,about,labels from users_info where user_id=?"

	err := u.ModelManager.QueryRow(sql, u.ID).Scan(&u.Info.ID, &u.Info.Avatar, &u.Info.Bg, &u.Info.About, &u.Info.Labels)

	if err == nil {

		return true
	}
	return false
}

// Add 添加
func (u *User) Add() bool {

	//插入数据
	stmt, err := u.ModelManager.Prepare("INSERT users SET unique_key=?,username=?,email=?,password=?,type=?,created_at=?")

	res, err := stmt.Exec(u.UniqueKey, u.Username, u.Email, u.Password, u.Type, u.CreatedAt)

	id, err := res.LastInsertId()

	if err == nil {
		u.ID = id
		return true
	}
	return false

}

// Auth 验证
func (u *User) Auth(password string) bool {

	sql := "select id,email,password,type,created_at,updated_at from users where username=?"
	err := u.ModelManager.QueryRow(sql, u.Username).Scan(&u.ID, &u.Email, &u.Password, &u.Type, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return false
	}

	if vendor.Md5(password) != u.Password {
		return false
	}

	return true
}

// Register 注册
func (u *User) Register() bool {

	u.UniqueKey = vendor.UniqueID()
	u.Password = vendor.Md5(u.Password)
	u.CreatedAt = time.Now().Unix()

	if u.Add() && u.AddInfo() {
		return true
	}
	return false
}

// CheckName 判断昵称是否可用
func (u *User) CheckName() bool {

	sql := "select id from users where username=?"
	err := u.ModelManager.QueryRow(sql, u.Username).Scan(&u.ID)

	if err == nil {
		return false
	}
	return true
}

// CheckEmail 判断邮箱是否可用
func (u *User) CheckEmail() bool {

	sql := "select id from users where email=?"
	err := u.ModelManager.QueryRow(sql, u.Email).Scan(&u.ID)

	if err == nil {
		return false
	}
	return true
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

// AddInfo 添加用户信息
func (u *User) AddInfo() bool {

	u.Info.UserID = u.ID
	u.Info.Avatar = "http://120.77.153.236:4869/38c934d558c5b12766781553c6279a32"
	u.Info.Bg = "http://120.77.153.236:4869/cdfe10a6b3eab163cdd3401388ecf8f6"
	u.Info.CreatedAt = time.Now().Unix()

	//插入数据
	stmt, err := u.ModelManager.Prepare("INSERT users_info SET user_id=?,avatar=?,bg=?,created_at=?")

	res, err := stmt.Exec(u.Info.UserID, u.Info.Avatar, u.Info.Bg, u.Info.CreatedAt)

	id, err := res.LastInsertId()

	if err != nil {
		return false
	}
	u.Info.ID = id
	return true
}

// UpdateInfo 修改用户信息
func (u *User) UpdateInfo() bool {
	u.Info.UpdatedAt = time.Now().Unix()
	stmt, err := u.ModelManager.Prepare("update users_ifo set avatar=?,bg=?,about=?,labels=?,updated_at=? where id=?")

	defer stmt.Close()
	defer u.CloseDb()

	if err != nil {
		return false
	}

	res, err := stmt.Exec(u.Info.Avatar, u.Info.Bg, u.Info.About, u.Info.Labels, u.Info.UpdatedAt, u.Info.ID)
	if err != nil {
		return false
	}

	affect, _ := res.RowsAffected()
	fmt.Println(affect)
	return true
}
