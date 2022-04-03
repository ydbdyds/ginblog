package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/blog.log" //日志地址
	linkName := "latest_log.log"

	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755) //控制权限
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = src //设置日志输出地址

	logger.SetLevel(logrus.DebugLevel) // 设置日志等级

	logWrite, _ := rotatelogs.New(
		filePath+"%Y%m%d.log",                     //分隔日志  设置保存日志的格式
		rotatelogs.WithMaxAge(7*24*time.Hour),     //最大保存时间 一周
		rotatelogs.WithRotationTime(24*time.Hour), //多久分隔一次
		rotatelogs.WithLinkName(linkName),         //建立软链接
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWrite,
		logrus.FatalLevel: logWrite,
		logrus.DebugLevel: logWrite,
		logrus.WarnLevel:  logWrite,
		logrus.ErrorLevel: logWrite,
		logrus.PanicLevel: logWrite,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //golang的time规定
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() //gin的中间件模型是洋葱模型 看到c.next就执行下一个中间件 很像递归 next下面的最后执行
		stoptime := time.Since(startTime)
		spendtime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stoptime.Nanoseconds())/1000000.0))) //花费时间
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    //状态码
		clientIp := c.ClientIP()           //客户端请求过来的ip
		userAgent := c.Request.UserAgent() //客户端的浏览器信息
		dataSize := c.Writer.Size()        //文件长度

		if dataSize < 0 {
			dataSize = 0
		}

		method := c.Request.Method   //请求方法
		path := c.Request.RequestURI //请求路径

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendtime,
			"IP":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 { //系统内部存在错误
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 { //错误码大于500
			entry.Error()
		} else if statusCode >= 400 { //警告
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
