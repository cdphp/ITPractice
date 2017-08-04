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
	c.operation = model.NewArticle()

	page, _ := strconv.Atoi(c.GetQuery("page"))

	conditions := make(map[string]string)
	conditions["conditions"] = "is_delete=0"

	data := c.operation.Find(conditions)
	limit := GetLimit()

	result := new(Result)
	result.ErrorNo = 0
	result.Data = Paginator(data, page, limit)

	JSONReturn(c.GetResponseWriter(), result)

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
	result := new(Result)

	tokenStr := c.GetRequest().Header.Get("token")

	tokenObj, errorNo := ValidateToken(tokenStr, c.GetResponseWriter(), c.GetRequest())
	if errorNo != 0 {
		result.ErrorNo = errorNo
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	c.operation = model.NewArticle()
	postData := c.GetPosts()

	if HasParam(postData, "title") == false || HasParam(postData, "digest") == false || HasParam(postData, "content") == false {
		result.ErrorNo = 24

		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	article, errorNo := c.operation.Write(postData["title"], postData["digest"], postData["content"], tokenObj.UserID)

	result.ErrorNo = errorNo
	result.Data = article

	JSONReturn(c.GetResponseWriter(), result)

	return
}

// Update 更新接口
func (c *ArticleController) Update() {
	c.operation = model.NewArticle()

	params := c.GetParams()
	id, _ := strconv.Atoi(params[1])

	postData := c.GetPosts()

	article, errorNo := c.operation.Get(id)

	result := new(Result)
	if errorNo != 0 {
		result.ErrorNo = errorNo
		JSONReturn(c.GetResponseWriter(), result)
		defer c.operation.CloseDb()
		return
	}

	if HasParam(postData, "title") {
		article.Title = postData["title"]
	}

	if HasParam(postData, "digest") {
		article.Digest = postData["digest"]
	}

	if HasParam(postData, "content") {
		article.Content = postData["content"]
	}

	if HasParam(postData, "labels") {
		article.Labels = postData["labels"]
	}

	if HasParam(postData, "clicks") {
		article.Clicks, _ = strconv.Atoi(postData["clicks"])
	}
	if article.Update() {
		result.ErrorNo = 0
	} else {
		result.ErrorNo = 25
	}

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return
}

// Delete 删除接口
func (c *ArticleController) Delete() {
	c.operation = model.NewArticle()

	params := c.GetParams()
	id, _ := strconv.Atoi(params[1])

	article, errorNo := c.operation.Get(id)

	result := new(Result)
	if errorNo != 0 {
		result.ErrorNo = errorNo
		JSONReturn(c.GetResponseWriter(), result)
		defer c.operation.CloseDb()
		return
	}

	if article.Delete() {
		result.ErrorNo = 0
	} else {
		result.ErrorNo = 26
	}

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return

}
