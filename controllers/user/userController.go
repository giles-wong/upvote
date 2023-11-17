package user

import (
	"giles.wang/upvote/common/system"
	userService "giles.wang/upvote/service/user"
	"github.com/giles-wong/general/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
}

type RegisterParams struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
}

// Register 用户注册
func (user UserController) Register(c *gin.Context) {
	//接收用户参数
	params := RegisterParams{}
	err := c.BindJSON(&params)
	zap.L().Info("用户注册接收参数", zap.Any("params", params))
	if err == nil {
		//参数验证
		if params.Name == "" || params.Password == "" || params.ConfirmPassword == "" || params.Phone == "" {
			response.GinFailure(c, system.RequestParamsNull, system.RequestParamsNullMsg)
			return
		}
		//传递参数
		serviceParam := userService.RegisterParam{
			Name:     params.Name,
			Password: params.Password,
			Phone:    params.Phone,
			Email:    params.Email,
		}

		user := userService.Register(&serviceParam)
		zap.L().Info("用户注册接收参数222", zap.Any("params", params))
		response.GinSuccess(c, user)
		return
	}
	response.GinFailure(c, system.RequestParamsNull, system.RequestParamsNullMsg)
}
