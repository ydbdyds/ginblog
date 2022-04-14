package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New() //路由初始化
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	//托管前端
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static") //路由路径 文件路径
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

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
		//个人设置
		auth.PUT("profile/:id", v1.UpdateProfile)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo) //查询单个用户
		router.GET("category", v1.GetCategory)
		router.GET("articles", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCategoryArt) //分类下的文章
		router.GET("article/info/:id", v1.GetArticleInfo) //查询文章详情
		router.POST("login", v1.Login)                    //登陆接口
		router.GET("category/:id", v1.GetCateInfo)        //查询单个分类
		//个人信息
		router.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort) //跑在这个端口
}
