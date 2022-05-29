package mongo

import (
	"context"

	errorHandler "github.com/adamkoro/aboutme/backend/errorHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func PingDb(client *mongo.Client, ctx *context.Context) bool {
	err := client.Ping(*ctx, readpref.Primary())
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
	}
	return errorExist
}
