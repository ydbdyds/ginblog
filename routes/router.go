package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() //路由初始化

	router := r.Group("api/v1") //在这个组内
	{
		//用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArticle)
		router.PUT("article/:id", v1.EditArticle)
		router.DELETE("article/:id", v1.DeleteArticle)

	}

	r.Run(utils.HttpPort) //跑在这个端口
}
