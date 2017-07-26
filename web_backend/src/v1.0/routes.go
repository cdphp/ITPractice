package main

import "net/http"

// Route struct
type Route struct {
	Name        string           //名称
	Method      string           //http方法(GET,POST..)
	Pattern     string           //uri
	HandlerFunc http.HandlerFunc //处理方法
}

// Routes res
type Routes []Route

// routes 自定义路由
var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"ArticleIndex", "GET", "/article", ArticleIndex},
}
