// 分类接口
package v1

import "github.com/gin-gonic/gin" // 引用上下文
import "tourbook/utils/errmsg"
import (
        "tourbook/model"
        "net/http"
        "strconv"
)

// 查询分类是否存在

// 添加分类

//var code int
// 添加分类
func AddCategory(c *gin.Context) {
        // todo 添加用户
        var data model.Category
        _ = c.ShouldBindJSON(&data) //
        code = model.CheckCategory(data.Name)
        if code == errmsg.SUCCSE {
                model.CreateCate(&data)
        }
        if code == errmsg.ERROR_CATENAME_USED {
                code = errmsg.ERROR_CATENAME_USED
        }

        c.JSON(http.StatusOK,gin.H{
                "status":code,
                "data":data,
                "message":errmsg.GetErrMsg(code),
        })
}

// 查询单个用户

// 查询分类列表
func GetCates(c *gin.Context) {
        // 分页功能 strconv.atoi 转换格式
        pageSize,_ := strconv.Atoi(c.Query("pagesize"))
        pageNum,_ := strconv.Atoi(c.Query("pagenum"))

        if pageSize == 0 {
                pageSize = -1
        }
        if pageNum == 0 {
                pageNum = -1
        }

        data := model.GetCates(pageSize,pageNum)
        code = errmsg.SUCCSE
        c.JSON(http.StatusOK,gin.H{
                "status":code,
                "data":data,
                "message":errmsg.GetErrMsg(code),
        })
}

// 编辑分类
func EditCate(c *gin.Context) {
        var data model.Category
        id,_ := strconv.Atoi(c.Param("id"))
        c.ShouldBindJSON(&data)
        code = model.CheckCategory(data.Name)
        if code == errmsg.SUCCSE {
                model.EditCate(id,&data)
        }
        if code == errmsg.ERROR_CATENAME_USED{
                c.Abort()
	}
//	data := model.GetCates(pageSize,pageNum)
//        code = errmsg.SUCCSE
        c.JSON(http.StatusOK,gin.H{
                "status":code,
//                "data":data,
                "message":errmsg.GetErrMsg(code),
        })
}

/*
// 删除分类
func DeleteCate(c *gin.Context){
        id,_ := strconv.Atoi(c.Param("id"))

        code = model.DeleteCate(id)

        c.JSON(http.StatusOK,gin.H{
                "status": code,
                "message": errmsg.GetErrMsg(code),
        })

}*/
// 删除分类
func DeleteCate(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
