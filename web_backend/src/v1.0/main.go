package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter() //加载路由

	log.Fatal(http.ListenAndServe(":8085", router)) //启动server
}
