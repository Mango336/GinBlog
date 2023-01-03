package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	fmt.Println(data)
	code := model.CheckLogin(data.Username, data.Password)
	var tokenStr string
	// var setTokenCode int = errmsg.SUCCESS
	if code == errmsg.SUCCESS {
		code, tokenStr = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   tokenStr,
	})
}
