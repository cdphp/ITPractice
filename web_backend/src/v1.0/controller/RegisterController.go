package controller

import (
	"v1.0/model"
	"v1.0/vendor"
)

// RegisterController struct
type RegisterController struct {
	vendor.Controller
	operation *model.User
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
	userType := 1
	c.operation = model.NewUser()
	user, errorNo := c.operation.Register(postData["username"], postData["email"], postData["password"], userType)

	result.ErrorNo = errorNo

	if errorNo == 0 {
		// 获取token
		tokenOperation := model.NewToken()
		auth := user.GetAuthName(user.Type)
		token, errorNo := tokenOperation.Obtian(user.ID, user.Username, auth, 7200)

		if errorNo == 0 {
			sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

			sess.Set("token", token)
			result.Data = token
		}

	}

	JSONReturn(c.GetResponseWriter(), result)
	defer c.operation.CloseDb()
	return
}
