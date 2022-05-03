package model

import (
	"context"
	"fmt"
	"ginblog/utils"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error
var rdb *redis.Client

//处理连接数据库

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.Host, utils.Port),
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping(context.TODO()).Result()

	if err != nil {
		fmt.Printf("连接数据库失败,检查参数!", err)
	}

	db.SingularTable(true)                                                   //单数模式
	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{}, &Comment{}) //数据库迁移  自动根据model建表 但是他会自动是复数

	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置数据库连接最大打开数。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置可重用连接的最长时间
	db.DB().SetConnMaxLifetime(10 * time.Second) //10秒钟 不能大于gin框架的timeout时间 r.run有一个默认时间

	//db.Close()
}
