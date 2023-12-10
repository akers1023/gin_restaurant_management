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

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order

		err := collection.OrderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the orders"})
		}
		c.JSON(http.StatusOK, order)
	}
}