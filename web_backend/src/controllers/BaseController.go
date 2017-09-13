package controllers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"lib"
	"models"

	seelog "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var logger seelog.LoggerInterface
var globalSessions *lib.Manager

func init() {
	db = Database()

	logger, _ = seelog.LoggerFromConfigAsFile(lib.GetCurrentDir() + "/configs/seelog.xml")
	logger.Info("Start Working")
	defer logger.Flush()

	globalSessions, _ = lib.NewSessionManager("memory", "goSessionid", 3600)
	go globalSessions.GC()
}

func Database() *gorm.DB {
	myConfig := new(lib.Config)
	myConfig.InitConfig(lib.GetCurrentDir() + "/configs/configs.ini")
	host := myConfig.Read("database", "host")
	port := myConfig.Read("database", "port")
	user := myConfig.Read("database", "user")
	password := myConfig.Read("database", "password")
	dbname := myConfig.Read("database", "dbname")

	//open a db connection
	orm, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return orm
}

const (
	//BASE64字符表,不要有重复
	base64Table        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	hashFunctionHeader = "hong"
	hashFunctionFooter = "ker"
)

// Substr 截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

// GetAuthName 获取等级名称
func GetAuthName(auth uint) string {
	var name string
	if auth == 1 {
		name = "User"
	} else if auth == 2 {
		name = "Master"
	} else if auth == 3 {
		name = "Custodian"
	} else if auth == 4 {
		name = "Manager"
	} else if auth == 22 {
		name = "Higher"
	}
	return name
}

// Md5 加密
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// UniqueID 生成Guid字串
func UniqueID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

var coder = base64.NewEncoding(base64Table)

// Base64Encode base64加密
func Base64Encode(str string) string {
	var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

// Base64Decode base64解密
func Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}

// Now return unix time
func Now() int64 {
	return time.Now().Unix()
}

// GetMsg 根据no获取对应msg
func GetMsg(no int) string {
	myConfig := new(lib.Config)

	myConfig.InitConfig(lib.GetCurrentDir() + "/configs/configs.ini")

	msg := myConfig.Read("error", strconv.Itoa(no))

	return msg
}

// IsEmail 验证邮箱格式
func IsEmail(email string) bool {
	b, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	return b
}

// GetLimit return limit
func GetLimit() int {
	return 10
}

// SendMail 发送邮件
func SendMail(to string, content map[string]string) {
	mail := lib.Mail{
		User:     GetDbConfig("mail_user"),
		Password: GetDbConfig("mail_pass"),
		Host:     GetDbConfig("mail_host"),
	}
	fmt.Println(mail)

	myConfig := new(lib.Config)
	myConfig.InitConfig(lib.GetCurrentDir() + "/configs/configs.ini")
	host := myConfig.Read("default", "host")
	var subject, body string

	if content["type"] == "register" {
		subject = "会员注册邮箱验证"
		requestURI := "/validate?type=" + content["type"] + "&key=" + content["key"]
		url := strings.Join([]string{host, requestURI}, "")

		body = `<!DOCTYPE html><html><head><meta charset="utf-8"><title>邮箱验证</title></head><style>.container{margin:0 auto;top:100px;position:relative;width:550px;height:300px;background:#fff;border-radius:5px;padding:30px}.container .content{padding:20px;font-size:14px;color:#666}.content div{margin:10px}</style><body><div style="background:#ebedeb;width:100%;height:600px"><div class="container"><div class="header"><div>尊敬的道友，` + to + ` ,您好：</div></div><div class="content"><div class="">感谢您加入我们。</div><div class="">请点击以下链接进行邮箱验证，以便开始使用您的账号</div><div class=""><a href="` + url + `" style="display:inline-block;color:#fff;line-height:40px;background-color:#1989fa;border-radius:5px;text-align:center;text-decoration:none;font-size:14px;padding:1px 30px">马上验证邮箱</a></div><div class="">如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：</div><div class=""><a href="` + url + `">` + url + `</a></div><div class="">如果您并未申请门派弟子，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者 <a href="#">联系掌门</a></div></div></div></div></body></html>		`

	} else if content["type"] == "forget" {
		subject = "找回密码邮箱验证"

		requestURI := "/validate?type=" + content["type"] + "&key=" + content["key"]
		url := strings.Join([]string{host, requestURI}, "")

		body = `<!DOCTYPE html><html><head><meta charset="utf-8"><title>邮箱验证</title></head><style>.container{margin:0 auto;top:100px;position:relative;width:550px;height:300px;background:#fff;border-radius:5px;padding:30px}.container .content{padding:20px;font-size:14px;color:#666}.content div{margin:10px}</style><body><div style="background:#ebedeb;width:100%;height:600px"><div class="container"><div class="header"><div>尊敬的道友，` + to + ` ,您好：</div></div><div class="content"><div class="">您正在进行找回密码操作。</div><div class="">请点击以下链接进行邮箱验证，以便开始使用您的账号</div><div class=""><a href="` + url + `" style="display:inline-block;color:#fff;line-height:40px;background-color:#1989fa;border-radius:5px;text-align:center;text-decoration:none;font-size:14px;padding:1px 30px">马上验证邮箱</a></div><div class="">如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：</div><div class=""><a href="` + url + `">` + url + `</a></div><div class="">如果您并未申请门派弟子，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者 <a href="#">联系掌门</a></div></div></div></div></body></html>		`

	}

	err := mail.SendToMail(to, subject, body, "html")

	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)

	}

	fmt.Println("Send mail success!")

}

// GetDbConfig 获取数据库里的配置项
func GetDbConfig(name string) string {
	var config models.Config
	db.Where("name=?", name).First(&config)

	return config.Value
}

// ValidateToken 验证token
func ValidateToken(token *models.Token, c *gin.Context) bool {
	sess := globalSessions.SessionStart(c.Writer, c.Request)

	sessionRes := sess.Get("token")

	fmt.Println("session:", sessionRes)
	if sessionRes == nil {
		if err := db.Where("token=?", token.Token).First(&token).Error; err != nil {
			fmt.Println("err", err)
			return false
		}
		fmt.Println("token", token)
	} else {
		token = sessionRes.(*models.Token)

		// token过期
		if Now()-token.CreatedAt > int64(token.Expire) {
			sess.Delete("token")
			return false
		}
	}

	return true
}

// Upload 上传文件
func UploadFile(filepath, key string) (string, error) {
	uploader := lib.Uploader{
		AccessKey: GetDbConfig("qiniu_accessKey"),
		SecretKey: GetDbConfig("qiniu_secretKey"),
		Bucket:    GetDbConfig("qiniu_bucket"),
	}

	res, err := uploader.Upload(filepath, key)
	url := GetDbConfig("qiniu_host") + res
	return url, err
}
