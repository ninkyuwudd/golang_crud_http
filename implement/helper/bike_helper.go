package helper

import (
	"encoding/json"
	"learn_http_interface/implement/service"
	"net/http"
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
