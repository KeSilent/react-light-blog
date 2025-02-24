package services

import "github.com/kesilent/react-light-blog/services/system"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}
