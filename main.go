package main

import (
	"Go-Student-API/controllers"
	"Go-Student-API/initializers"

	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/students", controllers.StudentsCreate)
	r.Run() 
}