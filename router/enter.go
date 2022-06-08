package router

import (
	"project/router/douseng"
	"project/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	DouSeng douseng.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
