package router

import (
	"giles.wang/upvote/config"
	userControllers "giles.wang/upvote/controllers/user"
	"giles.wang/upvote/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(config.Conf.Mode)
	r := gin.New()
	//增加日志以及异常捕获
	r.Use(logger.WithConfig(), logger.Recovery())
	//用户组
	user := r.Group("/user")
	{
		user.POST("/register", userControllers.UserController{}.Register)
	}

	return r
}
