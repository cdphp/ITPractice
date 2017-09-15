package controllers

import (
	"fmt"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MsgData 消息数据
type MsgData struct {
	Content  string `json:"content" binding:"required"`
	TargetID int    `json:"target_id" binding:"required"`
}

func init() {
	db = Database()
}

// CreateMessage 新增消息
func CreateMessage(c *gin.Context) {
	var msgData MsgData

	if err := c.BindJSON(&msgData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	TargetID := int64(msgData.TargetID)

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

	message := models.Message{
		UserID:   token.UserID,
		Content:  msgData.Content,
		TargetID: TargetID,
	}

	if err := db.Create(&message).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo":    errorNo,
		"message":    GetMsg(errorNo),
		"resourceId": message.ID,
	})

}

// ListMessage 消息列表
func ListMessage(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	row := GetLimit()

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
	var author models.User
	var authorProfile models.Profile
	db.Model(&token).Related(&author)
	db.Model(&author).Related(&authorProfile)

	var messages []models.Message
	var _messages []models.TransformedMessage

	db.Where("target_id=?", token.UserID).Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&messages)

	//transforms the users for building a good response
	for _, message := range messages {
		var user models.User
		var profile models.Profile
		db.Model(&message).Related(&user)
		db.Model(&user).Related(&profile)

		_messages = append(_messages, models.TransformedMessage{
			ID:           message.ID,
			Content:      message.Content,
			UserID:       message.UserID,
			AuthorName:   user.Username,
			AuthorAvatar: profile.Avatar,
			TargetID:     message.TargetID,
			TargetName:   author.Username,
			TargetAvatar: authorProfile.Avatar,
			CreatedAt:    message.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _messages,
	})
}
