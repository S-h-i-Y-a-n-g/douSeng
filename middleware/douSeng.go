package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/global"
	res "project/model/douseng/response"
	"project/utils"
)



func ValidateUser(token string,userID int64,c *gin.Context) error {
	//解析token
	j := &JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名

	userinfo, err := j.ParseTokenDouSeng(token)
	if err != nil {
		global.GSD_LOG.Error("token 解析失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DouSengUser{
			DSResponse: res.DSResponse{
				StatusMsg:  "token信息错误",
				StatusCode: 1,
			},
		},
		)
	}
	//验证信息同步
	if int64(userinfo.ID) != userID {
		c.JSON(http.StatusOK, res.DouSengUser{
			DSResponse: res.DSResponse{
				StatusMsg:  "token信息错误",
				StatusCode: 1,
			},
		},
		)
		return errors.New("token信息错误")
	}

	return err
}
