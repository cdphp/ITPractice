package model

import (
	"fmt"
	"strconv"
	"time"

	"v1.0/vendor"
)

// Article 文章信息
type Article struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Digest       string `json:"digest"`
	Content      string `json:"content"`
	UserID       int64  `json:"user_id"`
	Labels       string `json:"labels"`
	Clicks       int    `json:"clicks"`
	IsDelete     int    `json:"-"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	AuthorName   string `json:"author_name"`
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

// ListData 列表
func (a *Article) ListData(conditions map[string]string, page, row int) Articles {
	conditions["columns"] = "id,title,digest,content,user_id,created_at"

	start := (page - 1) * row
	conditions["limit"] = strconv.Itoa(start) + "," + strconv.Itoa(row)

	sql := a.SetConditions(conditions)
	fmt.Println(sql)
	rows, err := a.ModelManager.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	var result Articles
	for rows.Next() {

		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Digest, &article.Content, &article.UserID, &article.CreatedAt)
		if err == nil {
			result = append(result, article)
		} else {
			fmt.Println(err)
		}

	}
	defer rows.Close()

	return result
}

// Get 根据id获取数据
func (a *Article) Get() bool {

	sql := "select title,digest,content,user_id,labels,clicks,created_at from " + a.Resource + " where id=?"
	err := a.ModelManager.QueryRow(sql, a.ID).Scan(&a.Title, &a.Digest, &a.Content, &a.UserID, &a.Labels, &a.Clicks, &a.CreatedAt)

	if err != nil {
		return false
	}

	return true

}

// Add 添加
func (a *Article) Add() bool {
	a.CreatedAt = time.Now().Unix()
	//插入数据
	stmt, err := a.ModelManager.Prepare("INSERT " + a.Resource + " SET title=?,digest=?,content=?,user_id=?,created_at=?")

	res, err := stmt.Exec(a.Title, a.Digest, a.Content, a.UserID, a.CreatedAt)

	id, err := res.LastInsertId()

	defer stmt.Close()

	if err == nil {
		score := NewScore()
		user := NewUser()

		score.UserID = a.UserID
		score.Action = "write"
		score.Num = 10
		score.Type = 1

		user.ID = a.UserID

		if score.Add() && user.Upgrade(score.Num) {
			fmt.Println("奖励发送成功")
		}

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

	_, err = stmt.Exec(1, a.ID)
	if err != nil {
		return false
	}

	return true
}
