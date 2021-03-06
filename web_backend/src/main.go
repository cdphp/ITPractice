package main

import (
	"fmt"
	"lib"

	"controllers"

	"models"

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
	db.AutoMigrate(&models.Config{})
	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Score{})
	db.AutoMigrate(&models.Relation{})
	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Answer{})
	db.AutoMigrate(&models.Message{})

	defer db.Close()

	router := gin.Default()

	common := router.Group("/")
	{
		common.POST("/register", controllers.Register)
		common.POST("/login", controllers.Login)
		common.POST("/validate", controllers.Validate)
		common.POST("/mail", controllers.Mail)
		common.POST("/upload", controllers.Upload)
		common.POST("/forget", controllers.ForgetPass)
		common.GET("/oauth", controllers.Oauth)
	}

	users := router.Group("/user")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.ListUser)
		users.GET("/:id", controllers.FetchSingleUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
		users.POST("/resetPass", controllers.ResetPass)
	}

	articles := router.Group("/article")
	{
		articles.POST("/", controllers.CreateArticle)
		articles.GET("/", controllers.ListArticle)
		articles.GET("/:id", controllers.FetchSingleArticle)
		articles.PUT("/:id", controllers.UpdateArticle)
		articles.DELETE("/:id", controllers.DeleteArticle)
	}

	questions := router.Group("/question")
	{
		questions.POST("/", controllers.CreateQuestion)
		questions.GET("/", controllers.ListQuestion)
		questions.GET("/:id", controllers.FetchSingleQuestion)
		questions.PUT("/:id", controllers.UpdateQuestion)
		questions.DELETE("/:id", controllers.DeleteQuestion)
	}

	comments := router.Group("/comment")
	{
		comments.POST("/", controllers.CreateComment)
		comments.GET("/", controllers.ListComment)
	}

	answers := router.Group("/answer")
	{
		answers.POST("/", controllers.CreateAnswer)
		answers.GET("/", controllers.ListAnswer)
		answers.POST("/evalute", controllers.EvaluateAnswer)
	}

	companys := router.Group("/company")
	{
		companys.POST("/", controllers.CreateCompany)
		companys.GET("/", controllers.ListCompany)
	}

	relations := router.Group("/relation")
	{
		relations.POST("/", controllers.CreateRelation)
		relations.GET("/", controllers.ListRelation)
		relations.DELETE("/:id", controllers.DeleteRelation)
	}

	messages := router.Group("/message")
	{
		messages.POST("/", controllers.CreateMessage)
		messages.GET("/", controllers.ListMessage)
	}

	router.Run(":8085")

}

// Database init func
func Database() *gorm.DB {
	myConfig := new(lib.Config)
	myConfig.InitConfig(lib.GetCurrentDir() + "/configs/configs.ini")
	host := myConfig.Read("database", "host")
	port := myConfig.Read("database", "port")
	user := myConfig.Read("database", "user")
	password := myConfig.Read("database", "password")
	dbname := myConfig.Read("database", "dbname")

	//open a db connection
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
