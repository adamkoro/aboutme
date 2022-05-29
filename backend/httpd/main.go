package main

import (
	database "github.com/adamkoro/aboutme/backend/database/mongo"
)

func main() {
	//server := gin.Default()

	client, ctx := database.CreteConnection(5, "dev", "dev", "localhost", 27017)
	database.ConnectToDB(client, ctx)
	database.PingDb(client, ctx)
	database.DisconnectFromDb(client, ctx)

	//server.GET("/ping", handler.Ping())

	//server.Run(":8080")
}
