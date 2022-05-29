package handler

import (
	"net/http"

	database "github.com/adamkoro/aboutme/backend/database/mongo"
	"github.com/gin-gonic/gin"
)

func TestDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, ctx := database.CreteConnection(2, "dev", "dev", "localhost", 27017)
		database.ConnectToDB(client, ctx)
		errorExist := database.PingDb(client, ctx)
		database.DisconnectFromDb(client, ctx)

		if !errorExist {
			c.JSON(http.StatusOK, map[string]string{
				"message": "Successfully pinged the database",
			})
		} else {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Error occurred while trying to ping the database",
			})
		}
	}
}
