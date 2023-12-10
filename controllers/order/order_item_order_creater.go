package controllers

import (
	collection "restaurant_management/controllers/collections"
	"restaurant_management/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrderItemOrderCreator(order models.Order) string {

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	collection.OrderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}