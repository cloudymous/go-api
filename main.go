package main

import (
	"fmt"
	"go-api-git/handler"
	"go-api-git/student"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello go-api-git!")

	//student := student.Student{}

	dsn := "root:password@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Database Connected!")

	var student student.Student
	err = db.First(&student).Error
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println("Name :", student.Name)
	fmt.Println("Major :", student.Major)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/student/:id", handler.StudentGetHandler)

	router.Run(":8080")

}
