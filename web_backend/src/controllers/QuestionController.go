package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"models"

	"github.com/gin-gonic/gin"
)

type QuestionData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func init() {
	db = Database()
}

// CreateQuestion 创建问题
func CreateQuestion(c *gin.Context) {
	var questionData QuestionData

	c.BindJSON(&questionData)

	if questionData.Title == "" || questionData.Content == "" {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var token models.Token
	token.Token = c.GetHeader("Token")

	if ValidateToken(&token, c) == false {
		errorNo := 201
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	question := models.Question{
		Title:   questionData.Title,
		Content: questionData.Content,
		UserID:  token.UserID,
	}

	if err := db.Create(&question).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo":    errorNo,
		"message":    GetMsg(errorNo),
		"resourceId": question.ID,
	})
}

// ListQuestion 问题列表
func ListQuestion(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	var questions []models.Question
	var _questions []models.TransformedQuestion
	var total int
	row := GetLimit()

	db.Where("is_delete=0").Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&questions)
	db.Model(&models.Question{}).Where("is_delete=0").Count(&total)

	//transforms the questions for building a good response
	for _, question := range questions {
		var user models.User
		var profile models.Profile
		db.Model(&question).Related(&user)
		db.Model(&user).Related(&profile)
		_questions = append(_questions, models.TransformedQuestion{
			ID:        question.ID,
			Title:     question.Title,
			Content:   question.Content,
			UserID:    user.ID,
			Author:    user.Username,
			Avatar:    profile.Avatar,
			CreatedAt: question.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _questions,
		"total":   total,
	})
}

// FetchSingleQuestion 获取信息
func FetchSingleQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var question models.Question
	var user models.User
	var profile models.Profile

	if err = db.Where("id=?", id).First(&question).Error; err != nil {

		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	db.Model(&question).Related(&user)
	db.Model(&user).Related(&profile)
	_question := models.TransformedQuestion{
		ID:        question.ID,
		Title:     question.Title,
		Content:   question.Content,
		UserID:    user.ID,
		Author:    user.Username,
		Avatar:    profile.Avatar,
		CreatedAt: question.CreatedAt,
	}
	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _question,
	})

}

// UpdateQuestion 修改信息
func UpdateQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看question是否存在
	var question models.Question
	if err = db.Where("id=?", id).First(&question).Error; err != nil {
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
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 判断是否有权修改
	if token.UserID != question.UserID {
		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}
	var questionData QuestionData
	c.BindJSON(&questionData)

	var updateQuestion models.Question

	if questionData.Title != "" {
		updateQuestion.Title = questionData.Title
	}

	if questionData.Content != "" {
		updateQuestion.Content = questionData.Content
	}

	if err = db.Model(&question).UpdateColumns(updateQuestion).Error; err != nil {
		errorNo := 25
		c.JSON(http.StatusNotAcceptable, gin.H{
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

// DeleteQuestion 删除
func DeleteQuestion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看question是否存在
	var question models.Question
	if err = db.Where("id=?", id).First(&question).Error; err != nil {
		errorNo := 22
		c.JSON(http.StatusCreated, gin.H{
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
	if token.UserID != question.UserID {
		errorNo := 22
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if err := db.Model(&question).UpdateColumn("is_delete", true).Error; err != nil {
		fmt.Println("err:", err)
		errorNo := 26
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
