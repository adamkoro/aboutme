package mongo

import (
	"context"
	"fmt"
	"time"

	errorHandler "github.com/adamkoro/aboutme/backend/errorHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection setup functions

func createClient(username string, password string, address string, port int) *mongo.Client {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, address, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
	}
	return client
}

func setTimeout(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)
	return ctx
}

// Public functions

func DisconnectFromDb(client *mongo.Client, ctx *context.Context) {
	defer client.Disconnect(*ctx)
}

func CreteConnection(timeout int, username string, password string, address string, port int) (*mongo.Client, *context.Context) {
	client := createClient(username, password, address, port)
	ctx := setTimeout(time.Duration(timeout))
	return client, &ctx
}

func ConnectToDB(client *mongo.Client, ctx *context.Context) {
	err := client.Connect(*ctx)
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
	}
}
