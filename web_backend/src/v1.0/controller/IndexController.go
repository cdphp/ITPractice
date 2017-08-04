package controller

import (
	"fmt"

	"v1.0/vendor"
)

// IndexController struct
type IndexController struct {
	vendor.Controller
}

func (c *IndexController) Index() {
	sess := globalSessions.SessionStart(c.GetResponseWriter(), c.GetRequest())

	token := sess.Get("token")
	if token == nil {

	}

	fmt.Println(token)
}
