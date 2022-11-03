package CarParking

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewCarToSlot(writer http.ResponseWriter, req *http.Request) {
	var user UserDetails 
	var car CarDetails
	params:=mux.Vars(req)

	user_id := params["user_id"]
	car_id := params["car_id"]
	slot_id :=params["slot_id"]
	user_idobj, _ := primitive.ObjectIDFromHex(user_id)
	car_idobj, _ := primitive.ObjectIDFromHex(car_id)	
	slot_idobj, _ := primitive.ObjectIDFromHex(slot_id)

	user,userfound = GetUserFromID(user_idobj)
	car,carfound = GetCarFromID(car_idobj)
	slotfound = GetSlotFromID(slot_idobj)
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	filter := bson.M{"_id":slot_idobj}
	update := bson.M{"$set": bson.M{"occupancy":true,"carnumber":car.CarNumber,"fname":user.FName,"lname":user.LName ,"timein":time.Now()}}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if err.Err()==mongo.ErrNoDocuments{
		fmt.Println("Slot Details not Found in Existing Database")
	}else{
		fmt.Println("Car Added to Slot Succesfully")
	}
}

func GetUserFromID(user_idobj primitive.ObjectIDt) (UserDetails,bool){
	var user UserDetails
	found :=true 

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	cursor := collection.FindOne(context.TODO(), bson.M{"_id":user_idobj})
	cursor.Decode(&user)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Users Name Not Found in Existing Database please add User Details First or Enter a Correct One")
	}
	return user, found = true
}

func GetCarFromID(car_idobj primitive.ObjectID) (CarDetails,bool){
	var car CarDetails
	found :=false 

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("CarDetails")
	cursor:= collection.FindOne(context.TODO(), bson.M{"_id":car_idobj})
	cursor.Decode(&car)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Car Number Not Found in Existing Database please add Car Details First or Enter a Correct One")
		return found

	}
	return car,found = true
}

func GetSlotFromID(car_idobj primitive.ObjectID) (bool){
	found :=false
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("ParkingSlots")
	cursor:= collection.FindOne(context.TODO(), bson.M{"_id":slot_idobj})
	cursor.Decode(&slot)
	if cursor.Err() == mongo.ErrNoDocuments {
		fmt.Println("Slot Details Not Found in Existing Database please add New Slot Details First or Enter a Correct One")
		return found
	}
	return found = true
}