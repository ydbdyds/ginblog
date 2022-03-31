package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey) //秘钥
//jwt不存放密码 敏感信息不建议存放
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var code int

//jwt分为三部分 头部
//生成认证token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 3) //3个小时有效时间
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime), //过期时间
			Issuer:    "ydbdyds",          //签发人
		},
	}
	//newwithclaims将载荷和头部做连接 返回一个还没加签的token
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) //入参是签发方法 返回一个token
	token, err := reqClaim.SignedString(JwtKey)                      //加签
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token 把传入的token和token做一个对比
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil //用于验证的秘钥
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid { //格式化+解码  用到了Go的类型断言
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}

}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization") //拿到请求头

		if tokenHeader == "" { //如果请求为空
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)      //分割请求头
		if len(checkToken) != 2 && checkToken[0] != "Bearer" { //Bearer约定后面跟着token
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		key, tCheck := CheckToken(checkToken[1]) //把分割出来的拿去比对
		if tCheck == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME //判断过期
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()

	}
}
