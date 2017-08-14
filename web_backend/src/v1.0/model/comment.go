package model

import (
	"fmt"
	"strconv"
	"time"

	"v1.0/vendor"
)

// Comment struct
type Comment struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	Content      string `json:"content"`
	Type         int    `json:"type"`
	TargetID     int64  `json:"target_id"`
	RootID       int64  `json:"root_id"`
	IsDelete     int    `json:"-"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Author       *User  `json:"author"`
	vendor.Model `json:"-"`
}

// Comments array
type Comments []Comment

// NewComment 初始化
func NewComment() *Comment {
	c := new(Comment)
	c.Init("comments") //设置表名

	return c
}

// ListData 列表
func (c *Comment) ListData(conditions map[string]string, page, row int) Comments {
	conditions["columns"] = "id,user_id,content,target_id,root_id,created_at"

	start := (page - 1) * row
	conditions["limit"] = strconv.Itoa(start) + "," + strconv.Itoa(row)
	conditions["order"] = "created_at desc"

	sql := c.SetConditions(conditions)
	fmt.Println(sql)
	rows, err := c.ModelManager.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	var result Comments

	for rows.Next() {

		var comment Comment
		err = rows.Scan(&comment.ID, &comment.UserID, &comment.Content, &comment.TargetID, &comment.RootID, &comment.CreatedAt)
		if err == nil {
			user := NewUser()
			user.ID = comment.UserID
			if user.Get() {
				comment.Author = user
			}
			result = append(result, comment)
		} else {
			fmt.Println(err)
		}

	}
	defer rows.Close()
	defer c.CloseDb()
	return result
}

// Add func
func (c *Comment) Add() bool {
	if c.Type == 0 {
		c.Type = 1
	}
	c.CreatedAt = time.Now().Unix()
	//插入数据
	stmt, err := c.ModelManager.Prepare("INSERT " + c.Resource + " SET user_id=?,content=?,type=?,target_id=?,root_id=?,created_at=?")

	res, err := stmt.Exec(c.UserID, c.Content, c.Type, c.TargetID, c.RootID, c.CreatedAt)

	id, err := res.LastInsertId()

	defer stmt.Close()
	defer c.CloseDb()

	if err == nil {
		score := NewScore()
		user := NewUser()

		score.UserID = c.UserID
		score.Action = "comment"
		score.Num = 2
		score.Type = 1

		user.ID = c.UserID

		if score.Add() && user.Upgrade(score.Num) {
			fmt.Println("奖励发送成功")
		}

		c.ID = id
		return true
	}
	return false
}
