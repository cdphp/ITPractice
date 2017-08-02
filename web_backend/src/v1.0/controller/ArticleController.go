package controller

import (
	"fmt"
	"strconv"

	"v1.0/model"
	"v1.0/vendor"
)

// ArticleController struct
type ArticleController struct {
	vendor.Controller
	operation *model.Article
}

// Initialize 初始化
func (c *ArticleController) Initialize() {
	fmt.Println("Initialize article controller")
}

// Index 列表接口
func (c *ArticleController) Index() {

}

// Get 详情接口
func (c *ArticleController) Get() {
	c.operation = model.NewArticle()

	params := c.GetParams()
	id, _ := strconv.Atoi(params[1])

	article, errorNo := c.operation.Get(id)

	result := new(Result)
	result.ErrorNo = errorNo
	result.Data = article

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return
}

// Add 添加接口
func (c *ArticleController) Add() {
	c.operation = model.NewArticle()

	postData := c.GetPosts()

	result := new(Result)

	if HasParam(postData, "title") == false || HasParam(postData, "digest") == false || HasParam(postData, "content") == false {
		result.ErrorNo = 24

		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

	token := sess.Get("token").(*model.Token)

	article, errorNo := c.operation.Write(postData["title"], postData["digest"], postData["content"], token.UserID)

	result.ErrorNo = errorNo
	result.Data = article

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return
}

// Update 更新接口
func (c *ArticleController) Update() {

}

// Delete 删除接口
func (c *ArticleController) Delete() {

}
