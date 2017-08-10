package main

import (
	"fmt"
	"log"
	"net/http"

	"v1.0/controller"
	"v1.0/vendor"
)

func main() {

	router := vendor.NewRouter()

	router.RegisterController("Index", "/index", "GET", &controller.IndexController{}, "Index")

	router.RegisterController("UserList", "/user", "GET", &controller.UserController{}, "Index")
	router.RegisterController("UserAdd", "/user", "POST", &controller.UserController{}, "Add")
	router.RegisterController("UserGet", "/user/:id", "GET", &controller.UserController{}, "Get")
	router.RegisterController("UserUpdate", "/user/:id", "PUT", &controller.UserController{}, "Update")
	router.RegisterController("Login", "/login", "POST", &controller.LoginController{}, "Index")
	router.RegisterController("Register", "/register", "POST", &controller.RegisterController{}, "Index")

	router.RegisterController("ArticleList", "/article", "GET", &controller.ArticleController{}, "Index")
	router.RegisterController("ArticleAdd", "/article", "POST", &controller.ArticleController{}, "Add")
	router.RegisterController("ArticleGet", "/article/:id", "GET", &controller.ArticleController{}, "Get")
	router.RegisterController("ArticleUpdate", "/article/:id", "PUT", &controller.ArticleController{}, "Update")
	router.RegisterController("ArticleDelete", "/article/:id", "DELETE", &controller.ArticleController{}, "Delete")

	router.RegisterController("CommentList", "/comment", "GET", &controller.CommentController{}, "Index")
	router.RegisterController("CommentAdd", "/comment", "POST", &controller.CommentController{}, "Add")
	port := "8085"
	fmt.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router)) //启动server

}
