package main

import (
	"github.com/IlhamSetiaji/go-crud/controllers"
	"github.com/IlhamSetiaji/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	posts := r.Group("/posts")
	{
		posts.POST("/", controllers.CreatePost)
		posts.GET("/", controllers.GetPosts)
		posts.GET("/find/:id", controllers.FindPostById)
		posts.PUT("/update/:id", controllers.UpdatePost)
		posts.DELETE("/delete/:id", controllers.DeletePost)
	}
	r.Run()
}
