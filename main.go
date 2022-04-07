package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello go-api-git!")

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)

	router.Run(":8080")

}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "Response OK",
	})
}
