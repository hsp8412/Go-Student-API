package main

import (
	"Go-Student-API/initializers"
	"Go-Student-API/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Student{})
}