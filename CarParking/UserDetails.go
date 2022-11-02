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



type UserDetails struct{
	FName string
	LName string
	Age int
}


func AddUser(writer http.ResponseWriter, res *http.Request) {
	var user UserDetails
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&user)
	
	if err !=nil{
		fmt.Println(err)
	}

	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	result,err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}
	if result!= nil{
			fmt.Printf("User Details Inserted Successfully")
	}
}

func GetUser(writer http.ResponseWriter, res *http.Request) {

	var users []UserDetails
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	cursor, _ := collection.Find(context.TODO(), bson.M{})

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user UserDetails
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	json.NewEncoder(writer).Encode(users)

}


func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
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
		fmt.Printf("User Details Deleted Succesfully")
	}
}