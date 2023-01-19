package v1

import (
	"GinBlog/server"
	"GinBlog/utils/errmsg"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	fileData, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal("Upload file is error, err: ", err.Error())
		return
	}
	code, url := server.UploadFile(fileData, fileHeader.Size) // (文件内容访问接口, 文件大小)
	fmt.Println(code, url)
	if code == errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
