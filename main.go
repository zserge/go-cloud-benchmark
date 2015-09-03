package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":" + port)
}
