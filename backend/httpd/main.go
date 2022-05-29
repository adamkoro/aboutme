package main

import (
	"github.com/adamkoro/aboutme/backend/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/ping", handler.Ping())
	server.GET("/pingdb", handler.TestDB())

	server.Run(":8080")
}
