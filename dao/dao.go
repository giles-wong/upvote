package dao

import (
	"giles.wang/upvote/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MySqlDb), &gorm.Config{})
	if err != nil {
		zap.L().Error("数据库链接失败", zap.Any("error", err))
	}

	if Db.Error != nil {
		zap.L().Error("数据库初始化失败", zap.Any("error", Db.Error))
	}
	Db.DB()
}
