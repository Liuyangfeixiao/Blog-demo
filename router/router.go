package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/", controller.Index)
	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)

	// blog operation
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)

	e.GET("/postDetails", controller.PostDetails)

	e.Run(":8888")
}
