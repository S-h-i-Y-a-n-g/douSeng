package douseng

import (
	"fmt"
	"github.com/gin-gonic/gin"
	request "project/model/douseng/response"
	v1 "project/service/douseng"
)

type DouSengXYFApi struct{}

var v v1.DouSengXYFService

func (d *DouSengXYFApi) Action(c *gin.Context) {
	fmt.Println("++++++++++favitor_测试+++++++++")
	v.FavoriteService(request.Video{}, request.User{})
}
