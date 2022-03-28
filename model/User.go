package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"` //用户名
	Password string `gorm:"type:varchar(20);not null" json:"password"` //密码
	Role     int    `gorm:"type:int" json:"role"`                      //权限
}
