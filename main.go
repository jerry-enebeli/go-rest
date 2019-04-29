package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port     = ":8080"
	carStore = make([]car, 0)
)

type car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Color string `json:"color"`
}

func responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", createCarHandler).Methods("POST")
	router.HandleFunc("/", getCarsHandler).Methods("GET")
	router.HandleFunc("/{carID}", getCarHandler).Methods("GET")
	router.HandleFunc("/{carID}", updateCarHandler).Methods("PATCH")
	router.HandleFunc("/{carID}", deleteCarHandler).Methods("DELETE")

	http.ListenAndServe(port, router)
}

func createCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar car

	json.NewDecoder(r.Body).Decode(&newCar)

	carStore = append(carStore, newCar)

	responseJSON(w, http.StatusCreated, newCar)
}

func getCarHandler(w http.ResponseWriter, r *http.Request) {

	var carData car

	cardID := mux.Vars(r)["carID"]

	//find car id in the carstore
	for _, c := range carStore {

		if c.ID == cardID {
			fmt.Println(c)
			carData = c
			break
		}
	}

	//check if car not found
	if (car{}) == carData {
		responseJSON(w, http.StatusNotFound, carData)
	} else {
		responseJSON(w, http.StatusOK, carData)
	}

}

func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, http.StatusOK, carStore)
}

func updateCarHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	cardID := mux.Vars(r)["carID"]

	//find car id in the carstore
	for i, c := range carStore {

		if c.ID == cardID {
			carStore = append(carStore[:i], carStore[i+1:]...)
			break
		}
	}

	responseJSON(w, http.StatusOK, "data deleted")
}
