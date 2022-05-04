package main

import (
	"ginblog/model"
	"ginblog/routes"
	"ginblog/task"
)

func main() {
	//引用数据库
	model.InitDb()
	task.ExecuteCron() //执行定时任务
	routes.InitRouter()

}
