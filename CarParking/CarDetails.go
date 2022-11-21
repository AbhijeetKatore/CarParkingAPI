package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type CarDetails struct {
	CarNumber string `json:"CarNumber"`
	CarModel  string `json:"CarModel"`
}

func AddCarDetails(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "*")

	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)

	if err != nil {
		fmt.Println(err)
	}
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	result, err := collection.InsertOne(context.TODO(), car)
	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println("Car Details Inserted Successfully")
		CreateIndex(client, "CarDetails", "carnumber")
	}
}

func DeleteCarDetails(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	params := mux.Vars(req)
	car_number := params["car_number"]

	filter := bson.M{"carnumber": car_number}
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("Car Details Deleted Succesfully")
		CreateIndex(client, "CarDetails", "carnumber")
	}
}

func UpdateCarDetails(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	params := mux.Vars(req)
	car_number := params["car_number"]
	if err != nil {
		fmt.Println(err)
	}
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	filter := bson.M{"carnumber": car_number}

	update := bson.M{"$set": bson.M{"carnumber": car.CarNumber, "carmodel": car.CarModel}}
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if result != nil {
		fmt.Println("Data Updated Succesfully")
		CreateIndex(client, "CarDetails", "carnumber")
	}
}
