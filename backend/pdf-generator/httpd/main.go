package main

import (
	"github.com/adamkoro/aboutme/backend/pdf-generator/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping", handler.Ping())
	server.Run(":8080")
}
