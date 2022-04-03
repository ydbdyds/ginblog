package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //路由初始化

	auth := r.Group("api/v1") //在这个组内需要jwt中间件
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle) //删除文章
		auth.POST("upload", v1.Upload)               //上传文件
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCategory)
		router.GET("articles", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCategoryArt) //分类下的文章
		router.GET("article/info/:id", v1.GetArticleInfo) //查询文章详情
		router.POST("login", v1.Login)                    //登陆接口
	}

	r.Run(utils.HttpPort) //跑在这个端口
}
