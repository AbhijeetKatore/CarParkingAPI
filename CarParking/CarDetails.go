package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
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
	defer notUnique("The Car Number you have entered already exists please Enter Correct one.", writer)

	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&car)
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	result, _ := collection.InsertOne(context.TODO(), car)

	if result != nil {
		fmt.Println("Car Details Inserted Successfully")
		fmt.Fprintln(writer, "Car Details Inserted Successfully ")
		CreateIndex(client, "CarDetails", "carnumber")
	}
}

func DeleteCarDetails(writer http.ResponseWriter, req *http.Request) {
	type request struct {
		CarNumber string
	}
	var resp request
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&resp)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")

	result, _ := collection.DeleteMany(context.TODO(), resp)

	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("Car Details Deleted Succesfully")
		fmt.Fprintln(writer, "Car Details Deleted Successfully ")
		CreateIndex(client, "CarDetails", "carnumber")
	}
}

func UpdateCarDetails(writer http.ResponseWriter, req *http.Request) {
	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&car)
	params := mux.Vars(req)
	car_number := params["car_number"]

	fmt.Println(car_number)
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	filter := bson.M{"carnumber": car_number}

	update := bson.M{"$set": bson.M{"carnumber": car.CarNumber, "carmodel": car.CarModel}}
	result, _ := collection.UpdateMany(context.TODO(), filter, update)
	if result.MatchedCount > 0 {
		if result.ModifiedCount > 0 {
			fmt.Println("Car Details Updated Succesfully")
			CreateIndex(client, "CarDetails", "carnumber")
			fmt.Fprintln(writer, "Car Details Updated Succesfully")
		} else {
			fmt.Println("Car Details not Updated")
			fmt.Fprintln(writer, "Car Details not Updated")
		}
	} else {
		fmt.Println("Car Details did not matched with existing details to update")
		fmt.Fprintln(writer, "Car Details did not matched with existing details to update")
	}
}
