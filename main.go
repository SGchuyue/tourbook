package main

import "tourbook/routes"
import "tourbook/model"

func main() {

	// 引用数据库
	model.InitDb()

	routes.InitRouter()
}
