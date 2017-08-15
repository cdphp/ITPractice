package main

import (
	"fmt"

	"v2.0/controllers"
	"v2.0/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//Migrate the schema
	db := Database()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.Token{})
	defer db.Close()

	router := gin.Default()
	common := router.Group("/api")
	{
		common.POST("/register", controllers.Register)
		common.POST("/login", controllers.Login)
	}

	users := router.Group("/api/user")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.ListUser)
		users.GET("/:id", controllers.FetchSingleUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
	router.Run(":8085")

}

// Database init func
func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:hongker@/it_practice2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
