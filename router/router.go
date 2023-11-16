package router

import (
	"giles.wang/upvote/config"
	"giles.wang/upvote/controllers"
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
		user.GET("/list", controllers.UserController{}.GetUserList)
		user.POST("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/add", controllers.UserController{}.AddUser)
	}

	return r
}
