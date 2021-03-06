package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct { //考虑到用户可能通过postman这样的工具 直接创建有管理权限的账号 所以需要数据验证
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"` //用户名
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`  //密码
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限"`                    //权限1为管理员 2为普通用户 默认为2
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

// 更新查询  （供编辑用户使用 ）
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

//新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error //入参是接口 返回是db模型所以用错误处理来接受
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//查询用户列表 分页处理 要不然一次很多容易卡顿
func GetUsers(username string, pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int

	if username != "" {
		db.Select("id,username,role").Where(
			"username LIKE ?", username+"%",
		).Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)
		return users, total
	}
	db.Select("id,username,role").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)
	if err == gorm.ErrRecordNotFound {
		return users, 0
	}
	return users, total
}

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS

}

//编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{}) //先把Username等信息存到map中
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPw(password string) string {
	const KLen = 10                                                      //比特数
	salt := make([]byte, 8)                                              //相当于new一个 容量是8
	salt = []byte{21, 3, 43, 4, 22, 54, 53, 33}                          //加盐
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KLen) //成本参数16384 2^14(必须是2的幂) 这是官方推荐参数 可以随CPU性能增加 最后一位是长度
	if err != nil {
		log.Fatal(err) //放在日志中
	}

	fpassword := base64.StdEncoding.EncodeToString(HashPw) //拿到字符串形式的最后密码
	return fpassword
}

//登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 { //没有这个用户
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password { //密码错误
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 { //没有管理权限
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
