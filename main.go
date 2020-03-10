package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Rolls are models for sushi
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

// Rolls as a slice i.e. a variable array
var rolls []Roll

func main() {

	roll1 := Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chilli, Sauce, Nori, Rice"}
	roll2 := Roll{ID: "2", ImageNumber: "6", Name: "California Roll", Ingredients: "Crab, Avocado, Cucumber, Nori, Rice"}
	rolls = append(rolls, roll1, roll2)

	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/sushi", getRolls).Methods("Get")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("Get")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func getRolls(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range rolls {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createRoll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			delete(i)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			delete(i)
			break
		}
	}
	json.NewEncoder(w).Encode(rolls)
}

func delete(i int) {
	rolls = append(rolls[:i], rolls[i+1:]...)
}
