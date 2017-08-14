package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"v2.0/models"

	"github.com/gin-gonic/gin"
)

// UserData json
type UserData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser func
func CreateUser(c *gin.Context) {
	var userData UserData
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

	db := Database()
	if err := db.Create(&user).Error; err == nil {
		profile := models.Profile{
			UserID: user.ID,
		}
		db.Create(&profile)
		c.JSON(http.StatusCreated, gin.H{
			"errorNo":    0,
			"message":    "user created successfully!",
			"resourceId": user.ID,
		})

	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

}

// ListUser func
func ListUser(c *gin.Context) {
	page, _ := c.GetQuery("page")
	current, err := strconv.Atoi(page)
	if err != nil {
		current = 1
	}

	var profiles []models.Profile
	var _users []models.TransformedUser

	row := GetLimit()

	db := Database()
	db.Order("score desc").Offset((current - 1) * row).Limit(row).Find(&profiles)

	//transforms the users for building a good response
	for _, profile := range profiles {
		var user models.User
		db.Model(&profile).Related(&user)

		_users = append(_users, models.TransformedUser{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   profile.Avatar,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _users,
	})
}

// FetchSingleUser func
func FetchSingleUser(c *gin.Context) {

}

// UpdateUser func
func UpdateUser(c *gin.Context) {

}

// DeleteUser func
func DeleteUser(c *gin.Context) {

}
