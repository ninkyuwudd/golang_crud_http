package service

import (
	"learn_http_interface/implement/model"
	"learn_http_interface/implement/repository/bike"
)

type BikeService interface {
	GetBikeList() ([]*model.Bike, error)
}

type bikeService struct {
	repo bike.BikeRepository
}

func NewBikeRepository(repo bike.BikeRepository) BikeService {
	return &bikeService{repo}
}

func (s bikeService) GetBikeList() ([]*model.Bike, error) {
	return s.repo.AllBike()
}
