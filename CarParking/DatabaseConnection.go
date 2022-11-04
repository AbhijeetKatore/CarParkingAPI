package CarParking

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CreateIndex(client *mongo.Client,collec string,field string){
	collection := client.Database("CarParking").Collection(collec)
	indexModel := mongo.IndexModel{
			Keys: bson.M{
			field: 1, 
			}, Options: nil,
		}
	
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) 
		}
}