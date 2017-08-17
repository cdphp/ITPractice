package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"models"

	"github.com/gin-gonic/gin"
)

type CommentData struct {
	Content  string `json:"content" binding:"required"`
	Type     string `json:"type" binding:"required"`
	TargetID string `json:"target_id" binding:"required"`
	RootID   string `json:"root_id" binding:"required"`
}

func init() {
	db = Database()
}

// CreateComment 添加评论
func CreateComment(c *gin.Context) {
	var commentData CommentData

	if err := c.BindJSON(&commentData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	Type, err := strconv.ParseUint(commentData.Type, 10, 32)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}
	RootID, err := strconv.ParseInt(commentData.RootID, 10, 64)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}
	TargetID, err := strconv.ParseInt(commentData.TargetID, 10, 64)
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

	comment := models.Comment{
		Content:  commentData.Content,
		Type:     uint(Type),
		TargetID: TargetID,
		RootID:   RootID,
		UserID:   token.UserID,
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	// 奖励积分
	score := models.Score{
		UserID: token.UserID,
		Action: "comment",
		Num:    2,
		Type:   1,
	}
	if err := db.Create(&score).Error; err == nil {
		var profile models.Profile
		db.Where("user_id=?", token.UserID).First(&profile).UpdateColumn("score", profile.Score+score.Num)
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo":    errorNo,
		"message":    GetMsg(errorNo),
		"resourceId": comment.ID,
	})
}

// ListComment 文章列表
func ListComment(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	row := GetLimit()
	var comments []models.Comment
	var _comments []models.TransformedComment
	var total int

	targetID, err := strconv.Atoi(c.Query("target_id"))
	fmt.Println("target_id:", targetID)
	if err != nil {
		fmt.Println("err:", err)
		db.Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&comments)
		db.Model(&models.Comment{}).Count(&total)
	} else {
		db.Where("target_id=?", targetID).Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&comments)
		db.Model(&models.Comment{}).Where("target_id=?", targetID).Count(&total)
	}

	//transforms the users for building a good response
	for _, comment := range comments {
		var user models.User
		var profile models.Profile
		db.Model(&comment).Related(&user)
		db.Model(&user).Related(&profile)

		_comments = append(_comments, models.TransformedComment{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    user.ID,
			Author:    user.Username,
			Avatar:    profile.Avatar,
			Type:      comment.Type,
			TargetID:  comment.TargetID,
			RootID:    comment.RootID,
			CreatedAt: comment.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _comments,
		"total":   total,
	})
}

// DeleteComment 删除
func DeleteComment(c *gin.Context) {

}
