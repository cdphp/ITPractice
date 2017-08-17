package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"models"

	"github.com/gin-gonic/gin"
)

func init() {
	db = Database()
}

// UserData 用于接收row data json
type UserData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var userData UserData

	// 解析row data
	c.BindJSON(&userData)

	if userData.Username == "" || userData.Email == "" || userData.Password == "" {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if len(userData.Username) < 6 || len(userData.Username) > 30 {
		errorNo := 111
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if IsEmail(userData.Email) == false {
		errorNo := 112
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	user := models.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: Md5(userData.Password),
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

// ListUser 用户列表
func ListUser(c *gin.Context) {

	current, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		current = 1
	}

	var profiles []models.Profile
	var _users []models.TransformedUser

	row, err := strconv.Atoi(c.Query("row"))
	if err != nil {
		row = GetLimit()
	}

	db.Order("score desc").Offset((current - 1) * row).Limit(row).Find(&profiles)

	//transforms the users for building a good response
	for _, profile := range profiles {
		var user models.User
		db.Model(&profile).Related(&user)

		_users = append(_users, models.TransformedUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Avatar:   profile.Avatar,
			About:    profile.About,
			Labels:   profile.Labels,
			Score:    profile.Score,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _users,
	})
}

// FetchSingleUser 获取个人信息
func FetchSingleUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": 24,
			"message": err.Error(),
		})
		return
	}

	var user models.User
	var profile models.Profile

	if err = db.Where("id=?", id).First(&user).Related(&profile).Error; err != nil {

		fmt.Print("err:", err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": 22,
			"message": err.Error(),
		})
		return
	}

	_user := models.TransformedUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   profile.Avatar,
		About:    profile.About,
		Labels:   profile.Labels,
		Score:    profile.Score,
	}
	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _user,
	})

}

// UpdateUser 修改个人信息
func UpdateUser(c *gin.Context) {
	type UpdateData struct {
		Labels string `json:"labels" binding:"required"`
		About  string `json:"about" binding:"required"`
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看user是否存在
	var user models.User
	if err = db.Where("id=?", id).First(&user).Error; err != nil {
		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 验证token
	var token models.Token
	token.Token = c.GetHeader("Token")

	if ValidateToken(&token, c) == false {
		errorNo := 201
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 判断是否有权修改
	if token.UserID != id {
		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var updateData UpdateData
	c.BindJSON(&updateData)

	if err = db.Model(models.Profile{}).Where("user_id=?", token.UserID).UpdateColumns(models.Profile{Labels: updateData.Labels, About: updateData.About}).Error; err != nil {
		errorNo := 25
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
	})
	return
}

// DeleteUser 删除
func DeleteUser(c *gin.Context) {

}
