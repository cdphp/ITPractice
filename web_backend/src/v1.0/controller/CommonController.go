package controller

import (
	"fmt"
	"strings"

	"v1.0/model"
	"v1.0/vendor"
)

// CommonController struct
type CommonController struct {
	vendor.Controller
}

// Mail 发送邮件
func (c *CommonController) Mail() {

	user := model.NewUser()
	result := new(Result)

	postData := c.GetPosts()
	if HasParam(postData, "type") == false || HasParam(postData, "email") == false {
		result.ErrorNo = 24
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	fmt.Println("encode:", postData["email"].(string))
	email, err := vendor.Base64Decode(postData["email"].(string))

	if err != nil {
		fmt.Println("err:", err)
		result.ErrorNo = 108
		JSONReturn(c.GetResponseWriter(), result)
		return
	}
	user.Email = email
	if !user.FindByEmail() {
		result.ErrorNo = 108
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

	sendType := postData["type"].(string)

	myConfig := new(vendor.Config)
	myConfig.InitConfig(getCurrentDirectory() + "/config/configs.ini")
	host := myConfig.Read("default", "host")

	if sendType == "validate" {
		to := user.Email
		subject := "验证邮箱"
		key := user.UniqueKey

		requestURI := "/reg/validate?key=" + key
		url := strings.Join([]string{host, requestURI}, "")

		body := `<!DOCTYPE html><html><head><meta charset="utf-8"><title>邮箱验证</title></head><style>.container{margin:0 auto;top:100px;position:relative;width:550px;height:300px;background:#fff;border-radius:5px;padding:30px}.container .content{padding:20px;font-size:14px;color:#666}.content div{margin:10px}</style><body><div style="background:#ebedeb;width:100%;height:600px"><div class="container"><div class="header"><div>尊敬的修行者 ` + to + ` ,您好：</div></div><div class="content"><div class="">感谢您加入我们。</div><div class="">请点击以下链接进行邮箱验证，以便开始使用您的账号</div><div class=""><a href="` + url + `" style="display:inline-block;color:#fff;line-height:40px;background-color:#1989fa;border-radius:5px;text-align:center;text-decoration:none;font-size:14px;padding:1px 30px">马上验证邮箱</a></div><div class="">如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：</div><div class=""><a href="` + url + `">` + url + `</a></div><div class="">如果您并未申请门派弟子，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者 <a href="#">联系掌门</a></div></div></div></div></body></html>		`

		fmt.Println("send email")
		err := SendToMail(to, subject, body, "html")
		if err != nil {
			fmt.Println("Send mail error!")
			fmt.Println(err)
			result.ErrorNo = 27
			JSONReturn(c.GetResponseWriter(), result)
			return
		}

		result.ErrorNo = 0
		result.Data = &model.User{Email: user.Email}
		JSONReturn(c.GetResponseWriter(), result)
		return
	}

}
