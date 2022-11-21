package main

import (
	"net/http"
	"strings"
	"testing"
)

var client = &http.Client{}

func TestServer(test *testing.T) {
	httpreq, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Expected Status OK but got %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

// Test User
func TestGetUser(test *testing.T) {

	httpreq, _ := http.NewRequest("GET", "http://localhost:8080/user", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Expected Status OK but got %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}

	//OR

	// var req *http.Request
	// var recorder = httptest.NewRecorder()
	//,err:=http.NewRequest("POST","http://local",nil)
	// if err!=nil{
	// 	test.Fatalf(" Not Able to Create Request %v",err)
	// }
	// CarParking.GetUser(recorder, req)
	// res := recorder.Result()

	// op, _ := ioutil.ReadAll(res.Body)
	// test.Log(string(op))
}

func TestAddUser(test *testing.T) {
	body := strings.NewReader(`{
			"FName": "TestData",
			"LName": "TestData",
			"Age": 111
		}`)
	httpreq, err := http.NewRequest("POST", "http://localhost:8080/user", body)
	client := &http.Client{}
	res, err := client.Do(httpreq)

	if res.StatusCode != http.StatusOK {
		test.Fatalf("Expected Status OK but got %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}

	//             OR

	// body := strings.NewReader(`{
	// 	"FName": "TestData",
	// 	"LName": "TestData",
	// 	"Age": 111
	// }`)
	// httpreq, _ := http.NewRequest("POST", "", body)

	// CarParking.AddUser(recorder, httpreq)
	// res := recorder.Result()
	// if res.StatusCode != http.StatusOK {
	// 	test.Fatalf("User Details Not able to Add error: %v", res.StatusCode)
	// }
}

func TestDeleteUser(test *testing.T) {

	httpreq, err := http.NewRequest("DELETE", "http://localhost:8080/user/16", nil)
	client := &http.Client{}
	res, err := client.Do(httpreq)
	//CarParking.DeleteUser(recorder, httpreq)
	// res := recorder.Result()
	if res.StatusCode != http.StatusOK {
		test.Fatalf("User Details Not able to Delete error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestAddCar(test *testing.T) {
	body := strings.NewReader(`{
		"CarNumber" :"MH05AX4158",
		"CarModel"  :"Hyundai i10"
	}`)
	httpreq, err := http.NewRequest("POST", "http://localhost:8080/car", body)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("User Details Not able to Delete error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestUpdateCar(test *testing.T) {
	body := strings.NewReader(`{
		"CarNumber" :"MH05AX4158",
		"CarModel"  :"Hyundai i10 New"
	}`)
	httpreq, err := http.NewRequest("PUT", "http://localhost:8080/car/MH05AX4158", body)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Car Details Not able to Update error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestDeleteCar(test *testing.T) {
	httpreq, err := http.NewRequest("DELETE", "http://localhost:8080/car/MH05AX4158", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Car Details Not able to Delete error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestAddSlot(test *testing.T) {
	body := strings.NewReader(`{
		"FloorNumber" : 21,
		"UniqueSlotNumber" : 27,
		"Occupancy" :  false
	}`)
	httpreq, err := http.NewRequest("POST", "http://localhost:8080/slot", body)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Car Details Not Added error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestDeleteSlot(test *testing.T) {
	httpreq, err := http.NewRequest("DELETE", "http://localhost:8080/slot/3", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Slot Details Not able to Delete error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}
func TestUpdateSlot(test *testing.T) {
	body := strings.NewReader(`{
		"FloorNumber" : 111,
		"UniqueSlotNumber" : 111,
		"Occupancy" :  false
	}`)
	httpreq, err := http.NewRequest("PUT", "http://localhost:8080/slot/1", body)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Slot Details Not able to Update error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}
func TestGetSlot(test *testing.T) {
	// CarParking.GetFreeParkingSlots(recorder, req)
	// res := recorder.Result()
	httpreq, _ := http.NewRequest("GET", "http://localhost:8080/slot", nil)
	client := &http.Client{}
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Expected Status OK but got %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestAddNewCarToSlot(test *testing.T) {
	httpreq, err := http.NewRequest("PUT", "http://localhost:8080/cartoslot/2/MH05AX4158/1", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Slot Details Not able to Update error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}

func TestRemoveCarFromSlot(test *testing.T) {
	httpreq, err := http.NewRequest("DELETE", "http://localhost:8080/cartoslot/1", nil)
	res, err := client.Do(httpreq)
	if res.StatusCode != http.StatusOK {
		test.Fatalf("Slot Details Not able to Delete error: %v", res.StatusCode)
	}
	if err != nil {
		test.Fatalf("%v", err)
	}
}
