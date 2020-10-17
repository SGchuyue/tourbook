package routes

import (
	"github.com/gin-gonic/gin"
	"tourbook/api/v1"
	"tourbook/middleware"
	"tourbook/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Logger())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// user 模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("category", v1.GetCates)
		router.POST("login", v1.Login)
		router.GET("users", v1.GetUsers)
		router.GET("article", v1.GetArt)
		router.GET("article/list", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetCateArt)
		router.POST("article/add", v1.AddArticle)
	}

	r.Run(utils.HttpPort)
}
