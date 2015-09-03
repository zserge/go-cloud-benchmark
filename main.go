package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	addr := ":8080"
	if env := os.Getenv("PORT"); env != "" {
		addr = ":" + env
	}
	if env := os.Getenv("OPENSHIFT_GO_PORT"); env != "" {
		addr = os.Getenv("OPENSHIFT_GO_IP") + ":" + env
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(addr)
}
