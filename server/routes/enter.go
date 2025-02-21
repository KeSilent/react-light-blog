package router

import "github.com/kesilent/react-light-blog/routes/system"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
}
