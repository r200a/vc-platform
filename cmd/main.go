package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status":  "OK",
			"service": "VC",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"VC": "test",
		})
	})
	fmt.Println("Server starting on:8085")
	r.Run(":8085")
}
