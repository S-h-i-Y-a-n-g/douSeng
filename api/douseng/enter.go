package douseng

import "project/service"

type ApiGroup struct {
	DouSengPJHApi
	DouSengLYFApi
}

var (
	douSengPJHService = service.ServiceGroupApp.DouSengServiceGroup.DouSengPJHService
	douSengLYFService = service.ServiceGroupApp.DouSengServiceGroup.DouSengLYFService
	jwtService        = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
