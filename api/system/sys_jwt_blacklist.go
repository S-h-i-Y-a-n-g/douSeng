package system

import (
	"project/global"
	"project/model/common/response"
	"project/model/system"
	"project/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct {
}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /api/jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JoinInBlacklist(jwt); err != nil {
		global.GSD_LOG.Error("jwt作废失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
