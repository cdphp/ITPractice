package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := NewRouter() //加载路由
	port := "8085"
	fmt.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router)) //启动server

}
