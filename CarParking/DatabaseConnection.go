package CarParking

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongo_db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client

}

func CreateIndex(client *mongo.Client, collec string, field string) {
	collection := client.Database("CarParking").Collection(collec)
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			field: -1,
		}, Options: options.Index().SetUnique(true),
	}
	collection.Indexes().CreateOne(context.TODO(), indexModel)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func notUnique(s string, w http.ResponseWriter) {
	if r := recover(); r != nil {
		fmt.Println(s)
		fmt.Fprintln(w, s)
	}
}
