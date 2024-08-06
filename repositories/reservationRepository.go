package repositories

import (
	"fmt"
	"net/http"
	"template_soa/db"
	"template_soa/models"
)

type ReservationRepository struct {
	Database   *db.TripsDatabase
	HttpClient http.Client
}

func NewReservationRepository(database *db.TripsDatabase) *ReservationRepository {
	return &ReservationRepository{
		Database:   database,
		HttpClient: http.Client{},
	}
}

func (repository *ReservationRepository) MakeReservation(dataTrip *models.RequestTrip) (*models.Response, error) {
	driver := new(models.Driver)

	result := repository.Database.Connection.Where("id = ?", dataTrip.DriverId).First(&driver)
	if result.Error != nil {
		return &models.Response{
			Message:    "Bad Request",
			Details:    "Driver does not exist in the reservation system",
			Status:     http.StatusBadRequest,
			ApiVersion: "1.0",
		}, result.Error
	}

	response := models.Response{
		Message:        "Trip reserved successfully",
		Details:        fmt.Sprintf("The driver will arrived at %s", dataTrip.DateTrip.String()),
		DriverResponse: *driver,
		Status:         http.StatusOK,
		ApiVersion:     "1.0",
	}

	return &response, nil
}
