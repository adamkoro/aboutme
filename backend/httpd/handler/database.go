package handler

import (
	"net/http"

	database "github.com/adamkoro/aboutme/backend/database/mongo"
	"github.com/adamkoro/aboutme/backend/errorHandler"
	"github.com/gin-gonic/gin"
)

func PingDb() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client, c, err := database.CreateConnection(2, "dev", "dev", "localhost", 27017)
		errorExist := errorHandler.IsError(err)
		if errorExist {
			errorHandler.ErrorLogger.Println(err)
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error occurred while accessing the database",
			})
		}
		errorExist = database.PingDb(client, c)
		if errorExist {
			errorHandler.ErrorLogger.Println(err)
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error occurred while trying to ping the database",
			})
		}

		if !errorExist {
			ctx.JSON(http.StatusOK, map[string]string{
				"message": "Database is successfully pinged",
			})
		}
		database.DeleteConnection(client, c)
	}
}

//func TestData() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		//client, ctx := database.CreteConnection(2, "dev", "dev", "localhost", 27017)
//		//database.ConnectToDB(client, ctx)
//		//errorExist := database.PingDb(client, ctx)
//
//		//collection := client.Database("aboutme").Collection("pictures")
//		result, err := collection.Find(*ctx, bson.M{})
//		errorExist := errorHandler.IsError(err)
//		if errorExist {
//			errorHandler.ErrorLogger.Println(err)
//		}
//		var testdata []bson.M
//		err = result.All(*ctx, &testdata)
//		log.Println(testdata)
//
//		if !errorExist {
//			c.JSON(http.StatusOK, map[string]string{
//				"message": "Successfully getting data",
//			})
//		} else {
//			c.JSON(http.StatusInternalServerError, map[string]string{
//				"message": "Error occurred while getting data",
//			})
//		}
//		database.DisconnectFromDb(client, ctx)
//	}
//}
