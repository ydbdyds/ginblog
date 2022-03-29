package utils

//该文件主要用来做 go-ini的处理 包括serve和database处理
import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() { //使用ini的库做一个处理
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:", err)
	}
	LoadServer(file)
	LoadDara(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") //读[serve]中的appmode 默认参数debug
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadDara(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}