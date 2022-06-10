package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/service/douseng"
)

type DouSengXYFApi struct{}

var v v1.DouSengXYFService

func (d *DouSengXYFApi) Action(c *gin.Context) {
	//v.FavoriteService()
}
