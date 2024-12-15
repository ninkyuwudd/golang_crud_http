package bike

import "learn_http_interface/implement/model"

var bikes = []*model.Bike{
	{Id: 0, Name: "CBR150", Brand: "Honda", Vehicle: model.Vehicle{Id: 2, Name: "bike", Type: "2 wheels"}},
	{Id: 1, Name: "RXking", Brand: "Suzuki", Vehicle: model.Vehicle{Id: 2, Name: "bike", Type: "2 wheels"}},
}

type RepositoryBike interface {
	AllBike() ([]*model.Bike, error)
	GetBikeById(id int) (*model.Bike, error)
	AddBike(bike *model.Bike) *model.Bike
	DeleteBikeById(id int) error
	UpdateBike(bike *model.Bike) error
}

type repositoryBike struct {
}

func NewBikeRepository() RepositoryBike {
	return &repositoryBike{}
}

func (repo *repositoryBike) AllBike() ([]*model.Bike, error) {
	return bikes, nil
}

func (repo *repositoryBike) GetBikeById(id int) (*model.Bike, error) {
	for _, bike := range bikes {
		if bike.Id == id {
			return bike, nil
		}
	}
	return nil, nil
}

func (repo *repositoryBike) AddBike(bike *model.Bike) *model.Bike {
	bikes = append(bikes, bike)
	return bike
}

func (repo *repositoryBike) DeleteBikeById(id int) error {
	for i, bike := range bikes {
		if bike.Id == id {
			bikes[i] = bikes[len(bikes)-1]
			bikes = bikes[:len(bikes)-1]
			return nil
		}
	}
	return nil
}

func (repo *repositoryBike) UpdateBike(bike *model.Bike) error {
	for i, b := range bikes {
		if b.Id == bike.Id {
			bikes[i] = bike
			return nil
		}
	}
	return nil
}
