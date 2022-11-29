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
	"go.mongodb.org/mongo-driver/bson"
)

type UserDetails struct {
	// First Name of the user
	// in: string
	// required: true
	FName string `json:"FName"`
	// Last Name of the user
	// in: string
	// required: true
	LName string `json:"LName"`
	// Age of the user
	// in: integer
	// required: true
	Age int `json:"Age"`
	// ID of the user
	// in: integer
	// required: true
	UserID int `json:"UserID"`
}

func AddUser(writer http.ResponseWriter, res *http.Request) {
	// POST request
	//
	// # Insert documentation
	//
	// ---
	// produces:
	// - application/json
	//
	// parameters:
	//
	//   - name: Body
	//     in: body
	//     schema:
	//
	//         items:
	//             "$ref": "#/definitions/UserDetails"
	//         type: object
	//
	// responses:
	//
	//   '200':
	//        description: user response
	//        schema:
	//          type: object
	//          items:
	//            "$ref": "#/definitions/UserDetails"
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	defer notUnique("The User ID you have entered already exists please use another one.")
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
		defer CreateIndex(client, "Users", "userid")
		fmt.Fprintln(writer, "User Details Inserted Successfully and UserID is ", user.UserID)
	}
}
func notUnique(s string) {
	if r := recover(); r != nil {
		fmt.Println(s)
	}
}

//GET request
// # Insert documentation
// ---
// produces:
// - application/json
//
// responses:
//
//   '200':
//        description: user response
//        schema:
//          type: object
//          items:
//            "$ref": "#/definitions/UserDetails"

func GetUser(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	isPageQuery := query.Has("page")
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	var users []UserDetails
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	if isPageQuery {
		page, _ := strconv.ParseInt(query["page"][0], 10, 64)
		var limit int64 = 10
		if page > 0 {
			_, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Filter(bson.M{}).Decode(&users).Find()
			if err != nil {
				panic(err)
			}
			if users == nil {
				fmt.Fprintln(writer, "You have reached the END of the Users List")
			} else {
				json.NewEncoder(writer).Encode(users)
				temp, _ := strconv.Atoi(query["page"][0])
				temp += 1
				currPage := strconv.Itoa(temp)
				nextPage := strings.Replace("/user?page=0", "0", currPage, 12)
				writer.Header().Set("Content-Type", "application/json")
				fmt.Fprintln(writer, "", nextPage)
			}
		} else {
			fmt.Fprintln(writer, "This page does not exist.")
		}
	} else {
		// fmt.Fprintln(writer, "This page does not exist.")
		cursor, _ := collection.Find(context.TODO(), bson.M{})
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var user UserDetails
			err := cursor.Decode(&user)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
			json.NewEncoder(writer).Encode(users)
		}
	}
}

func DeleteUser(writer http.ResponseWriter, req *http.Request) {
	type response struct {
		UserID int
	}
	var resp response
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&resp)
	if err != nil {
		fmt.Println(err)
	}
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "*")
	client := ConnectDatabase()
	collection := client.Database("CarParking").Collection("Users")
	result, err := collection.DeleteMany(context.TODO(), resp)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount == 0 {
		fmt.Println("Data didn't Match to Delete")
	} else {
		fmt.Println("User Details Deleted Succesfully")
		CreateIndex(client, "Users", "userid")
		fmt.Fprintln(writer, "User Details Deleted Succesfully")
	}
}
