package model

import (
	"context"
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` //关联 一对多
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
	View    uint64 `gorm:"type:int" json:"view"`
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
	article.Addview() //点击事件
	//cnt, _ := rdb.Get(context.Background(), "view:article:"+strconv.Itoa(int(article.ID))).Uint64()
	//if cnt > 0 {
	//	article.View += cnt
	//}
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

		for i, a := range articleList { //查看redis缓存是否有数据
			redisCnt := GetView(&a)
			//fmt.Println(a)
			if redisCnt > 0 {
				articleList[i].View += redisCnt
			}
		}
		//text
		//CheckAndUpdate()
		//text
		return articleList, errmsg.SUCCESS, total
	}
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Preload("Category").Where("title LIKE ?", title+"%").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}

	for i, a := range articleList { //查看redis缓存是否有数据
		redisCnt := GetView(&a)
		//fmt.Println(a)
		if redisCnt > 0 {
			articleList[i].View += redisCnt
		}
	}

	return articleList, errmsg.SUCCESS, total
}

//编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{}) //先把Username等信息存到map中
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["view"] = data.View
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

//文章点击 已测试
func (article *Article) Addview() {
	rdb.Incr(context.Background(), "view:article:"+strconv.Itoa(int(article.ID))) //拼接键值对
}

//获取文章点击量
func GetView(article *Article) uint64 {
	countStr, _ := rdb.Get(context.Background(), "view:article:"+strconv.Itoa(int(article.ID))).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	//fmt.Println("文章id:", article.ID, "点击:", count)
	return count
}

//redis缓存同步mysql
func CheckAndUpdate() {
	ctx := context.Background()
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, "*", 10).Result() //键 光标(到哪了) 10个每次
		if err != nil {
			logrus.Println("ScanError", err)
			break
		}

		for _, key := range keys {
			strValue, _ := rdb.Get(context.Background(), key).Result()
			value, _ := strconv.ParseUint(strValue, 10, 64)
			fmt.Println(value)
			if value == 0 {
				continue
			}
			temp := strings.Split(key, ":") //分割出id temp[2]
			//fmt.Printf("%v %v\n", temp[2], value)
			id, _ := strconv.Atoi(temp[2])
			article, _ := GetArticleInfo(id)
			//fmt.Println(article.View)
			article.View += value
			fmt.Println(article.View)
			EditArticle(id, &article)          //更新
			rdb.Del(context.Background(), key) //删除缓存
		}

		if cursor == 0 {
			break
		}

	}
}
