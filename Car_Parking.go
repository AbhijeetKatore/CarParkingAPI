package main

import (
	"CarParking/CarParking"
	"net/http"

	"github.com/gorilla/mux"
)
var errConnection error
func main(){
	 Connection()
}

func Connection()error{
	Router := mux.NewRouter()
	Router.HandleFunc("/user", CarParking.AddUser).Methods("POST")
	Router.HandleFunc("/user", CarParking.GetUser).Methods("GET")
	Router.HandleFunc("/user/{_userid}", CarParking.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/car", CarParking.AddCarDetails).Methods("POST")
	Router.HandleFunc("/car/{car_number}", CarParking.DeleteCarDetails).Methods("DELETE")
	Router.HandleFunc("/car/{car_number}", CarParking.UpdateCarDetails).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.AddParkingSlot).Methods("POST")
	Router.HandleFunc("/slot/{_uniqueslotid}", CarParking.DeleteParkingSlots).Methods("DELETE")
	Router.HandleFunc("/slot/{_uniqueslotid}", CarParking.UpdateParkingSlot).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.GetFreeParkingSlots).Methods("GET")
	Router.HandleFunc("/cartoslot/{_userid}/{carnumber}/{_uniqueslotid}", CarParking.AddNewCarToSlot).Methods("PUT")
	Router.HandleFunc("/cartoslot/{_uniqueslotid}",CarParking.DeleteCarFromSlot).Methods("DELETE")

	errConnection:=http.ListenAndServe(":8080", Router)
	return errConnection
}
