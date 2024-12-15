package service

import (
	"learn_http_interface/implement/model"
	"learn_http_interface/implement/repository/vehicle"
)

type VehicleService interface {
	GetVehicleList() ([]*model.Vehicle, error)
	PutVehicle(vehicle *model.Vehicle) *model.Vehicle
	GetVehicleById(id int) (*model.Vehicle, error)
	DeleteVehicle(id int) error
	UpdateVehicle(vehicle *model.Vehicle) error
}

type vehicleService struct {
	repo vehicle.VehicleRepository
}

func NewVehicleService(repo vehicle.VehicleRepository) VehicleService {
	return &vehicleService{repo}
}

func (v vehicleService) GetVehicleList() ([]*model.Vehicle, error) {
	return v.repo.AllVehicle()
}

func (v vehicleService) PutVehicle(data *model.Vehicle) *model.Vehicle {
	return v.repo.AddVehicle(data)
}

func (v vehicleService) GetVehicleById(id int) (*model.Vehicle, error) {

	return v.repo.GetVehicleById(id)
}

func (v vehicleService) DeleteVehicle(id int) error {
	return v.repo.DeleteVehicleById(id)
}

func (v vehicleService) UpdateVehicle(data *model.Vehicle) error {
	return v.repo.UpdateVehicle(data)
}
