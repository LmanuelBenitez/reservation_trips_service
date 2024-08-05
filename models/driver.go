package models

type Driver struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Telephone string `json:"telephone"`
	VehicleId string `json:"vehiculeId"`
}
