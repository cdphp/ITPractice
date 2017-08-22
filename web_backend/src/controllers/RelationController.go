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

// CreateRelation 创建文章
func CreateRelation(c *gin.Context) {
	type UserRelation struct {
		MasterID int64 `json:"master_id" binding:"required"`
	}

	var userRelation UserRelation
	if err := c.BindJSON(&userRelation); err != nil {
		fmt.Println(err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
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

	if token.UserID == userRelation.MasterID {
		errorNo := 203
		c.JSON(http.StatusOK, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var relation models.Relation

	if err := db.Where("is_delete=0 and master_id=? and pupil_id=?", userRelation.MasterID, token.UserID).First(&relation).Error; err == nil {
		//已有数据
		errorNo := 202
		c.JSON(http.StatusOK, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	relation.MasterID = userRelation.MasterID
	relation.PupilID = token.UserID

	if err := db.Create(&relation).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	// 奖励积分
	score := models.Score{
		UserID: token.UserID,
		Action: "relation",
		Num:    5,
		Type:   1,
	}
	if err := db.Create(&score).Error; err == nil {
		var profile models.Profile
		db.Where("user_id=?", token.UserID).First(&profile).UpdateColumn("score", profile.Score+score.Num)
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
	})
}

// ListRelation 创建文章
func ListRelation(c *gin.Context) {
	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	row := GetLimit()

	var relations []models.Relation
	var _relations []models.TransformedRelation
	var total int

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		db.Where("is_delete=0").Offset((current - 1) * row).Limit(row).Find(&relations)
		db.Model(&models.Relation{}).Where("is_delete=0").Count(&total)
	} else {
		db.Where("is_delete=0 and master_id=?", uid).Offset((current - 1) * row).Limit(row).Find(&relations)
		db.Model(&models.Relation{}).Where("is_delete=0 and master_id=?", uid).Count(&total)
	}

	//transforms the users for building a good response
	for _, relation := range relations {
		var user models.User
		var profile models.Profile
		db.Model(&relation).Related(&user, "Pupil")
		db.Model(&user).Related(&profile)
		_relations = append(_relations, models.TransformedRelation{
			ID:        relation.ID,
			MasterID:  relation.MasterID,
			PupilID:   relation.PupilID,
			Username:  user.Username,
			Avatar:    profile.Avatar,
			Labels:    profile.Labels,
			Type:      relation.Type,
			CreatedAt: relation.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _relations,
		"total":   total,
	})

}

// DeleteRelation 创建文章
func DeleteRelation(c *gin.Context) {

}
