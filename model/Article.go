package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //关联 一对多
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

//查询分类文章 根据分类id 返回切片
func GetCategoryArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	db.Preload("Category").Where("cid =?", id).Find(&articleList).Count(&total)
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Cid = ?", id).Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIT, 0
	}
	return articleList, errmsg.SUCCESS, total
}

//查询文章详情
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST //返回空结构体并且报错
	}
	return article, errmsg.SUCCESS
}

//查询文章列表 分页处理 要不然一次很多容易卡顿 通过预加载来加载关联 查找article时就预加载相关分类  通过order实现更新时间排序
func GetArticle(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	if title == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Preload("Category").Find(&articleList).Error
		// 单独计数
		db.Model(&articleList).Count(&total)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
		return articleList, errmsg.SUCCESS, total
	}
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?", title+"%").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
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
