package bike

import "learn_http_interface/implement/model"

var bikes = []*model.Bike{
	{Id: 0, Name: "CBR150", Brand: "Honda", Vehicle: model.Vehicle{Id: 2, Name: "bike", Type: "2 wheels"}},
	{Id: 1, Name: "RXking", Brand: "Suzuki", Vehicle: model.Vehicle{Id: 2, Name: "bike", Type: "2 wheels"}},
}

type BikeRepository interface {
	AllBike() ([]*model.Bike, error)
}

type bikeRepository struct {
}

func NewBikeRepository() BikeRepository {
	return &bikeRepository{}
}

func (repo *bikeRepository) AllBike() ([]*model.Bike, error) {
	return bikes, nil
}
