package crawler

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	CheckError(err)

	database = client.Database("cv")
}

func saveProductDetail(productDetail *ProductDetail) error {

	collection := database.Collection("products")
	// Set client options
	// check if product exist
	var result ProductDetail
	err := collection.FindOne(context.TODO(), bson.M{"id": productDetail.ID}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("Inserting product detail", productDetail.ID)
		// insert product
		_, err := collection.InsertOne(context.TODO(), productDetail)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
