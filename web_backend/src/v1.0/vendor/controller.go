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

// Init func
func (c *Controller) Init(ct *Context, controllerName string) {
	c.ct = ct
	c.name = controllerName
	//c.Initialize()
}

// Initialize func
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

// GetPosts 获取post数据
func (c *Controller) GetPosts() map[string]string {
	request := c.GetRequest()
	request.ParseForm()

	postData, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()

	result := make(map[string]string)
	if err := json.Unmarshal([]byte(postData), &result); err != nil {
		panic(err)
	}

	return result
}
