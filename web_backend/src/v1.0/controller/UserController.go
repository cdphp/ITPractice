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
}

var operation *model.User

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
	operation = model.NewUser()

	params := c.GetParams()
	uid, _ := strconv.Atoi(params[1])

	user, err := operation.Get(uid)

	result := make(map[string]interface{})

	if err != nil {
		result["errorNo"] = err.No
		result["errorMsg"] = err.Msg
	} else {
		result["errorNo"] = 0
		result["errorMsg"] = "没有错误"
	}
	result["data"] = &user
	defer operation.CloseDb()
	c.JSONReturn(result)

}

// Add func
func (c *UserController) Add() {
	postData := c.GetPostData()
	fmt.Println(postData)
}

// Update func
func (c *UserController) Update() {
	postData := c.GetPostData()
	fmt.Println(postData)
}

// Delete func
func (c *UserController) Delete() {
	postData := c.GetPostData()
	fmt.Println(postData)
}
