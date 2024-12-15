package model

type Bike struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Brand   string  `json:"brand"`
	Vehicle Vehicle `json:"vehicle"`
}
