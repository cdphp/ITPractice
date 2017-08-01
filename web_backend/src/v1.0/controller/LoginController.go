package controller

import (
	"fmt"

	"v1.0/model"
	"v1.0/vendor"
)

// LoginController struct
type LoginController struct {
	vendor.Controller
	operation *model.User
}

// Index 登录验证
func (c *LoginController) Index() {
	postData := c.GetPosts()

	result := new(Result)

	if HasParam(postData, "username") == false || HasParam(postData, "password") == false {
		result.ErrorNo = 24

		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	c.operation = model.NewUser()

	user, errorNo := c.operation.Auth(postData["username"], postData["password"])
	if errorNo == 0 {
		sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

		sess.Set("id", user.ID)
		fmt.Println("session set id:", user.ID)
	}
	result.ErrorNo = errorNo
	result.Data = user

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return
}
