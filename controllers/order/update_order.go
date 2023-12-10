package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	collection "restaurant_management/controllers/collections"
	"restaurant_management/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := collection.MenuCollection.FindOne(ctx, bson.M{"tabled_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "menu", Value: order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := collection.OrderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$st", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}