package routes

import (
	"tourbook/utils"
	"github.com/gin-gonic/gin"
//	"net/http"
	"tourbook/api/v1"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
	// 用于测试
	/*	router.GET("hello",func(c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"msg":"ok",
			})
		})
	*/
		// user 模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user/:id",v1.DeleteUser)

		// 分类模块的路由接口
		router.POST("category/add",v1.AddCategory)
                router.GET("category",v1.GetCates)
                router.PUT("category/:id",v1.EditCate)
                router.DELETE("category/:id",v1.DeleteCate)
		// 文章模块的路由接口
	}

	r.Run(utils.HttpPort)
}
