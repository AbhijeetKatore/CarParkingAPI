package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type SlotDetails struct {
	FloorNumber      int
	UniqueSlotNumber int
	Occupancy        bool
	UniqueSlotID     int
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
	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	slot.UniqueSlotID = int(count) + 1
	result, err := collection.InsertOne(context.TODO(), slot)

	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println("Slot Details Inserted Successfully")
		CreateIndex(client, "ParkingSlots", "uniqueslotid")

	}

}

func DeleteParkingSlots(writer http.ResponseWriter, req *http.Request) {
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	params := mux.Vars(req)
	uniqueslotid, err := strconv.Atoi(params["_uniqueslotid"])
	filter := bson.M{"uniqueslotid": uniqueslotid}
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("Slot Details Deleted Succesfully")
		CreateIndex(client, "ParkingSlots", "uniqueslotid")

	}
}
func UpdateParkingSlot(writer http.ResponseWriter, req *http.Request) {
	var slot SlotDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&slot)
	if err != nil {
		fmt.Println(err)
	}
	params := mux.Vars(req)
	uniqueslotid, err := strconv.Atoi(params["_uniqueslotid"])
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"uniqueslotid": uniqueslotid}
	update := bson.M{"$set": bson.M{"floornumber": slot.FloorNumber, "uniqueslotnumber": slot.UniqueSlotNumber, "occupancy": slot.Occupancy}}

	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if result != nil {
		fmt.Println("Slot Details Updated Succesfully")
		CreateIndex(client, "ParkingSlots", "uniqueslotid")
	}
}

func GetFreeParkingSlots(writer http.ResponseWriter, req *http.Request) {
	EnableCors(&writer)
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
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(slots)
}
