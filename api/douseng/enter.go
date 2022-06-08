package douseng

import "project/service"

type ApiGroup struct {
	DouSengPJHApi
}

var (
	douSengPJHService = service.ServiceGroupApp.DouSengServiceGroup.DouSengPJHService
)