package vendor

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

// routeInfo struct
type routeInfo struct {
	name           string         //名称
	pattern        string         //路径
	method         string         //get,post,put...
	params         map[int]string //参数
	controllerType reflect.Type   //控制器类型
	action         string         //响应
}

var routePath = []routeInfo{}

// NewRouter 返回一个Router实例
func NewRouter() *Router {
	return new(Router)
}

// Router struct
type Router struct {
	Route map[string]routeInfo
}

// 实现Handler接口
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isFound := false

	for _, p := range router.Route {

		reg, err := regexp.Compile(p.pattern)
		if err != nil {
			continue
		}

		if reg.MatchString(r.URL.String()) && p.method == r.Method {

			matches := reg.FindStringSubmatch(r.URL.String())

			if len(matches[0]) != len(r.URL.Path) {
				continue
			}
			params := make(map[int]string)
			fmt.Println("length of params:", len(p.params))
			if len(p.params) > 0 {
				//add url parameters to the query param map
				values := r.URL.Query()

				k := 1
				for i, match := range matches[1:] {
					values.Add(p.params[i], match)
					params[k] = match
					k++
				}

			}
			fmt.Println("params:", params)
			vc := reflect.New(p.controllerType)
			//fmt.Println("vc:", vc)
			//fmt.Println("type:", p.controllerType)
			init := vc.MethodByName("Init")
			//fmt.Println("init:", init)
			in := make([]reflect.Value, 2)
			ct := &Context{ResponseWriter: w, Request: r, Params: params}

			in[0] = reflect.ValueOf(ct)
			in[1] = reflect.ValueOf(p.controllerType.Name())
			//fmt.Println("in", in)

			init.Call(in)

			method := vc.MethodByName(p.action)
			in = make([]reflect.Value, 0)
			method.Call(in)

			isFound = true

		}
	}
	if !isFound {
		fmt.Fprintf(w, "404 Page not found")
	}

}

// RegisterController func
func (router *Router) RegisterController(routeName string, path string, method string, c ControllerInterface, action string) {
	if router.Route == nil {
		router.Route = make(map[string]routeInfo)
	}

	parts := strings.Split(path, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			//a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	path = strings.Join(parts, "/")
	_, regexErr := regexp.Compile(path)
	if regexErr != nil {

		//TODO add error handling here to avoid panic
		panic(regexErr)
	}
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	fmt.Println(path, params, t)

	router.Route[routeName] = routeInfo{routeName, path, method, params, t, action}
}
