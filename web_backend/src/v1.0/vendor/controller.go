package vendor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ControllerInterface interface
type ControllerInterface interface {
	Init(ct *Context, controllerName string)
	Get()
}

// Context struct
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         map[int]string
}

// Controller struct
type Controller struct {
	ct   *Context
	name string
}

// Init set context and name
func (c *Controller) Init(ct *Context, controllerName string) {
	c.ct = ct
	c.name = controllerName
	c.GetResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
	c.GetResponseWriter().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	//c.Initialize()
}

// Initialize run before action
func (c *Controller) Initialize() {
	fmt.Println("Initialize:", c.name)
}

// Get func
func (c *Controller) Get() {
	fmt.Println("hello")
}

// GetControllerName func
func (c *Controller) GetControllerName() string {
	return c.name
}

// GetContext func
func (c *Controller) GetContext() *Context {
	return c.ct
}

// GetParams func
func (c *Controller) GetParams() map[int]string {
	return c.ct.Params
}

// GetResponseWriter func
func (c *Controller) GetResponseWriter() http.ResponseWriter {
	return c.ct.ResponseWriter
}

// GetRequest func
func (c *Controller) GetRequest() *http.Request {
	return c.ct.Request
}

// GetQuery 获取get参数
func (c *Controller) GetQuery(param string) string {
	query := c.ct.Request.URL.Query()

	return query.Get(param)
}

// GetPosts 获取post数据
func (c *Controller) GetPosts() map[string]interface{} {
	request := c.GetRequest()

	request.ParseForm()
	fmt.Println(request)

	postData, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()

	result := make(map[string]interface{})
	if err := json.Unmarshal([]byte(postData), &result); err != nil {
		panic(err)
	}

	return result
}
