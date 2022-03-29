package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"` //用户名
	Password string `gorm:"type:varchar(20);not null" json:"password"` //密码
	Role     int    `gorm:"type:int" json:"role"`                      //权限
}

//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error //入参是接口 返回是db模型所以用错误处理来接受
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//查询用户列表 分页处理 要不然一次很多容易卡顿
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

////编辑用户
//func EditUser(id int)(code int)  {
//
//}
//
////删除用户