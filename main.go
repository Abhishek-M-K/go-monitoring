package main

import (
	"github.com/Abhishek-M-K/go-monitoring/initializers"
	"github.com/Abhishek-M-K/go-monitoring/routes"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvs()	
	initializers.ConnectToDB()
}

func main(){
  	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Go Monitoring API!",
		})
	})

	postsGroup := router.Group("/posts")
	routes.PostRoutes(postsGroup)
  	router.Run() // listen and serve on 0.0.0.0:8080 by default unless a PORT environment variable was defined
}