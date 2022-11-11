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

type UserDetails struct {
	FName  string
	LName  string
	Age    int
	UserID int
}

func AddUser(writer http.ResponseWriter, res *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "*")

	var user UserDetails
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	user.UserID = int(count) + 1
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println("User Details Inserted Successfully and UserID is ", user.UserID)
		CreateIndex(client, "Users", "userid")
	}
}

func GetUser(writer http.ResponseWriter, res *http.Request) {
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

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
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)

}

func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	params := mux.Vars(req)
	_userid := params["_userid"]
	userid, err := strconv.Atoi(_userid)
	filter := bson.M{"userid": userid}
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("User Details Deleted Succesfully")
		CreateIndex(client, "Users", "userid")
	}
}
