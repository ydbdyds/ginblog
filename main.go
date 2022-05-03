package main

import (
	"ginblog/model"
	"ginblog/routes"
	"ginblog/task"
)

func main() {
	//引用数据库
	model.InitDb()
	task.ExecuteCron()
	routes.InitRouter()

}
