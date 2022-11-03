package main

import (
	"CarParking/CarParking"
	"net/http"

	"github.com/gorilla/mux"
)
func main(){
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



	http.ListenAndServe(":8080", Router)
}
