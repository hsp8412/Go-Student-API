package main

import (
	"Go-Student-API/initializers"
	"Go-Student-API/routes"

	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.String(200, "Welcome to students api.")
	})
	routes.StudentsRoute(r)
	r.Run() 
}


