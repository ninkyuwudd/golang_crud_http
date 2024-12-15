package helper

import (
	"encoding/json"
	"learn_http_interface/implement/model"
	"learn_http_interface/implement/service"
	"net/http"
	"strconv"
)

type BikeHelper struct {
	service service.BikeService
}

func NewBikeHelper(service service.BikeService) *BikeHelper {
	return &BikeHelper{service}
}

func (helper BikeHelper) GetAllBike(w http.ResponseWriter, r *http.Request) {
	bike, err := helper.service.GetBikeList()

	if err != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bike)

	if err != nil {
		return
	}

}

func (helper BikeHelper) AddNewBike(w http.ResponseWriter, r *http.Request) {
	var bike *model.Bike
	if err := json.NewDecoder(r.Body).Decode(&bike); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdItem := helper.service.AddBike(bike)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(createdItem)
	if err != nil {
		return
	}

}

func (helper BikeHelper) GetBikeById(w http.ResponseWriter, r *http.Request) {
	readId := r.URL.Query().Get("id")
	getId, errConv := strconv.Atoi(readId)
	if errConv != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}
	bike, err := helper.service.GetBikeById(getId)
	if err != nil {
		//http.Error(w, "data vehicle kosong", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(bike)
	if err != nil {
		return
	}
}

func (helper BikeHelper) DeleteBikeById(w http.ResponseWriter, r *http.Request) {
	readId := r.URL.Query().Get("id")
	getId, errConv := strconv.Atoi(readId)
	if errConv != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)

	}
	err := helper.service.DeleteBikeById(getId)
	if err != nil {
		http.Error(w, "data vehicle kosong", http.StatusInternalServerError)
	}
}
