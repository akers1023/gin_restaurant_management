package controllers

import (
	"context"
	"log"
	"net/http"
	collection "restaurant_management/controllers/collections"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := collection.TableCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing table items"})
		}
		var allTables []bson.M
		if err = result.All(ctx, &allTables); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allTables)
	}
}
