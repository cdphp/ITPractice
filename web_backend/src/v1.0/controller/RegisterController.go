package controller

import (
	"v1.0/model"
	"v1.0/vendor"
)

// RegisterController struct
type RegisterController struct {
	vendor.Controller
}

// Index 注册
func (c *RegisterController) Index() {
	postData := c.GetPosts()

	result := new(Result)

	if HasParam(postData, "username") == false || HasParam(postData, "password") == false || HasParam(postData, "email") == false {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user := model.NewUser()
	user.Username = postData["username"]
	if !user.CheckName() {
		result.ErrorNo = 103
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user.Email = postData["email"]
	if !user.CheckEmail() {
		result.ErrorNo = 104
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	user.Password = postData["password"]
	user.Type = 1

	if user.Register() {
		// 获取token
		token := model.NewToken()
		auth := user.GetAuthName(user.Type)

		if token.Obtian(user.ID, user.Username, auth, 7200) {
			result.ErrorNo = 0
			sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

			sess.Set("token", token)
			result.Data = token
		} else {
			result.ErrorNo = 105
		}

	} else {
		result.ErrorNo = 106
	}

	JSONReturn(c.GetResponseWriter(), result)
	return
}
