package main

import (
	"fmt"
	"github.com/sbui056/url-shortener/handler"
	"github.com/sbui056/url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go URL Shortener",
		})
	})

	r.POST("/create-short-rul", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		fmt.Printf("Failed to start server - Error: %v\n", err)
		return
	}
}