package service

import (
	"project/service/douseng"
	"project/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.SysGroup
	DouSengServiceGroup douseng.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
