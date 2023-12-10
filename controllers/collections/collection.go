package collection

import (
	"restaurant_management/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)


var Validate = validator.New()
var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var MenuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var FoodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var InvoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")
var OrderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var OrderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")
var TableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")
