package douseng

import "project/service"

type ApiGroup struct {
	DouSengPJHApi
	DouSengLYFApi
	DouSengXYFApi
	DouShengWYHApi
}

var (
	douSengPJHService = service.ServiceGroupApp.DouSengServiceGroup.DouSengPJHService
	douSengLYFService = service.ServiceGroupApp.DouSengServiceGroup.DouSengLYFService
	douShengWYHService = service.ServiceGroupApp.DouSengServiceGroup.DouShengWYHService
	jwtService        = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
