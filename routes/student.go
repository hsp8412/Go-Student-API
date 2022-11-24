package routes

import (
	"Go-Student-API/controllers"

	"github.com/gin-gonic/gin"
)

func StudentsRoute(r *gin.Engine){
	r.POST("/students", controllers.StudentsCreate)
	r.GET("/students", controllers.StudentsIndex)
	r.GET("/students/:id", controllers.StudentsShow)
	r.PUT("/students/:id", controllers.StudentsUpdate)
	r.DELETE("/students/:id",controllers.StudentsDelete)
}