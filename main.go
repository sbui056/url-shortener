package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go URL Shortener",
		})
	})

	err := r.Run(":9808")
	if err != nil {
		fmt.Printf("Failed to start server - Error: %v\n", err)
		return
	}
}