package main

import (
	"tourbook/routes"
	"tourbook/model"
)

func main() {

	// 引用数据库
	model.InitDb()
	// 调用路由
	routes.InitRouter()
}
