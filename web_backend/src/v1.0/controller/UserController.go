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
	c.operation = model.NewUser()

	params := c.GetParams()
	id, _ := strconv.Atoi(params[1])

	user, errorNo := c.operation.Get(id)

	fmt.Println(user)
	result := new(Result)
	result.ErrorNo = errorNo
	result.Data = &user

	JSONReturn(c.GetResponseWriter(), result)

}

// Add func
func (c *UserController) Add() {
	postData := c.GetPosts()
	fmt.Println(postData)
}

// Update func
func (c *UserController) Update() {
	postData := c.GetPosts()
	fmt.Println(postData)
}

// Delete func
func (c *UserController) Delete() {
	postData := c.GetPosts()
	fmt.Println(postData)
}
