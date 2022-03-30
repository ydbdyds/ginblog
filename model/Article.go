package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category //关联
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error //入参是接口 返回是db模型所以用错误处理来接受
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//todo 查询分类文章
//todo 查询文章详情

//todo 查询文章列表 分页处理 要不然一次很多容易卡顿
func GetArticle(pageSize int, pageNum int) []Article {
	var article []Article
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return article
}

//编辑分类
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{}) //先把Username等信息存到map中
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
