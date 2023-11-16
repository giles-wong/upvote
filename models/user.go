package models

import (
	"giles.wang/upvote/dao"
)

type User struct {
	Model
	Id       int64  `gorm:"column:id;" json:"id"`                               //用户Id  唯一标识
	Name     string `gorm:"column:name;type:varchar(64);" json:"name"`          //用户姓名
	Password string `gorm:"column:password;type:varchar(128);" json:"password"` //用户密码
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`        //用户手机号
	Email    string `gorm:"column:email;type:varchar(128);" json:"email"`       //用户邮箱
}

func (User) TableName() string {
	return "gin_test_user"
}

func GetUserInfo(id int64) (User, error) {
	var user User

	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(user *User) int64 {

	err := dao.Db.Create(user)
	if err != nil {
		return user.Id
	}
	return 0
}
