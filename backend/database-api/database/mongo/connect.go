package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/adamkoro/aboutme/backend/database-api/errorHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection setup functions

func createConnString(username string, password string, address string, port int) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, address, port)
}

func createClient(connectionString string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
		return nil, err
	}
	return client, nil
}

func setTimeout(timeout time.Duration) (context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout*time.Second)
	return ctx, nil
}

// Public functions

func CreateConnection(timeout int, username string, password string, address string, port int) (*mongo.Client, *context.Context, error) {
	connString := createConnString(username, password, address, port)
	client, err := createClient(connString)
	errorExist := errorHandler.IsError(err)
	if errorExist {
		return nil, nil, err
	}

	ctx, err := setTimeout(time.Duration(timeout))
	errorExist = errorHandler.IsError(err)
	if errorExist {
		return nil, nil, err
	}

	err = client.Connect(ctx)
	errorExist = errorHandler.IsError(err)
	if errorExist {
		return nil, nil, err
	}

	return client, &ctx, nil
}

func DeleteConnection(client *mongo.Client, ctx *context.Context) {
	client.Disconnect(*ctx)
}
