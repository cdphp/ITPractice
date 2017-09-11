package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"models"

	"github.com/gin-gonic/gin"
)

type AnswerData struct {
	Content    string `json:"content" binding:"required"`
	QuestionID string `json:"question_id" binding:"required"`
}

func init() {
	db = Database()
}

// CreateAnswer 添加回答
func CreateAnswer(c *gin.Context) {
	var answerData AnswerData

	if err := c.BindJSON(&answerData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	QuestionID, err := strconv.ParseInt(answerData.QuestionID, 10, 64)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

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

	answer := models.Answer{
		Content:    answerData.Content,
		QuestionID: QuestionID,
		UserID:     token.UserID,
	}

	if err := db.Create(&answer).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo":    errorNo,
		"message":    GetMsg(errorNo),
		"resourceId": answer.ID,
	})
}

// ListAnswer 回答列表
func ListAnswer(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	row := GetLimit()
	var answers []models.Answer
	var _answers []models.TransformedAnswer
	var total int

	questionID, err := strconv.Atoi(c.Query("question_id"))
	fmt.Println("question_id:", questionID)
	if err != nil {
		fmt.Println("err:", err)
		db.Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&answers)
		db.Model(&models.Answer{}).Count(&total)
	} else {
		db.Where("question_id=?", questionID).Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&answers)
		db.Model(&models.Answer{}).Where("question_id=?", questionID).Count(&total)
	}

	//transforms the users for building a good response
	for _, answer := range answers {
		var user models.User
		var profile models.Profile
		db.Model(&answer).Related(&user)
		db.Model(&user).Related(&profile)

		_answers = append(_answers, models.TransformedAnswer{
			ID:         answer.ID,
			Content:    answer.Content,
			UserID:     user.ID,
			Author:     user.Username,
			Avatar:     profile.Avatar,
			QuestionID: answer.QuestionID,
			CreatedAt:  answer.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _answers,
		"total":   total,
	})
}

// DeleteAnswer 删除
func DeleteAnswer(c *gin.Context) {

}
