package routes

import (
	"learn_http_interface/implement/registry"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	registed := registry.NewAppRegistry()

	mux.HandleFunc("/vehicle", registed.VehicleHelper.GetAllVehicle)
	mux.HandleFunc("/vehicle/add", registed.VehicleHelper.AddNewVehicles)
	mux.HandleFunc("/vehicle/id", registed.VehicleHelper.GetVehiclesById)
	mux.HandleFunc("/vehicle/delete", registed.VehicleHelper.DeleteVehiclesById)
	mux.HandleFunc("/vehicle/update", registed.VehicleHelper.UpdateVehicle)
	mux.HandleFunc("/bike", registed.BikeHelper.GetAllBike)

	return mux
}
