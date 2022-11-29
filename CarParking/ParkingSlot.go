package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type SlotDetails struct {
	FloorNumber      int  `json:"FloorNumber"`
	UniqueSlotNumber int  `json:"UniqueSlotNumber"`
	Occupancy        bool `json:"Occupancy"`
	UniqueSlotID     int  `json:"UniqueSlotID"`
}

func AddParkingSlot(writer http.ResponseWriter, req *http.Request) {
	var slot SlotDetails
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&slot)

	if err != nil {
		fmt.Println(err)
	}
	defer notUnique("The Slot Details you have entered already exists please Enter Correct one.")

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
		fmt.Fprintln(writer, "Slot Details Inserted Successfully ")
	}
}

func DeleteParkingSlots(writer http.ResponseWriter, req *http.Request) {
	type response struct {
		UniqueSlotID int
	}
	decoder := json.NewDecoder(req.Body)
	var resp response
	decoder.Decode(&resp)
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	// params := mux.Vars(req)
	// uniqueslotid, err := strconv.Atoi(params["_uniqueslotid"])
	// filter := bson.M{"uniqueslotid": uniqueslotid}
	result, err := collection.DeleteMany(context.TODO(), resp)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("Slot Details Deleted Succesfully")
		CreateIndex(client, "ParkingSlots", "uniqueslotid")
		fmt.Fprintln(writer, "Slot Details Deleted Succesfully")

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
	if result.ModifiedCount > 0 {
		fmt.Println("Slot Details Updated Succesfully")
		CreateIndex(client, "ParkingSlots", "uniqueslotid")
		fmt.Fprintln(writer, "Slot Details Updated Succesfully")

	}
}

func GetFreeParkingSlots(writer http.ResponseWriter, req *http.Request) {
	EnableCors(&writer)
	query := req.URL.Query()
	isPageQuery := query.Has("page")
	writer.Header().Set("Content-Type", "application/json")

	var slots []SlotDetails
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"occupancy": false}

	if isPageQuery {
		page, _ := strconv.ParseInt(query["page"][0], 10, 64)
		var limit int64 = 10
		if page > 0 {
			_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(filter).Decode(&slots).Find()
			if err != nil {
				panic(err)
			}
			if slots == nil {
				fmt.Fprintln(writer, "You have reached the END of the Slots List")
			} else {
				json.NewEncoder(writer).Encode(slots)
				temp, _ := strconv.Atoi(query["page"][0])
				temp += 1
				currPage := strconv.Itoa(temp)
				nextPage := strings.Replace("/slot?page=0", "0", currPage, 12)
				writer.Header().Set("Content-Type", "application/json")
				fmt.Fprintln(writer, "", nextPage)
			}
		} else {
			fmt.Fprintln(writer, "This page does not exist.")
		}
	} else {
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
}
