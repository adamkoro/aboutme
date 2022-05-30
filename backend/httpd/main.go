package main

import (
	"github.com/adamkoro/aboutme/backend/errorHandler"
	"github.com/adamkoro/aboutme/backend/httpd/initdb"
)

func main() {
	client, ctx, err := initdb.CreateConnection()
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
	}
	initdb.CreateDb(client, ctx)

	//gin.SetMode(gin.ReleaseMode)
	//server := gin.Default()

	//server.GET("/ping", handler.Ping())
	//server.GET("/pingdb", handler.PingDb())
	//server.GET("/testdata", handler.TestData())

	//server.Run(":8080")
}
