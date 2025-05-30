package services

import (
	initializer "github.com/kesilent/react-light-blog/services/Initializer"
	"github.com/kesilent/react-light-blog/services/autocode"
	"github.com/kesilent/react-light-blog/services/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	InitializerGroup   initializer.InitializerGroup
	AutoCodeGroup      autocode.AutoCodeGroup
}
