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
	fmt.Fprintln(c.GetResponseWriter(), "Hello,welcome to user module")
	return

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
	uid, _ := strconv.Atoi(params[1])

	user, errorNo := c.operation.Get(uid)

	result := new(Result)
	result.ErrorNo = errorNo
	result.Data = &user

	defer c.operation.CloseDb()

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
