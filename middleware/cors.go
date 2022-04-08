package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//跨域配置
func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowAllOrigins: true, //跨域域名 允许所有的跨域
		//AllowOrigins: []string{"*"},这个与上面的等效
		AllowMethods:  []string{"*"}, //允许的请求方法 全部允许
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length", "Authorization"},
		//AllowCredentials: true, //是否发送cookie请求

		MaxAge: 12 * time.Hour,
	})

}
