package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/pkg/config"
	"github.com/r200a/vc-platform/storage/db"
)

func main() {
	// Load Config,
	cfg := config.Load()
	// making git main
	// Database Setup
	db.Connect(cfg.DBURL)

	// Setup Router
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

	// Setup Server
	fmt.Println("Server starting on:", cfg.Port)
	r.Run(":" + cfg.Port)
}
