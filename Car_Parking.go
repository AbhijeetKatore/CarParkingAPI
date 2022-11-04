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
	Router.HandleFunc("/user/{_id}", CarParking.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/car", CarParking.AddCarDetails).Methods("POST")
	Router.HandleFunc("/car/{_id}", CarParking.DeleteCarDetails).Methods("DELETE")
	Router.HandleFunc("/car/{_id}", CarParking.UpdateCarDetails).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.AddParkingSlot).Methods("POST")
	Router.HandleFunc("/slot/{_id}", CarParking.DeleteParkingSlots).Methods("DELETE")
	Router.HandleFunc("/slot/{_id}", CarParking.UpdateParkingSlot).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.GetFreeParkingSlots).Methods("GET")
	Router.HandleFunc("/cartoslot/{user_id}/{car_id}/{slot_id}", CarParking.AddNewCarToSlot).Methods("PUT")
	Router.HandleFunc("/cartoslot/{slot_id}",CarParking.DeleteCarFromSlot).Methods("GET")

	errConnection:=http.ListenAndServe(":8080", Router)
	return errConnection
}
