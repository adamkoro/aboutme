package mongo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	errorHandler "github.com/adamkoro/aboutme/backend/errorHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection setup functions

func createConnString(username string, password string, address string, port string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", username, password, address, port)
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

func CreateConnection(timeout string, username string, password string, address string, port string) (*mongo.Client, *context.Context, error) {
	connString := createConnString(username, password, address, port)
	client, err := createClient(connString)
	errorExist := errorHandler.IsError(err)
	if errorExist {
		return nil, nil, err
	}

	timeOut, err := strconv.Atoi(timeout)
	errorExist = errorHandler.IsError(err)
	if errorExist {
		return nil, nil, err
	}

	ctx, err := setTimeout(time.Duration(timeOut))
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
	defer client.Disconnect(*ctx)
}
