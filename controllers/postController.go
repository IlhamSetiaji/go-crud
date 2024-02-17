package controllers

import (
	"github.com/IlhamSetiaji/go-crud/initializers"
	"github.com/IlhamSetiaji/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var payload struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	c.Bind(&payload)

	post := models.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create post",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post has been created successfully",
		"data":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to get posts",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Posts has been fetched successfully",
		"data":    posts,
	})
}

func FindPostById(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to find post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post has been fetched successfully",
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to find post",
		})
		return
	}

	var payload struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	c.Bind(&payload)

	post.Title = payload.Title
	post.Body = payload.Body

	result = initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post has been updated successfully",
		"data":    post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Failed to find post",
		})
		return
	}

	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post has been deleted successfully",
	})
}
