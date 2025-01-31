package controllers

import (
	"github.com/Abhishek-M-K/go-monitoring/initializers"
	"github.com/Abhishek-M-K/go-monitoring/models"
	"github.com/gin-gonic/gin"
)

func Posts (c *gin.Context) {
	// Get data from the request body
	var body struct{
		Title string 
		Body string 
	}

	c.Bind(&body)

	// Create
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	
	// Return the response
  	c.JSON(201, gin.H{
		"data": post,
  		"message": "Post created successfully!",
  	})
}

func GetPosts (c *gin.Context) {

	// Get all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return the response
	c.JSON(200, gin.H{
		"data": posts,
		"message": "Posts retrieved successfully!",
	})
}

func GetPostByIndex (c *gin.Context) {
	// Get index by URL parameter
	id := c.Param("id")

	// Get post by index
	var post models.Post
	initializers.DB.First(&post, id)

	// Return the response
	c.JSON(200, gin.H{
		"data": post,
		"message": "Post retrieved successfully!",
	})
}

func UpdatePost (c *gin.Context) {
	// Get index by URL parameter
	id := c.Param("id")

	// Get body
	var body struct {
		Title string 
		Body string
	}

	c.Bind(&body)

	// Find the post
	var post models.Post
	exists := initializers.DB.First(&post, id)

	if exists.Error != nil {
		c.JSON(404, gin.H{
			"message":"Post not found",
		})
		return
	}

	// Update the post
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	if result.Error != nil {
		c.JSON(501, gin.H{
			"message":"Failed to update the post",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": result,
		"message":"Post updated successfully",
	})
}

func DeletePost (c *gin.Context) {
	// Get index from URL parameters
	id := c.Param("id")

	// Delete post
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message":"Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":"Post deleted successfully",
	})
}