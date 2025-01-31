package main

import (
	"strconv"
	"time"

	"github.com/Abhishek-M-K/go-monitoring/initializers"
	"github.com/Abhishek-M-K/go-monitoring/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init(){
	initializers.LoadEnvs()	
	initializers.ConnectToDB()
	initializers.InitMetrics()
}

func main(){
  	router := gin.Default()

	// Metrics middleware
	router.Use(ReqMetrics)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Go Monitoring API!",
		})
	})

	// Metrics route
	router.GET("/metrics",gin.WrapH(promhttp.Handler()))

	postsGroup := router.Group("/posts")
	routes.PostRoutes(postsGroup)
  	router.Run() // listen and serve on 0.0.0.0:8080 by default unless a PORT environment variable was defined
}


func ReqMetrics (c *gin.Context){
	start := time.Now()

	c.Next()

	duration := time.Since(start).Seconds()

	// Increment the counter
	initializers.RequestCounter.WithLabelValues(
		c.Request.Method, c.FullPath(), strconv.Itoa(c.Writer.Status()),
	).Inc()

	// Observe the duration
	initializers.RequestMetrics.WithLabelValues(
		c.Request.Method, c.FullPath(), strconv.Itoa(c.Writer.Status()),
	).Observe(duration)
}