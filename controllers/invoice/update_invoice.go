package controllers

import (
	"context"
	"fmt"
	"net/http"
	collection "restaurant_management/controllers/collections"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		invoiceId := c.Param("invoice_id")

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"invoice_id": invoiceId}

		var updateObj primitive.D

		if invoice.Payment_method != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_method", Value: invoice.Payment_method})
		}

		if invoice.Payment_status != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_status", Value: invoice.Payment_status})
		}

		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: invoice.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		result, err := collection.InvoiceCollection.UpdateOne(
			ctx,
			filter,
			bson.D{

				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		if err != nil {
			msg := fmt.Sprintf("invoice item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}