package routes

import (
	"GinBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	rp := r.Group("api/v1")
	{
		rp.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		})
	}
	r.Run(utils.HttpPort)
}
