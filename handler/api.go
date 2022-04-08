package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentGetHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}
