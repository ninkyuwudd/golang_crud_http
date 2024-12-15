package vehicle

import (
	"database/sql"
	"learn_http_interface/implement/model"
	"learn_http_interface/utils"
)

type VehicleRepository interface {
	AllVehicle() ([]*model.Vehicle, error)
	AddVehicle(vehicle *model.Vehicle) *model.Vehicle
	GetVehicleById(id int) (*model.Vehicle, error)
	DeleteVehicleById(id int) error
	UpdateVehicle(vehicle *model.Vehicle) error
}

type vehicleRepository struct {
	conn *sql.DB
}

func NewVehicleRository() VehicleRepository {
	return &vehicleRepository{
		conn: utils.MysqlConnection(),
	}
}

func (repo *vehicleRepository) AllVehicle() ([]*model.Vehicle, error) {

	db, err := repo.conn.Query("select * from vehicle")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var vehicles []*model.Vehicle
	for db.Next() {
		var vehicle model.Vehicle
		db.Scan(&vehicle.Id, &vehicle.Name, &vehicle.Type)
		vehicles = append(vehicles, &vehicle)
	}

	return vehicles, nil
}

func (repo *vehicleRepository) AddVehicle(vehicle *model.Vehicle) *model.Vehicle {

	db, err := repo.conn.Exec("insert into vehicle(name, type) values(?, ?)", vehicle.Name, vehicle.Type)

	if err != nil {
		return nil
	}

	id, err := db.LastInsertId()
	if err != nil {
		return nil
	}
	vehicle.Id = int(id)

	return vehicle
}

func (repo *vehicleRepository) GetVehicleById(id int) (*model.Vehicle, error) {
	vehicle := &model.Vehicle{}

	err := repo.conn.QueryRow("select * from vehicle where id = ?", id).Scan(&vehicle.Id, &vehicle.Name, &vehicle.Type)

	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (repo *vehicleRepository) DeleteVehicleById(id int) error {

	_, err := repo.conn.Exec("delete from vehicle where id = ?", id)

	return err
}

func (repo *vehicleRepository) UpdateVehicle(vehicleData *model.Vehicle) error {
	query := "UPDATE `vehicle` SET `name`=?,`type`=? WHERE id = ?"
	_, err := repo.conn.Exec(query, vehicleData.Name, vehicleData.Type, vehicleData.Id)
	return err
}
