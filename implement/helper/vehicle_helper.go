package helper

import (
	"encoding/json"
	"learn_http_interface/implement/model"
	"strconv"

	"learn_http_interface/implement/service"
	"net/http"
)

type VehicleHelper struct {
	vehicleService service.VehicleService
}

func NewVehicleHelper(service service.VehicleService) *VehicleHelper {
	return &VehicleHelper{service}
}

func (helper *VehicleHelper) GetAllVehicle(w http.ResponseWriter, r *http.Request) {
	vehicle, err := helper.vehicleService.GetVehicleList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(vehicle)
	if err != nil {
		return
	}
}

func (helper *VehicleHelper) AddNewVehicles(w http.ResponseWriter, r *http.Request) {
	var vehicle *model.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdItem := helper.vehicleService.PutVehicle(vehicle)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(createdItem)
	if err != nil {
		return
	}

}

func (helper *VehicleHelper) GetVehiclesById(w http.ResponseWriter, r *http.Request) {
	readId := r.URL.Query().Get("id")
	getId, errConv := strconv.Atoi(readId)
	if errConv != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}
	vehicle, err := helper.vehicleService.GetVehicleById(getId)
	if err != nil {
		//http.Error(w, "data vehicle kosong", http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(vehicle)
	if err != nil {
		return
	}
}

func (helper *VehicleHelper) DeleteVehiclesById(w http.ResponseWriter, r *http.Request) {
	readId := r.URL.Query().Get("id")
	getId, errConv := strconv.Atoi(readId)
	if errConv != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)

	}

	err := helper.vehicleService.DeleteVehicle(getId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`data deleted successfully`))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (helper *VehicleHelper) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle *model.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)

	}

	updatedVehicle := helper.vehicleService.UpdateVehicle(vehicle)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(updatedVehicle)
	if err != nil {
		return
	}

}
