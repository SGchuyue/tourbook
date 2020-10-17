package main

import (
	"tourbook/model"
	"tourbook/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 调用路由
	routes.InitRouter()
}
