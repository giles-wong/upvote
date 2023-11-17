package user

import (
	"fmt"
	"giles.wang/upvote/models"
	"github.com/giles-wong/general/encrypt"
	"github.com/giles-wong/general/snowflake"
	"go.uber.org/zap"
)

type RegisterParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func Register(param *RegisterParam) int64 {
	//绑定参数
	user := &models.User{Name: param.Name, Phone: param.Phone, Email: param.Email}
	//密码 md5加密
	user.Password = encrypt.Md5(param.Password)
	//生成雪花算法Id
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	user.Id = node.GetId()
	zap.L().Info("用户添加最终生成的结构体", zap.Any("user", user))
	return models.AddUser(user)
}
