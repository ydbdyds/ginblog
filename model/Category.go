package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//分类用不到软删除 尽量简化结构体
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"` //分类
}

//查询用分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error //入参是接口 返回是db模型所以用错误处理来接受
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//查询分类列表 分页处理 要不然一次很多容易卡顿
func GetCategory(pageSize int, pageNum int) ([]Category, int) {
	var category []Category
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&category).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return category, total
}

//编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{}) //先把Username等信息存到map中
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的所有文章

//删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
