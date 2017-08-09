package controller

import (
	"fmt"

	"v1.0/model"
	"v1.0/vendor"
)

// LoginController struct
type LoginController struct {
	vendor.Controller
}

// Initialize 初始化
func (c *LoginController) Initialize() {
	fmt.Println("login controller intialize ")

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

	user := model.NewUser()
	user.Username = postData["username"].(string)

	var errorNo int
	if user.Auth(postData["password"].(string)) {
		// 获取token
		token := model.NewToken()
		auth := user.GetAuthName(user.Type)

		if token.Obtian(user.ID, user.Username, auth, 7200) {
			errorNo = 0
			sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

			sess.Set("token", token)
			result.Data = token
		} else {
			errorNo = 105
		}

	} else {
		errorNo = 102
	}

	result.ErrorNo = errorNo
	JSONReturn(c.GetResponseWriter(), result)
	return
}
