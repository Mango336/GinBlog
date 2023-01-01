// 相关article模块的业务操作 Controler
package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBind(&data)
	// fmt.Println("=====Add Article:===\n", data)
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的文章
func GetArtInCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // -1表示取消 limit
	}
	if pageNum == 0 {
		pageNum = -1 // -1表示取消 offset
	}
	cid, _ := strconv.Atoi(c.Param("cid")) // 从参数中获取分类id
	code, data := model.GetArtInCate(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // 从参数中获取文章id
	code, data := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询文章列表
func GetArtList(c *gin.Context) {
	// 从前端传来的Query中 获取pagesize和pagenum字段
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 // -1表示取消 limit
	}
	if pageNum == 0 {
		pageNum = -1 // -1表示取消 offset
	}
	code, data := model.GetArtList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	c.ShouldBindJSON(&data)
	code := model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
