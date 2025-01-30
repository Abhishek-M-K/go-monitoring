package main

import (
	"github.com/Abhishek-M-K/go-monitoring/initializers"
	"github.com/Abhishek-M-K/go-monitoring/models"
)

func init(){
	initializers.LoadEnvs()	
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}