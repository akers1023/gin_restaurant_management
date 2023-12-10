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

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := collection.OrderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}