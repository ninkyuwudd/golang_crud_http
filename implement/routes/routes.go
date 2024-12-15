package routes

import (
	"learn_http_interface/implement/registry"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	registed := registry.NewAppRegistry()
	registedVehicle := registed.VehicleHelper
	registeredBike := registed.BikeHelper
	mux.HandleFunc("/vehicle", registedVehicle.GetAllVehicle)
	mux.HandleFunc("/vehicle/add", registedVehicle.AddNewVehicles)
	mux.HandleFunc("/vehicle/id", registedVehicle.GetVehiclesById)
	mux.HandleFunc("/vehicle/delete", registedVehicle.DeleteVehiclesById)
	mux.HandleFunc("/vehicle/update", registedVehicle.UpdateVehicle)
	mux.HandleFunc("/bike", registeredBike.GetAllBike)
	mux.HandleFunc("/bike/id", registeredBike.GetBikeById)
	mux.HandleFunc("/bike/add", registeredBike.AddNewBike)

	return mux
}
