package routes

import (
	"tourbook/utils"
	"github.com/gin-gonic/gin"
//	"net/http"
	"tourbook/api/v1"
	"tourbook/middleware"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Logger())
	auth := r.Group("api/v1")
		auth.Use(middleware.JwtToken())
	{
	// 用于测试
	/*	router.GET("hello",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"msg":"ok",
			})
		})
	*/
		// user 模块的路由接口
	//	router.POST("user/add",v1.AddUser)
	//	auth.GET("users",v1.GetUsers)
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)

		// 分类模块的路由接口
		auth.POST("category/add",v1.AddCategory)
              //  router.GET("category",v1.GetCates)
                auth.PUT("category/:id",v1.EditCate)
                auth.DELETE("category/:id",v1.DeleteCate)
		// 文章模块的路由接口
		//auth.POST("article/add",v1.AddArticle)
		//auth.get("article",v1.GetArt)
		auth.PUT("article/:id",v1.EditArt)
		auth.DELETE("article/:id",v1.DeleteArt)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add",v1.AddUser)
		router.GET("category",v1.GetCates)
		router.POST("login",v1.Login)
		router.GET("users",v1.GetUsers)
		router.GET("article",v1.GetArt)
		router.GET("article/list",v1.GetCateArt)
		router.GET("article/info/:id",v1.GetCateArt)
		router.POST("article/add",v1.AddArticle)
	}

	r.Run(utils.HttpPort)
}
