package CarParking

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewCarToSlot(writer http.ResponseWriter, req *http.Request) {
	var user UserDetails 
	var car CarDetails
	params:=mux.Vars(req)

	user_id,_ :=strconv.Atoi( params["_userid"])
	carnumber := params["carnumber"]
	uniqueslotid,_ :=strconv.Atoi( params["_uniqueslotid"])

	user,userFound := GetUserFromID(user_id)
	car,carFound := GetCarFromID(carnumber)
	slotFound := GetSlotFromID(uniqueslotid)
	
	if (userFound && carFound && slotFound){
		client := ConnectDatabase()
		collection := client.Database("CarParking").Collection("ParkingSlots")
		filter := bson.M{"uniqueslotid":uniqueslotid}
		update := bson.M{"$set": bson.M{"occupancy":true,"carnumber":car.CarNumber,"fname":user.FName,"lname":user.LName,"user_id":user.UserID ,"timein":time.Now()}}
		err := collection.FindOneAndUpdate(context.TODO(), filter, update)
		if err.Err()==mongo.ErrNoDocuments{
			fmt.Println("Slot Details not Found in Existing Database")
		}else{
			fmt.Println("Car Added to Slot Succesfully")
		}
	}
}

func GetUserFromID(user_id int) (UserDetails,bool){
	var user UserDetails
	found :=false 

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	cursor := collection.FindOne(context.TODO(), bson.M{"userid":user_id})
	cursor.Decode(&user)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Users Name Not Found in Existing Database please add User Details First or Enter a Correct One")
		return user, found
	}
	found = true
	return user, found
}

func GetCarFromID(carnumber string) (CarDetails,bool){
	var car CarDetails
	found :=false 
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	cursor:= collection.FindOne(context.TODO(), bson.M{"carnumber":carnumber})
	cursor.Decode(&car)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Car Number Not Found in Existing Database please add Car Details First or Enter a Correct One")
		return car,found
	}
	found = true
	return car,found
}

func GetSlotFromID(uniqueslotid int) (bool){
	var slot SlotDetails
	found :=false
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	cursor:= collection.FindOne(context.TODO(), bson.M{"uniqueslotid":uniqueslotid})
	cursor.Decode(&slot)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Slot Details Not Found in Existing Database please add New Slot Details First or Enter a Correct One")
		return found
	}
	found = true
	return found
}

func DeleteCarFromSlot(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// slot_id := params["_uniqueslotid"]
	uniqueslotid,_ := strconv.Atoi(params["_uniqueslotid"])

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter:=bson.D{{"uniqueslotid",uniqueslotid}}
	result,err:=collection.DeleteMany(context.TODO(),filter)
	if err != nil {
    	log.Fatal(err)
	}
	if result.DeletedCount == 0{
		fmt.Println("Car Details Not Found to Delete from Slot")
	}else{
		fmt.Println("Car Details Removed from Parking Slot")
	}
}
