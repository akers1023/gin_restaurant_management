package controllers

import (
	"context"
	"net/http"
	collection "restaurant_management/controllers/collections"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table models.Table

		err := collection.TableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the tables"})
		}
		c.JSON(http.StatusOK, table)
	}
}