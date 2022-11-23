package controllers

import (
	"Go-Student-API/initializers"
	"Go-Student-API/models"
	"time"

	"github.com/gin-gonic/gin"
)

func StudentsCreate(c *gin.Context){

	student := models.Student{
		FirstName: "Steven",
		MiddleName: "",
		LastName: "He",
		DOB: time.Date(1998, time.Month(8), 4, 0, 0, 0, 0, time.UTC),
		StudentID: "30113342",
		Age: 24,
		IsCurrentStudent: true,
		Gender: "male",
	}

	result := initializers.DB.Create(&student) 

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"student": student,
	})
}