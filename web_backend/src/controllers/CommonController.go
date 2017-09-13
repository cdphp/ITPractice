package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"models"

	"lib"

	"github.com/gin-gonic/gin"
)

func init() {
	db = Database()
}

// LoginData 用于接收登录的row data json
type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegUserData 用于接收注册的row data json
type RegUserData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ValidateData struct {
	Key string `json:"key" binding:"required"`
}

// Register 创建用户
func Register(c *gin.Context) {
	var regUserData RegUserData

	// 解析row data
	if err := c.BindJSON(&regUserData); err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if len(regUserData.Username) < 6 || len(regUserData.Username) > 30 {
		errorNo := 111
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if IsEmail(regUserData.Email) == false {
		errorNo := 112
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if HasUser(regUserData.Username) {
		errorNo := 103
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if HasEmail(regUserData.Email) {
		errorNo := 104
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	user := models.User{
		Username:  regUserData.Username,
		Email:     regUserData.Email,
		Password:  Md5(regUserData.Password),
		UniqueKey: UniqueID(),
	}
	fmt.Println(user)

	if err := db.Create(&user).Error; err == nil {
		profile := models.Profile{
			UserID: user.ID,
			Avatar: "http://ouecw69lw.bkt.clouddn.com/profile_big.jpg",
		}
		db.Create(&profile)

		// 发送验证邮件
		content := make(map[string]string)
		content["type"] = "validate"
		content["key"] = user.UniqueKey
		SendMail(user.Email, content)

		errorNo := 0
		c.JSON(http.StatusCreated, gin.H{
			"errorNo":    errorNo,
			"message":    GetMsg(errorNo),
			"resourceId": user.ID,
		})

	} else {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

}

// HasUser 检查用户名是否存在
func HasUser(username string) bool {

	var user models.User
	db.Where("username=?", username).First(&user)

	if user.ID != 0 {
		return true
	}
	return false
}

// HasEmail 检查邮箱是否存在
func HasEmail(email string) bool {

	var user models.User
	db.Where("email=?", email).First(&user)

	if user.ID != 0 {
		return true
	}
	return false
}

// Login 用户登录
func Login(c *gin.Context) {
	var loginData LoginData

	// 解析row data
	if err := c.BindJSON(&loginData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var user models.User
	//验证账号是邮箱还是用户名
	if IsEmail(loginData.Account) {
		db.Where("email=? and is_delete=0", loginData.Account).First(&user)
	} else {
		db.Where("username=? and is_delete=0", loginData.Account).First(&user)
	}

	if user.ID == 0 {
		errorNo := 101
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if Md5(loginData.Password) != user.Password {
		errorNo := 102
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if user.State == 0 { // 需要验证邮箱
		errorNo := 110
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
			"email":   user.Email,
		})
		return
	}

	//验证通过，生成token
	token := models.Token{
		Token:  Md5(strconv.FormatInt(time.Now().UnixNano(), 10)),
		UserID: user.ID,
		Auth:   GetAuthName(user.Type),
	}

	db.Create(&token)
	if token.ID != 0 {
		var user models.User
		var profile models.Profile
		db.Model(&token).Related(&user).Related(&profile)

		_token := models.TransformedToken{
			ID:        token.ID,
			Token:     token.Token,
			Username:  user.Username,
			UserID:    user.ID,
			Avatar:    profile.Avatar,
			Expire:    token.Expire,
			Auth:      token.Auth,
			CreatedAt: token.CreatedAt,
		}

		errorNo := 0
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
			"data":    _token,
		})

	} else {
		errorNo := 105
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
	}

}

// Validate 邮箱验证
func Validate(c *gin.Context) {
	var validateData ValidateData

	// 解析row data
	c.BindJSON(&validateData)

	if validateData.Key == "" {
		errorNo := 24
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var user models.User

	if err := db.Where("unique_key=?", validateData.Key).First(&user).Error; err != nil {
		errorNo := 101
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if user.State == 0 {
		if err := db.Model(&user).UpdateColumn(models.User{State: 1, UpdatedAt: Now()}).Error; err != nil {
			fmt.Println(err)
			errorNo := 25
			c.JSON(http.StatusBadGateway, gin.H{
				"errorNo": errorNo,
				"message": GetMsg(errorNo),
			})
			return
		}
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
	})
	return

}

func Mail(c *gin.Context) {
	type MailData struct {
		Type  string `json:"type" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	var mailData MailData

	if err := c.BindJSON(&mailData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看邮箱是否存在
	var user models.User

	if err := db.Where("email=?", mailData.Email).First(&user).Error; err != nil {
		errorNo := 113
		c.JSON(http.StatusOK, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 发送验证邮件
	content := make(map[string]string)
	content["type"] = mailData.Type
	content["key"] = user.UniqueKey
	SendMail(user.Email, content)

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
	})
	return
}

// Upload 上传
func Upload(c *gin.Context) {
	type UploadData struct {
		Type    string `json:"type" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	var uploadData UploadData

	// 解析row data
	if err := c.BindJSON(&uploadData); err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var localFile, key string

	if uploadData.Type == "image" { //上传图片

		key = Md5(strconv.FormatInt(time.Now().UnixNano(), 10))

		var baseName string
		baseName = "/tmp/" + key

		res, err := lib.SaveImageToDisk(baseName, uploadData.Content)
		if err != nil {
			fmt.Println(err)
			return
		}
		localFile = res
	}

	url, err := UploadFile(localFile, key)

	if err != nil {
		errorNo := 28
		c.JSON(http.StatusOK, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return

	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"url":     url,
	})
	return
}

// 忘记密码
func ForgetPass(c *gin.Context) {
	// ForgetData 用于接收忘记密码的row data
	type ForgetData struct {
		Account string `json:"account" binding:"required"`
	}
	var forgetData ForgetData

	// 解析row data
	if err := c.BindJSON(&forgetData); err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查找账户
	var user models.User
	//验证账号是邮箱还是用户名
	if IsEmail(forgetData.Account) {
		db.Where("email=? and is_delete=0", forgetData.Account).First(&user)
	} else {
		db.Where("username=? and is_delete=0", forgetData.Account).First(&user)
	}

	if user.ID == 0 {
		errorNo := 101
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"email":   user.Email,
	})
	return

}

// Oauth 验证
func Oauth(c *gin.Context) {
	code := c.Query("code")

	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader("client_id=a5d9a5d4595fdb831368&client_secret=3527e024d285fff4b6614af690acbb6fa010a08c&redirect_uri=http://127.0.0.1:8010/login&code="+code),
	)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var dat map[string]string
	json.Unmarshal(body, &dat)
	fmt.Println(dat)

	if _, ok := dat["access_token"]; ok {
		// 获取用户信息
		resp, err = http.Get("https://api.github.com/user?access_token=" + dat["access_token"])
		body, err = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err == nil {
			var userinfo map[string]interface{}
			json.Unmarshal(body, &userinfo)

			id := int(userinfo["id"].(float64))
			username := userinfo["name"].(string)
			email := userinfo["email"].(string)
			avatar := userinfo["avatar_url"].(string)

			var user models.User
			db.Where("github_id=? and is_delete=0", id).First(&user)
			// 检查用户是否存在
			if user.ID != 0 { //直接登录
				fmt.Println("user:", user)

				//验证通过，生成token
				token := models.Token{
					Token:  Md5(strconv.FormatInt(time.Now().UnixNano(), 10)),
					UserID: user.ID,
					Auth:   GetAuthName(user.Type),
				}

				db.Create(&token)
				if token.ID != 0 {
					var user models.User
					var profile models.Profile
					db.Model(&token).Related(&user).Related(&profile)

					_token := models.TransformedToken{
						ID:        token.ID,
						Token:     token.Token,
						Username:  user.Username,
						UserID:    user.ID,
						Avatar:    profile.Avatar,
						Expire:    token.Expire,
						Auth:      token.Auth,
						CreatedAt: token.CreatedAt,
					}

					errorNo := 0
					c.JSON(http.StatusCreated, gin.H{
						"errorNo": errorNo,
						"message": GetMsg(errorNo),
						"data":    _token,
					})
					return

				} else {
					errorNo := 105
					c.JSON(http.StatusCreated, gin.H{
						"errorNo": errorNo,
						"message": GetMsg(errorNo),
					})
					return
				}
			} else { //需要注册
				user := models.User{
					Username:  username,
					Email:     email,
					Password:  Md5(username),
					GithubID:  id,
					State:     1,
					UniqueKey: UniqueID(),
				}
				fmt.Println(user)

				if err := db.Create(&user).Error; err == nil {
					profile := models.Profile{
						UserID: user.ID,

						Avatar: avatar,
					}
					db.Create(&profile)

					errorNo := 0
					c.JSON(http.StatusCreated, gin.H{
						"errorNo":    errorNo,
						"message":    GetMsg(errorNo),
						"need_login": 1,
					})
					return

				} else {
					c.JSON(http.StatusCreated, gin.H{
						"errorNo": 23,
						"message": err.Error(),
					})
					return
				}
			}
		}

	}
	errorNo := 201
	c.JSON(http.StatusCreated, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
	})
	return

}
