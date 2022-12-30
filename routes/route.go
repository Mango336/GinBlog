package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	rp := r.Group("api/v1")
	{
		// rp.GET("hello", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		// })

		// rp.GET("user/:id", v1.UserExist)

		// 用户模块
		rp.POST("user/add", v1.AddUser)
		rp.GET("users", v1.GetUsers)
		rp.PUT("user/:id", v1.EditUser)
		rp.DELETE("user/:id", v1.DeleteUser)
		// 分类模块
		rp.POST("cate/add", v1.AddCategory)
		rp.GET("cates", v1.GetCategory)
		rp.PUT("cate/:id", v1.EditCategory)
		rp.DELETE("cate/:id", v1.DeleteCategory)
	}
	r.Run(utils.HttpPort)
}
