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
}

// Initialize 初始化
func (c *ArticleController) Initialize() {
	fmt.Println("Initialize article controller")

}

// Index 列表接口
func (c *ArticleController) Index() {
	article := model.NewArticle()

	page, err := strconv.Atoi(c.GetQuery("page"))
	if err != nil {
		page = 1
	}

	conditions := make(map[string]string)
	conditions["conditions"] = "is_delete=0"
	_, err = strconv.Atoi(c.GetQuery("uid"))
	if err == nil {
		conditions["conditions"] += " and user_id=" + c.GetQuery("uid")
	}

	limit := GetLimit()
	data := article.ListData(conditions, page, limit)
	fmt.Println("data:", data)
	result := new(Result)
	result.ErrorNo = 0
	result.Data = data

	JSONReturn(c.GetResponseWriter(), result)

}

// Get 详情接口
func (c *ArticleController) Get() {
	article := model.NewArticle()
	result := new(Result)

	params := c.GetParams()
	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	article.ID = id
	if !article.Get() {
		result.ErrorNo = 22
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	user := model.NewUser()
	user.ID = article.UserID
	if !user.Get() {
		result.ErrorNo = 22
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	article.AuthorName = user.Username

	result.ErrorNo = 0
	result.Data = article

	JSONReturn(c.GetResponseWriter(), result)
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

	postData := c.GetPosts()

	if HasParam(postData, "title") == false || HasParam(postData, "digest") == false || HasParam(postData, "content") == false {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	article := model.NewArticle()
	article.Title = postData["title"].(string)
	article.Digest = postData["digest"].(string)
	article.Content = postData["content"].(string)
	article.UserID = tokenObj.UserID
	if !article.Add() {
		result.ErrorNo = 23
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	result.ErrorNo = 0
	result.Data = article

	JSONReturn(c.GetResponseWriter(), result)
	return
}

// Update 更新接口
func (c *ArticleController) Update() {
	article := model.NewArticle()
	result := new(Result)

	params := c.GetParams()
	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	postData := c.GetPosts()

	article.ID = id
	if !article.Get() {
		result.ErrorNo = 22
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	if HasParam(postData, "title") {
		article.Title = postData["title"].(string)
	}

	if HasParam(postData, "digest") {
		article.Digest = postData["digest"].(string)
	}

	if HasParam(postData, "content") {
		article.Content = postData["content"].(string)
	}

	if HasParam(postData, "labels") {
		article.Labels = postData["labels"].(string)
	}

	if HasParam(postData, "clicks") {
		article.Clicks, _ = strconv.Atoi(postData["clicks"].(string))
	}
	if article.Update() {
		result.ErrorNo = 0
	} else {
		result.ErrorNo = 25
	}

	JSONReturn(c.GetResponseWriter(), result)
	return
}

// Delete 删除接口
func (c *ArticleController) Delete() {
	article := model.NewArticle()
	result := new(Result)

	params := c.GetParams()
	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	article.ID = id
	if article.Delete() {
		result.ErrorNo = 0
	} else {
		result.ErrorNo = 26
	}

	JSONReturn(c.GetResponseWriter(), result)
	return

}
