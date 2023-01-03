package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"GinBlog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	// 需要token验证的分组
	rpToken := r.Group("api/v1")
	rpToken.Use(middleware.JwtTokenHandler())
	{
		// 用户模块
		rpToken.PUT("user/:id", v1.EditUser)      // 编辑用户信息
		rpToken.DELETE("user/:id", v1.DeleteUser) // 删除用户
		// 分类模块
		rpToken.POST("cate/add", v1.AddCategory)      // 添加分类
		rpToken.PUT("cate/:id", v1.EditCategory)      // 编辑分类信息
		rpToken.DELETE("cate/:id", v1.DeleteCategory) // 删除分类
		// 文章模块
		rpToken.POST("art/add", v1.AddArticle)      // 添加文章
		rpToken.PUT("art/:id", v1.EditArticle)      // 编辑文章信息
		rpToken.DELETE("art/:id", v1.DeleteArticle) // 删除文章
	}

	rp := r.Group("api/v1") // GET方法 不需要token验证
	{
		rp.POST("login", v1.Login)            // 登录接口
		rp.POST("user/add", v1.AddUser)       // 添加用户
		rp.GET("users", v1.GetUsers)          // 获取分类列表
		rp.GET("cates", v1.GetCategory)       // 获取分类列表
		rp.GET("arts", v1.GetArtList)         // 获取文章列表
		rp.GET("art/info/:id", v1.GetArtInfo) // 获取单个文章信息
		rp.GET("art/:cid", v1.GetArtInCate)   // 获取某个分类的文章
	}
	r.Run(utils.HttpPort)
}
