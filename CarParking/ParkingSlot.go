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

type SlotDetails struct {
	FloorNumber      int
	UniqueSlotNumber int
	Occupancy bool
}

func AddParkingSlot(writer http.ResponseWriter, req *http.Request) {
	var slot SlotDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&slot)

	if err != nil {
		fmt.Println(err)
	}

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	result, err := collection.InsertOne(context.TODO(), slot)

	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Printf("Slot Details Inserted Successfully")
	}
}

func DeleteParkingSlots(writer http.ResponseWriter, req *http.Request){
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
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
		fmt.Printf("Slot Details Deleted Succesfully")
	}
}
func UpdateParkingSlot(writer http.ResponseWriter, req *http.Request){
	var slot SlotDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&slot)
	
	params := mux.Vars(req)
	_id := params["_id"]
	pid, _ := primitive.ObjectIDFromHex(_id)
	if err!=nil{
		fmt.Println(err)
	}

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"_id": pid}
	update := bson.M{"$set": bson.M{"floornumber": slot.FloorNumber, "uniqueslotnumber": slot.UniqueSlotNumber,"occupancy":slot.Occupancy}}

	result,err := collection.UpdateMany(context.TODO(), filter, update)
	if result!=nil{
		fmt.Println("Slot Details Updated Succesfully")
	}
}

func GetFreeParkingSlots(writer http.ResponseWriter, req *http.Request){
	var slots []SlotDetails
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"occupancy": false}
	cursor, _ := collection.Find(context.TODO(), filter)
	
	for cursor.Next(context.TODO()) {
		var slot SlotDetails
		err := cursor.Decode(&slot)
		if err != nil {
			log.Fatal(err)
		}
		slots = append(slots, slot)
	}
	json.NewEncoder(writer).Encode(slots)

}


