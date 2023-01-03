package main

import (
	"GinBlog/model"
	"GinBlog/routes"

	"fmt"
)

func main() {
	fmt.Println("hello Golang!!!")
	// r := gin.Default()
	// r.LoadHTMLFiles("./index.tmpl")
	// r.GET("/susu", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", nil)
	// })
	// r.Run(":8888")

	model.InitDb()
	routes.InitRouter()
}
