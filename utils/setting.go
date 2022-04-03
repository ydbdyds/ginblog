package utils

//该文件主要用来做 go-ini的处理 包括serve和database处理
import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() { //使用ini的库做一个处理
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") //读[serve]中的appmode 默认参数debug
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("3h45h3iASs")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File) { //对象存储配置
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
