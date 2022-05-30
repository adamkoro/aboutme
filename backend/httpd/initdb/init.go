package initdb

import (
	"context"
	"time"

	database "github.com/adamkoro/aboutme/backend/database/mongo"
	"github.com/adamkoro/aboutme/backend/errorHandler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const databaseName = "aboutme"

func getDatabases(client *mongo.Client, ctx context.Context) ([]string, error) {
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
		return nil, err
	}
	return databases, nil
}

func searchDb(database []string) string {
	for _, db := range database {
		if db == databaseName {
			return db
		}
		continue
	}
	return ""
}

func isDbExist(client *mongo.Client, ctx *context.Context) (bool, error) {
	dbs, err := getDatabases(client, *ctx)
	errorExist := errorHandler.IsError(err)
	if errorExist {
		return false, err
	}

	db := searchDb(dbs)
	if db == databaseName {
		errorHandler.InfoLogger.Println("Database is exist")
		database.DeleteConnection(client, ctx)
		return true, nil
	}
	errorHandler.InfoLogger.Println("Database is not exist")
	return false, nil
}

func CreateConnection() (*mongo.Client, *context.Context, error) {
	client, ctx, err := database.CreateConnection(2, "dev", "dev", "localhost", 27017)
	errorExist := errorHandler.IsError(err)
	if errorExist {
		errorHandler.ErrorLogger.Println(err)
		return nil, nil, err
	}
	errorHandler.InfoLogger.Println("Connection to MongoDB is established")
	return client, ctx, nil
}

func CreateDb(client *mongo.Client, ctx *context.Context) error {
	isDbExist, _ := isDbExist(client, ctx)
	if !isDbExist {
		errorHandler.InfoLogger.Println("Database is creating")
		currentTime := time.Now()
		aboutmeDatabase := client.Database(databaseName)
		picturesCollections := aboutmeDatabase.Collection("metadata")
		_, err := picturesCollections.InsertOne(*ctx, bson.D{
			{Key: "databaseIsCreated", Value: currentTime.String()},
		})
		errorExist := errorHandler.IsError(err)
		if errorExist {
			errorHandler.ErrorLogger.Println(err)
			return err
		}
		database.DeleteConnection(client, ctx)
		return nil
	}

	return nil
}
