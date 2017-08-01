package config

import "net/http"
import "v1.0/controller"

// Route struct
type Route struct {
	Name        string           //名称
	Method      string           //http方法(GET,POST..)
	Pattern     string           //uri
	HandlerFunc http.HandlerFunc //处理方法
}

// Routes res
type Routes []Route

// GetRoutes function
func GetRoutes() []Route {
	// routes 自定义路由
	routes := Routes{
		Route{"Index", "GET", "/", &controller.UserController{}},
		//Route{"ArticleIndex", "GET", "/article", ArticleIndex},
		//Route{"GetUser", "GET", "/user/{id:[0-9]+}", GetUser},
		//Route{"Register", "POST", "/user", AddUser},
		//Route{"Login", "POST", "/login", Login},
	}

	return routes
}
