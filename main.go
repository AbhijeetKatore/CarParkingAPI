package main

import (
	"CarParking/CarParking"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//go:generate swagger generate spec -m -o ./swagger.yaml

	Connection()
}

func Connection() {
	Router := mux.NewRouter()
	Router.HandleFunc("/user", CarParking.AddUser).Methods("POST")
	Router.HandleFunc("/user", CarParking.GetUser).Methods("GET")
	Router.HandleFunc("/user", CarParking.DeleteUser).Methods("DELETE")
	Router.HandleFunc("/car", CarParking.AddCarDetails).Methods("POST")
	Router.HandleFunc("/car", CarParking.DeleteCarDetails).Methods("DELETE")
	Router.HandleFunc("/car/{car_number}", CarParking.UpdateCarDetails).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.AddParkingSlot).Methods("POST")
	Router.HandleFunc("/slot", CarParking.DeleteParkingSlots).Methods("DELETE")
	Router.HandleFunc("/slot/{_uniqueslotid}", CarParking.UpdateParkingSlot).Methods("PUT")
	Router.HandleFunc("/slot", CarParking.GetFreeParkingSlots).Methods("GET")
	Router.HandleFunc("/cartoslot/{_userid}/{carnumber}/{_uniqueslotid}", CarParking.AddNewCarToSlot).Methods("PUT")
	Router.HandleFunc("/cartoslot", CarParking.DeleteCarFromSlot).Methods("DELETE")
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		fmt.Fprintln(w, "Application Running Succefully")
		w.WriteHeader(200) })
	Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Api Calls Working")
		w.WriteHeader(200)
	}).Methods("GET")

	http.ListenAndServe(":8080", Router)

}
