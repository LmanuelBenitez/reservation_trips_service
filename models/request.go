package models

import "time"

type RequestTrip struct {
	Id                 string    `json:"id"`
	PassengerName      string    `json:"passengerName"`
	PassengerLastName  string    `json:"passengerLastName"`
	NumberOfPassengers int       `json:"numberOfPassengers"`
	DriverId           string    `json:"driverId"`
	DateTrip           time.Time `json:"dateTrip"`
	Status             int       `json:"status"`
	StartingPlace      string    `json:"startingPlace"`
	FinishedPlace      string    `json:"finishedPlace"`
}
