package controllers

import (
	"fmt"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	db = Database()
}

// CompanyData 公司数据
type CompanyData struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateCompany 添加
func CreateCompany(c *gin.Context) {
	var companyData CompanyData

	if err := c.BindJSON(&companyData); err != nil {
		fmt.Println("err:", err)
		errorNo := 24
		c.JSON(http.StatusNotAcceptable, gin.H{
			"errorNo": errorNo,
			"message": GetMsg(errorNo),
		})
		return
	}

	company := models.Company{
		Name:        companyData.Name,
		Description: companyData.Description,
	}

	if err := db.Create(&company).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"errorNo": 23,
			"message": err.Error(),
		})
	}

	errorNo := 0
	c.JSON(http.StatusCreated, gin.H{
		"errorNo":    errorNo,
		"message":    GetMsg(errorNo),
		"resourceId": company.ID,
	})

}

// ListCompany 公司列表
func ListCompany(c *gin.Context) {

	current, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		current = 1
	}

	var companys []models.Company
	var _companys []models.TransformedCompany
	var total int
	row := GetLimit()

	name := c.Query("name")
	if name == "" {
		db.Where("is_delete=0").Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&companys)
		db.Model(&models.Company{}).Where("is_delete=0").Count(&total)
	} else {
		db.Where("is_delete=0 and name like ?", "%"+name+"%").Order("created_at desc").Offset((current - 1) * row).Limit(row).Find(&companys)
		db.Model(&models.Company{}).Where("is_delete=0 and name like ?", "%"+name+"%").Count(&total)
	}

	//transforms the companys for building a good response
	for _, company := range companys {

		_companys = append(_companys, models.TransformedCompany{
			ID:          company.ID,
			Name:        company.Name,
			Description: company.Description,
			CreatedAt:   company.CreatedAt,
		})
	}

	errorNo := 0
	c.JSON(http.StatusOK, gin.H{
		"errorNo": errorNo,
		"message": GetMsg(errorNo),
		"data":    _companys,
		"total":   total,
	})
}

func UpCompany(c *gin.Context) {
	type UpData struct {
		CompanyID int `json:"company_id" binding:"required"`
	}
}

func DownCompany(c *gin.Context) {

}
