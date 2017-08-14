package controller

import (
	"fmt"
	"strings"

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
	user.Username = postData["username"].(string)
	if !user.CheckName() {
		result.ErrorNo = 103
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user.Email = postData["email"].(string)
	if !user.CheckEmail() {
		result.ErrorNo = 104
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	user.Password = postData["password"].(string)
	user.Type = 1

	if user.Register() {
		if c.SendValidateMail(user.Email, user.UniqueKey) {
			result.ErrorNo = 0
		} else {
			result.ErrorNo = 107
		}

	} else {
		result.ErrorNo = 106
	}

	JSONReturn(c.GetResponseWriter(), result)
	return
}

// Validate 邮箱验证
func (c *RegisterController) Validate() {
	user := model.NewUser()
	result := new(Result)

	postData := c.GetPosts()
	if HasParam(postData, "key") == false {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user.UniqueKey = postData["key"].(string)
	if !user.FindByKey() {
		result.ErrorNo = 108
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	if user.State != 0 {
		result.ErrorNo = 109
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	user.State = 1
	if user.Update() {
		result.ErrorNo = 0
		//result.Data = user
	} else {
		result.ErrorNo = 25
	}
	JSONReturn(c.GetResponseWriter(), result)
	return

}

// SendValidateMail 发送验证邮箱
func (c *RegisterController) SendValidateMail(email, key string) bool {
	to := email
	subject := "验证邮箱"

	myConfig := new(vendor.Config)
	myConfig.InitConfig(getCurrentDirectory() + "/config/configs.ini")
	host := myConfig.Read("default", "host")

	requestURI := "/reg/validate?key=" + key
	url := strings.Join([]string{host, requestURI}, "")

	body := `<!DOCTYPE html><html><head><meta charset="utf-8"><title>邮箱验证</title></head><style>.container{margin:0 auto;top:100px;position:relative;width:550px;height:300px;background:#fff;border-radius:5px;padding:30px}.container .content{padding:20px;font-size:14px;color:#666}.content div{margin:10px}</style><body><div style="background:#ebedeb;width:100%;height:600px"><div class="container"><div class="header"><div>尊敬的修行者 ` + email + ` ,您好：</div></div><div class="content"><div class="">感谢您加入我们。</div><div class="">请点击以下链接进行邮箱验证，以便开始使用您的账号</div><div class=""><a href="` + url + `" style="display:inline-block;color:#fff;line-height:40px;background-color:#1989fa;border-radius:5px;text-align:center;text-decoration:none;font-size:14px;padding:1px 30px">马上验证邮箱</a></div><div class="">如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：</div><div class=""><a href="` + url + `">` + url + `</a></div><div class="">如果您并未申请门派弟子，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者 <a href="#">联系掌门</a></div></div></div></div></body></html>		`

	fmt.Println("send email")
	err := SendToMail(to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
		return false
	}

	fmt.Println("Send mail success!")
	return true
}
