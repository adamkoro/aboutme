package main

import (
	"github.com/adamkoro/aboutme/backend/database-api/httpd/env"
	"github.com/adamkoro/aboutme/backend/database-api/httpd/handler"
	"github.com/adamkoro/aboutme/backend/database-api/httpd/initdb"
	"github.com/gin-gonic/gin"
)

func main() {
	env.CheckEnvs()
	client, ctx, _ := initdb.CreateConnection()

	initdb.CreateDb(client, ctx)

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/ping", handler.Ping())
	server.GET("/pingdb", handler.PingDb())

	server.Run(":8080")
}
