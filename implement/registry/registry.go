package registry

import (
	"learn_http_interface/implement/helper"
	"learn_http_interface/implement/repository/bike"
	"learn_http_interface/implement/repository/vehicle"
	"learn_http_interface/implement/service"
)

type AppRegistry struct {
	VehicleHelper *helper.VehicleHelper
	BikeHelper    *helper.BikeHelper
}

func NewAppRegistry() *AppRegistry {

	vehicleHelper := helper.NewVehicleHelper(service.NewVehicleService(vehicle.NewVehicleRository()))

	bikeHelper := helper.NewBikeHelper(service.NewBikeRepository(bike.NewBikeRepository()))

	return &AppRegistry{
		VehicleHelper: vehicleHelper,
		BikeHelper:    bikeHelper,
	}

}
