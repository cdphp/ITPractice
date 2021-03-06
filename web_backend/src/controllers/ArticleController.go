package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"models"

	"github.com/gin-gonic/gin"
)

type ArticleData struct {
	Title   string `json:"title" binding:"required"`
	Digest  string `json:"digest" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func init() {
	db = Database()
}

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	var articleData ArticleData

	c.BindJSON(&articleData)

	if articleData.Title == "" || articleData.Digest == "" || articleData.Content == "" {
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

	article := models.Article{
		Title:   articleData.Title,
		Content: articleData.Content,
		Digest:  articleData.Digest,
		UserID:  token.UserID,
	}

	if err := db.Create(&article).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	// 奖励积分
	score := models.Score{
		UserID: token.UserID,
		Action: "article",
		Num:    10,
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
		"resourceId": article.ID,
	})
}

// ListArticle 文章列表
func ListArticle(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	var articles []models.Article
	var _articles []models.TransformedArticle
	var total int
	row := GetLimit()

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		db.Where("is_delete=0").Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&articles)
		db.Model(&models.Article{}).Where("is_delete=0").Count(&total)
	} else {
		db.Where("is_delete=0 and user_id=?", uid).Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&articles)
		db.Model(&models.Article{}).Where("is_delete=0 and user_id=?", uid).Count(&total)
	}

	//transforms the users for building a good response
	for _, article := range articles {
		var user models.User
		var profile models.Profile
		db.Model(&article).Related(&user)
		db.Model(&user).Related(&profile)
		_articles = append(_articles, models.TransformedArticle{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Digest:    article.Digest,
			UserID:    user.ID,
			Author:    user.Username,
			Avatar:    profile.Avatar,
			Labels:    article.Labels,
			Clicks:    article.Clicks,
			CreatedAt: article.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _articles,
		"total":   total,
	})
}

// FetchSingleArticle 获取信息
func FetchSingleArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	var article models.Article
	var user models.User
	var profile models.Profile

	if err = db.Where("id=?", id).First(&article).Error; err != nil {

		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	//增加点击量
	var token models.Token
	token.Token = c.GetHeader("Token")

	if !(ValidateToken(&token, c) == true && token.UserID == article.UserID) {
		db.Model(&article).UpdateColumn("clicks", article.Clicks+1)
	}

	db.Model(&article).Related(&user)
	db.Model(&user).Related(&profile)
	_article := models.TransformedArticle{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Digest:    article.Digest,
		UserID:    user.ID,
		Author:    user.Username,
		Avatar:    profile.Avatar,
		Labels:    article.Labels,
		Clicks:    article.Clicks,
		CreatedAt: article.CreatedAt,
	}
	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _article,
	})

}

// UpdateArticle 修改个人信息
func UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看article是否存在
	var article models.Article
	if err = db.Where("id=?", id).First(&article).Error; err != nil {
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
	if token.UserID != article.UserID {
		errorNo := 22
		c.JSON(http.StatusNoContent, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}
	var articleData ArticleData
	c.BindJSON(&articleData)

	var updateArticle models.Article

	if articleData.Title != "" {
		updateArticle.Title = articleData.Title
	}

	if articleData.Digest != "" {
		updateArticle.Digest = articleData.Digest
	}

	if articleData.Content != "" {
		updateArticle.Content = articleData.Content
	}

	if err = db.Model(&article).UpdateColumns(updateArticle).Error; err != nil {
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

// DeleteArticle 删除
func DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	// 查看article是否存在
	var article models.Article
	if err = db.Where("id=?", id).First(&article).Error; err != nil {
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
	if token.UserID != article.UserID {
		errorNo := 22
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	if err := db.Model(&article).UpdateColumn("is_delete", true).Error; err != nil {
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
