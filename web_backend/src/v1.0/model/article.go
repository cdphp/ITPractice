package model

import (
	"fmt"
	"time"

	"v1.0/vendor"
)

// Article 文章信息
type Article struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Digest    string `json:"digest"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	Labels    string `json:"labels"`
	Clicks    int    `json:"clicks"`
	IsDelete  int    `json:"-"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`

	vendor.Model `json:"-"`
}

// Articles array
type Articles []Article

// NewArticle 初始化
func NewArticle() *Article {
	a := new(Article)
	a.Init("articles") //设置表名

	return a
}

// Get 根据id获取数据
func (a *Article) Get(id int) (*Article, int) {

	sql := "select id,title,digest,content,user_id,labels,clicks,created_at from " + a.Resource + " where id=?"
	err := a.ModelManager.QueryRow(sql, id).Scan(&a.ID, &a.Title, &a.Digest, &a.Content, &a.UserID, &a.Labels, &a.Clicks, &a.CreatedAt)

	if err != nil {
		return nil, 101
	}

	return a, 0

}

// Write 添加
func (a *Article) Write(title, digest, content string, userID int64) (*Article, int) {
	a.Title = title
	a.Digest = digest
	a.Content = content
	a.UserID = userID
	a.CreatedAt = time.Now().Unix()

	if a.Add() {
		return a, 0
	}
	return nil, 23

}

// Add 添加
func (a *Article) Add() bool {

	//插入数据
	stmt, err := a.ModelManager.Prepare("INSERT " + a.Resource + " SET title=?,digest=?,content=?,user_id=?,created_at=?")

	res, err := stmt.Exec(a.Title, a.Digest, a.Content, a.UserID, a.CreatedAt)

	id, err := res.LastInsertId()

	defer stmt.Close()
	defer a.CloseDb()

	if err == nil {
		a.ID = id
		return true
	}
	return false

}

// Update 更新
func (a *Article) Update() bool {
	a.UpdatedAt = time.Now().Unix()
	stmt, err := a.ModelManager.Prepare("update " + a.Resource + " set title=?,digest=?,content=?,labels=?,clicks=?,updated_at=? where id=?")

	defer stmt.Close()
	defer a.CloseDb()

	if err != nil {
		return false
	}

	res, err := stmt.Exec(a.Title, a.Digest, a.Content, a.Labels, a.Clicks, a.UpdatedAt, a.ID)
	if err != nil {
		return false
	}

	affect, _ := res.RowsAffected()
	fmt.Println(affect)
	return true
}

// Delete 删除
func (a *Article) Delete() bool {

	stmt, err := a.ModelManager.Prepare("update " + a.Resource + " set is_delete=? where id=?")
	if err != nil {
		return false
	}
	defer stmt.Close()
	defer a.CloseDb()

	_, err = stmt.Exec(1, a.ID)
	if err != nil {
		return false
	}

	return true
}
