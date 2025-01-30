package routes

import (
	"github.com/Abhishek-M-K/go-monitoring/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.RouterGroup){
	router.GET("/", controllers.GetPosts)
	router.POST("/", controllers.Posts)
	router.GET("/:id", controllers.GetPostByIndex)
	router.PATCH("/:id", controllers.UpdatePost)
	router.DELETE("/:id", controllers.DeletePost)
}