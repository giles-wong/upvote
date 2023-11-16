package controllers

import (
	"net/http"

	"giles.wang/upvote/service"
	"github.com/giles-wong/general/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct{}

type UserInfoParam struct {
	Name     string `json:"name"`
	PassWord string `json:"passWord"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (u UserController) AddUser(ctx *gin.Context) {
	//接收请求参数
	param := UserInfoParam{}
	err := ctx.BindJSON(&param)
	zap.L().Info("接收用户添加参数", zap.Any("params", param))
	if err == nil {
		userServiceParam := &service.UserInfoParam{Name: param.Name, PassWord: param.PassWord, Phone: param.Phone, Email: param.Email}
		id := service.AddUser(userServiceParam)

		ctx.JSON(
			http.StatusOK,
			response.Success(id),
		)
	}

	ctx.JSON(
		http.StatusOK,
		response.Failure(2001, err),
	)
}

func (u UserController) GetUserInfo(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		response.Success("123456"),
	)
}

func (u UserController) GetUserList(c *gin.Context) {
	num1 := 1
	num2 := 0
	num3 := num1 / num2

	c.JSON(
		http.StatusOK,
		response.Success(num3),
	)
}
