package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"v2.0/models"
)

var db *gorm.DB

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

// Register 创建用户
func Register(c *gin.Context) {
	var regUserData RegUserData

	// 解析row data
	c.BindJSON(&regUserData)

	if regUserData.Username == "" || regUserData.Email == "" || regUserData.Password == "" {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if len(regUserData.Username) < 6 || len(regUserData.Username) > 30 {
		errorNo := 111
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if IsEmail(regUserData.Email) == false {
		errorNo := 112
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if HasUser(regUserData.Username) {
		errorNo := 103
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if HasEmail(regUserData.Email) {
		errorNo := 104
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	user := models.User{
		Username: regUserData.Username,
		Email:    regUserData.Email,
		Password: Md5(regUserData.Password),
	}
	fmt.Println(user)

	if err := db.Create(&user).Error; err == nil {
		profile := models.Profile{
			UserID: user.ID,
			Avatar: "http://ouecw69lw.bkt.clouddn.com/profile_big.jpg",
		}
		db.Create(&profile)
		errorNo := 0
		c.JSON(http.StatusCreated, gin.H{
			"errorNo":    errorNo,
			"message":    GetMsg(errorNo),
			"resourceId": user.ID,
		})

	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
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
	c.BindJSON(&loginData)

	if loginData.Account == "" || loginData.Password == "" {
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
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if Md5(loginData.Password) != user.Password {
		errorNo := 102
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
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
			ID:       token.ID,
			Token:    token.Token,
			Username: user.Username,
			Avatar:   profile.Avatar,
			Expire:   token.Expire,
			Auth:     token.Auth,
		}

		errorNo := 0
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 105,
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

// GetAuthName 获取等级名称
func GetAuthName(auth uint) string {
	var name string
	if auth == 1 {
		name = "User"
	} else if auth == 2 {
		name = "Master"
	} else {
		name = "Admin"
	}
	return name
}
