package controllers

import (
	"Go-Student-API/initializers"
	"Go-Student-API/models"
	"time"

	"github.com/gin-gonic/gin"
)

func StudentsCreate(c *gin.Context){
	var body struct{
		FirstName string
		MiddleName string
		LastName string
		StudentID string
		Age uint8
		DOB time.Time
		Gender string
		IsCurrentStudent bool
	}

	c.Bind(&body)

	student := models.Student{
		FirstName: body.FirstName,
		MiddleName: body.MiddleName,
		LastName: body.LastName,
		DOB: body.DOB,
		StudentID: body.StudentID,
		Age: body.Age,
		IsCurrentStudent: body.IsCurrentStudent,
		Gender: body.Gender,
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

func StudentsIndex(c *gin.Context){
	var students []models.Student
	initializers.DB.Find(&students)

	c.JSON(200, gin.H{
		"students":students,
	})
}

func StudentsShow(c *gin.Context){
	id := c.Param("id")

	var student models.Student
	initializers.DB.First(&student, id)

	c.JSON(200, gin.H{
		"student":student,
	})
}

func StudentsUpdate(c *gin.Context){
	id := c.Param("id")
	var body struct{
		FirstName string
		MiddleName string
		LastName string
		StudentID string
		Age uint8
		DOB time.Time
		Gender string
		IsCurrentStudent bool
	}

	c.Bind(&body)

	var student models.Student
	initializers.DB.First(&student, id)

	initializers.DB.Model(&student).Updates(models.Student{
		FirstName: body.FirstName,
		MiddleName: body.MiddleName,
		LastName: body.LastName,
		DOB: body.DOB,
		StudentID: body.StudentID,
		Age: body.Age,
		IsCurrentStudent: body.IsCurrentStudent,
		Gender: body.Gender,
	})

	c.JSON(200, gin.H{
		"student":student,
	})
}

func StudentsDelete(c *gin.Context){
	id := c.Param("id")

	var student models.Student
	initializers.DB.First(&student, id)

	initializers.DB.Delete(&models.Student{}, id)

	c.JSON(200, gin.H{
		"student":student,
	})
}