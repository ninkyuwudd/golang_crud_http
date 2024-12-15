package service

import (
	"learn_http_interface/implement/model"
	"learn_http_interface/implement/repository/bike"
)

type BikeService interface {
	GetBikeList() ([]*model.Bike, error)
	AddBike(bike *model.Bike) *model.Bike
	GetBikeById(id int) (*model.Bike, error)
	DeleteBikeById(id int) error
	UpdateBike(bike *model.Bike) error
}

type bikeService struct {
	repo bike.RepositoryBike
}

func NewBikeRepository(repo bike.RepositoryBike) BikeService {
	return &bikeService{repo}
}

func (s bikeService) GetBikeList() ([]*model.Bike, error) {
	return s.repo.AllBike()
}

func (s bikeService) AddBike(data *model.Bike) *model.Bike {
	return s.repo.AddBike(data)
}

func (s bikeService) GetBikeById(id int) (*model.Bike, error) {
	return s.repo.GetBikeById(id)
}

func (s bikeService) DeleteBikeById(id int) error {
	return s.repo.DeleteBikeById(id)
}

func (s bikeService) UpdateBike(data *model.Bike) error {
	return s.repo.UpdateBike(data)
}
