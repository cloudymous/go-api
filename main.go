package main

import (
	"fmt"
	"go-api-git/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello go-api-git!")

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/student/:id", handler.StudentGetHandler)

	router.Run(":8080")

}
