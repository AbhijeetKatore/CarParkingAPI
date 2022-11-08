package main

import (
	"CarParking/CarParking"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestServer(test *testing.T){
	req,err:=http.NewRequest("GET","http://localhost:8080/user",nil)
		if err!=nil{
			test.Fatalf(" Not Able to Create Request %v",err)
		}
	recorder := httptest.NewRecorder()
	CarParking.GetFreeParkingSlots(recorder,req)

	res:=recorder.Result()
	test.Log(res.StatusCode,"a")
	if res.StatusCode != http.StatusOK{
		test.Fatalf("Expected Status OK but got %v",res.StatusCode)
	}
	op,_ := ioutil.ReadAll(res.Body)	
	test.Log(op)
}

// func TestGetUser(test *testing.T){
// 	req,err:=http.NewRequest("GET","http://localhost:8080/user",nil)
// 	if err!=nil{
// 		test.Fatalf("Request Not Able to Create %v",err)
// 	}
// 	recorder := httptest.NewRecorder()
// 	GetUser(recorder,req)
// 	res:=recorder.Result()
// 	defer res.Body.Close()
// 	if res.StatusCode != http.StatusOK{
// 		test.Fatalf("Expected Status OK but got %v",res.StatusCode)
// 	}
// 	op,_ := ioutil.ReadAll(res.Body)	
// 	test.Log(op)

// }

// func GetUser(writer http.ResponseWriter, res *http.Request) {

// 	var users []CarParking.UserDetails
// 	client := CarParking.ConnectDatabase()
// 	collection := client.Database("CarParking").Collection("Users")
// 	cursor, _ := collection.Find(context.TODO(), bson.M{})

// 	defer cursor.Close(context.TODO())

// 	for cursor.Next(context.TODO()) {
// 		var user CarParking.UserDetails
// 		err := cursor.Decode(&user)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, user)
// 	}
// 	writer.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(writer).Encode(users)
// }