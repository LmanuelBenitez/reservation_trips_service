package models

type Response struct {
	Message        string
	Details        string
	DriverResponse Driver
	Status         int
	ApiVersion     string
}
