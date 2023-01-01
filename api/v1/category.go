// 相关category模块的业务操作 Controler
package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBind(&data)
	code := model.CheckCategory(data.Name)
	// 分类名未使用
	if code == errmsg.SUCCESS {
		code2 := model.CreateCategory(&data)
		var message string
		if code2 == errmsg.ERROR { // 创建分类错误
			message = errmsg.GetErrMsg(code2)
		} else { // 创建成功
			message = "创建分类成功..."
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code2,
			"message": message,
		})
	}
	// 分类名已存在
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 查询分类列表
func GetCategory(c *gin.Context) {
	// 从前端传来的Query中 获取pagesize和pagenum字段
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // -1表示取消 limit
	}
	if pageNum == 0 {
		pageNum = -1 // -1表示取消 offset
	}
	data := model.GetCategory(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	c.ShouldBindJSON(&data)
	code, cate := model.GetCateInfo(id) // 根据id获取分类的所有信息
	var message string
	if code == errmsg.SUCCESS {
		// 1. 分类名没修改 ==> 可以直接编辑分类信息
		if data.Name == cate.Name {
			model.EditCategory(id, &data)
			message = errmsg.GetErrMsg(code)
		} else { // 2. 分类名修改了 ==> 考虑修改后的名字是否重复
			code2 := model.CheckCategory(data.Name)
			if code2 == errmsg.SUCCESS { // 不重复
				model.EditCategory(id, &data)
			}
			message = errmsg.GetErrMsg(code2)
		}
	} else {
		message = errmsg.GetErrMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": message,
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
