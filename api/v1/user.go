// 相关user模块的业务操作 Controler
package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)
	code := model.CheckUser(data.Username)
	// 用户名未使用
	if code == errmsg.SUCCESS {
		code2 := model.CreateUser(&data)
		var message string
		if code2 == errmsg.ERROR { // 创建用户错误
			message = errmsg.GetErrMsg(code2)
		} else { // 创建成功
			message = "创建用户成功..."
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code2,
			"message": message,
		})
	}
	// 用户名已存在
	if code == errmsg.ERROR_USERNAME_USED {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	// 从前端传来的Query中 获取pagesize和pagenum字段
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	// fmt.Println(pageSize, pageNum)
	if pageSize == 0 {
		pageSize = -1 // -1表示取消 limit
	}
	if pageNum == 0 {
		pageNum = -1 // -1表示取消 offset
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	c.ShouldBindJSON(&data)
	// fmt.Println(data)
	code, user := model.GetUserInfo(id) // 根据id获取用户的所有信息
	var message string
	if code == errmsg.SUCCESS {
		// 1. 用户名没修改 ==> 可以直接编辑用户信息
		if data.Username == user.Username {
			model.EditUser(id, &data)
			message = errmsg.GetErrMsg(code)
		} else { // 2. 用户名修改了 ==> 考虑修改后的名字是否重复
			code2 := model.CheckUser(data.Username)
			if code2 == errmsg.SUCCESS { // 不重复
				model.EditUser(id, &data)
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

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
