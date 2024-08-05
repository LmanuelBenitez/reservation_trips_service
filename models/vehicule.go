package models

type Vehicle struct {
	Id                  string `json:"id"`
	VehicleRegistration string `json:"vehicleRegistration"`
	VehicleCapacity     int    `json:"vehicleCapacity"`
	Brand               string `json:"brand"`
	Color               string `json:"color"`
}
