package controller

import (
	"fmt"
	"strconv"

	"v1.0/model"
	"v1.0/vendor"
)

// UserController struct
type UserController struct {
	vendor.Controller
	operation *model.User
}

// Index function
func (c *UserController) Index() {
	user := model.NewUser()

	page, err := strconv.Atoi(c.GetQuery("page"))
	if err != nil {
		page = 1
	}

	limit := GetLimit()
	data := user.ListData(page, limit)

	result := new(Result)
	result.ErrorNo = 0
	result.Data = data

	JSONReturn(c.GetResponseWriter(), result)

}

// List function
func (c *UserController) List() {
	fmt.Fprintln(c.GetResponseWriter(), "Hello,welcome to user hello function")
	return

}

// Get function
func (c *UserController) Get() {
	user := model.NewUser()
	result := new(Result)

	params := c.GetParams()
	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user.ID = id
	if !user.Get() {
		result.ErrorNo = 22
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	fmt.Println(user)

	result.ErrorNo = 0
	result.Data = user

	JSONReturn(c.GetResponseWriter(), result)
	return
}

// Add func
func (c *UserController) Add() {
	postData := c.GetPosts()
	fmt.Println(postData)
}

// Update func
func (c *UserController) Update() {
	result := new(Result)
	params := c.GetParams()
	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	postData := c.GetPosts()
	user := model.NewUser()
	user.ID = id

	if HasParam(postData, "labels") {
		user.Info.Labels = postData["labels"].(string)
	}

	if HasParam(postData, "about") {
		user.Info.About = postData["about"].(string)
	}

	if !user.UpdateInfo() {
		result.ErrorNo = 25
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	result.ErrorNo = 0
	JSONReturn(c.GetResponseWriter(), result)
	return

}

// Delete func
func (c *UserController) Delete() {
	postData := c.GetPosts()
	fmt.Println(postData)
}
