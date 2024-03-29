package CarParking

import (
	"context"
	"encoding/json"
	"fmt"
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
	decoder.Decode(&slot)
	defer notUnique("The Slot Details you have entered already exists please Enter Correct one.", writer)
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	// count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	// slot.UniqueSlotID = int(count) + 1
	result, _ := collection.InsertOne(context.TODO(), slot)
	if result != nil {
		fmt.Println("Slot Details Inserted Successfully")
		defer CreateIndex(client, "ParkingSlots", "uniqueslotid")
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
	result, _ := collection.DeleteMany(context.TODO(), resp)

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
	decoder.Decode(&slot)

	params := mux.Vars(req)
	uniqueslotid, _ := strconv.Atoi(params["_uniqueslotid"])
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"uniqueslotid": uniqueslotid}
	update := bson.M{"$set": bson.M{"floornumber": slot.FloorNumber, "uniqueslotnumber": slot.UniqueSlotNumber, "occupancy": slot.Occupancy}}

	result, _ := collection.UpdateMany(context.TODO(), filter, update)
	if result.MatchedCount > 0 {
		if result.ModifiedCount > 0 {
			fmt.Println("Slot Details Updated Succesfully")
			CreateIndex(client, "ParkingSlots", "uniqueslotid")
			fmt.Fprintln(writer, "Slot Details Updated Succesfully")

		} else {
			fmt.Println("Slot Details not Updated")
			fmt.Fprintln(writer, "Slot Details not Updated")
		}
	} else {
		fmt.Println("Slot Details did not matched with existing details to update")
		fmt.Fprintln(writer, "Slot Details did not matched with existing details to update")
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
			cursor.Decode(&slot)
			slots = append(slots, slot)
		}
		json.NewEncoder(writer).Encode(slots)
	}
}
