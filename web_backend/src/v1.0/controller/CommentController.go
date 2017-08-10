package controller

import (
	"fmt"
	"strconv"

	"v1.0/model"
	"v1.0/vendor"
)

// CommentController struct
type CommentController struct {
	vendor.Controller
}

// Initialize 初始化
func (c *CommentController) Initialize() {
	fmt.Println("Initialize comment controller")

}

// Index 列表接口
func (c *CommentController) Index() {
	comment := model.NewComment()

	page, err := strconv.Atoi(c.GetQuery("page"))
	if err != nil {
		page = 1
	}

	conditions := make(map[string]string)
	conditions["conditions"] = "is_delete=0"

	_, err = strconv.Atoi(c.GetQuery("target_id"))
	if err == nil {
		conditions["conditions"] += " and target_id=" + c.GetQuery("target_id")
	}

	limit := GetLimit()
	comments := comment.ListData(conditions, page, limit)

	result := new(Result)
	result.ErrorNo = 0
	result.Data = comments

	JSONReturn(c.GetResponseWriter(), result)
}

// Add 添加接口
func (c *CommentController) Add() {
	result := new(Result)

	tokenStr := c.GetRequest().Header.Get("token")

	tokenObj, errorNo := ValidateToken(tokenStr, c.GetResponseWriter(), c.GetRequest())
	if errorNo != 0 {
		result.ErrorNo = errorNo
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	postData := c.GetPosts()

	if HasParam(postData, "target_id") == false || HasParam(postData, "content") == false {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	comment := model.NewComment()
	var err error
	comment.TargetID, err = strconv.ParseInt(postData["target_id"].(string), 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	comment.Content = postData["content"].(string)
	comment.UserID = tokenObj.UserID

	if HasParam(postData, "type") {
		comment.Type = postData["type"].(int)
	}

	if HasParam(postData, "root_id") {
		comment.RootID = postData["root_id"].(int64)
	}

	if !comment.Add() {
		result.ErrorNo = 23
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	result.ErrorNo = 0
	result.Data = comment

	JSONReturn(c.GetResponseWriter(), result)
	return
}
