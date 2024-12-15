package vehicle

import (
	"errors"
	"learn_http_interface/implement/model"
)

type VehicleRepositoryMock interface {
	AllVehicle() ([]*model.Vehicle, error)
	AddVehicle(vehicle *model.Vehicle) *model.Vehicle
	GetVehicleById(id int) (*model.Vehicle, error)
	DeleteVehicleById(id int) error
	UpdateVehicle(vehicle *model.Vehicle) error
}

var vehiclesList = []*model.Vehicle{
	&model.Vehicle{Id: 1, Name: "car", Type: "4 wheels"},
	&model.Vehicle{Id: 2, Name: "bike", Type: "2 wheels"},
}

type vehicleRepositoryMock struct {
	vehiclesList []*model.Vehicle
}

func NewVehicleRositoryMock() VehicleRepositoryMock {
	return &vehicleRepositoryMock{
		vehiclesList: vehiclesList,
	}
}

func (repo *vehicleRepositoryMock) AllVehicle() ([]*model.Vehicle, error) {
	return repo.vehiclesList, nil
}

func (repo *vehicleRepositoryMock) AddVehicle(vehicle *model.Vehicle) *model.Vehicle {
	repo.vehiclesList = append(repo.vehiclesList, vehicle)
	return vehicle
}

func (repo *vehicleRepositoryMock) GetVehicleById(id int) (*model.Vehicle, error) {
	for _, vehicle := range repo.vehiclesList {
		if vehicle.Id == id {
			return vehicle, nil
		}
	}
	return nil, errors.New("Vehicle Not Found")
}

func (repo *vehicleRepositoryMock) DeleteVehicleById(id int) error {

	for i, vehicle := range repo.vehiclesList {
		if vehicle.Id == id {
			repo.vehiclesList[i] = repo.vehiclesList[len(repo.vehiclesList)-1]
			repo.vehiclesList = repo.vehiclesList[:len(repo.vehiclesList)-1]
			return nil
		}
	}

	return errors.New("Vehicle Not Found")

}

func (repo *vehicleRepositoryMock) UpdateVehicle(vehicleUpdate *model.Vehicle) error {
	for i, vehicle := range repo.vehiclesList {
		if vehicle.Id == vehicleUpdate.Id {
			repo.vehiclesList[i] = vehicleUpdate
			return nil
		}
	}
	return errors.New("Vehicle Not Found")
}
