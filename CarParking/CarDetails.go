package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarDetails struct {
	CarNumber string
	CarModel  string
}

func AddCarDetails(writer http.ResponseWriter, req *http.Request) {
	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	
	if err !=nil{
		fmt.Println(err)
	}
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	result,err := collection.InsertOne(context.TODO(), car)
	if err != nil {
		panic(err)
	}
	if result!= nil{
			fmt.Printf("Car Details Inserted Successfully")
	}

}


func DeleteCarDetails(writer http.ResponseWriter, req *http.Request){
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	params := mux.Vars(req)
	_id := params["_id"]
	pid, _ := primitive.ObjectIDFromHex(_id)

	filter := bson.M{"_id": pid}
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Data didn't Match to Delete")
	}else{
		fmt.Printf("Car Details Deleted Succesfully")
	}
}

func UpdateCarDetails(writer http.ResponseWriter, req *http.Request){
	var car CarDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&car)
	
	params := mux.Vars(req)
	_id := params["_id"]
	pid, _ := primitive.ObjectIDFromHex(_id)

	
	if err!=nil{
		fmt.Println(err)
	}

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	filter := bson.M{"_id": pid}
	update := bson.M{"$set": bson.M{"carnumber": car.CarNumber, "carmodel": car.CarModel}}

	result,err := collection.UpdateMany(context.TODO(), filter, update)
	fmt.Println(result, err)
	if result!=nil{
		fmt.Println("Data Updated Succesfully")
	}
}
