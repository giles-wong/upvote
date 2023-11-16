package main

import (
	"giles.wang/upvote/config"
	"giles.wang/upvote/pkg/logger"
	"giles.wang/upvote/router"
)

func main() {
	//加载配置文件
	if err := config.Init("./config/config.json"); err != nil {
		panic(err)
	}
	// 初始化日志配置
	logger.InitLogger(config.Conf.LogConfig)

	r := router.Router()

	r.Run()
}
