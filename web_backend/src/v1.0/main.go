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
	router.RegisterController("Index", "/", "GET", &controller.IndexController{}, "Index")
	router.RegisterController("UserAdd", "/user", "POST", &controller.UserController{}, "Add")
	router.RegisterController("UserGet", "/user/:id", "GET", &controller.UserController{}, "Get")
	router.RegisterController("UserUpdate", "/user/:id", "PUT", &controller.UserController{}, "Update")
	router.RegisterController("Login", "/login", "POST", &controller.LoginController{}, "Index")
	router.RegisterController("Register", "/register", "POST", &controller.RegisterController{}, "Index")
	router.RegisterController("ArticleAdd", "/article", "POST", &controller.ArticleController{}, "Add")
	router.RegisterController("ArticleGet", "/article/:id", "GET", &controller.ArticleController{}, "Get")

	port := "8085"
	fmt.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router)) //启动server

}
