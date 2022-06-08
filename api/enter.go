package v1

import (
	"project/api/douseng"
	"project/api/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	DouSengApiGroup douseng.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
