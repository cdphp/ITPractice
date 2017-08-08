package controller

import (
	"fmt"

	"v1.0/vendor"
)

// StaticController struct
type StaticController struct {
	vendor.Controller
}

// Initialize 初始化
func (c *StaticController) Initialize() {
	fmt.Println("Static controller intialize ")

}

// Index 首页
func (c *StaticController) Index() {

}
