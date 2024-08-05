package repositories

import (
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

	repository.Database.Connection.Where("id = ?", dataTrip.DriverId).First(&driver)
	// if err != nil {
	// 	logger.Error("ReservationRepository", "MakeReservation", err.Error())
	// }

	response := models.Response{
		Message:    "",
		Details:    "",
		Status:     http.StatusOK,
		ApiVersion: "1.0",
	}

	return &response, nil
}
