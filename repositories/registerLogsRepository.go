package repositories

import (
	"template_soa/db"
	"template_soa/models"
)

type RegisterLogsRepository struct {
	Database *db.TripsDatabase
}

func NewRegisterLogsRepository(database *db.TripsDatabase) *RegisterLogsRepository {
	return &RegisterLogsRepository{
		Database: database,
	}
}

func (repository *RegisterLogsRepository) CreateLogTrip(tripData *models.RequestTrip) error {

	result := repository.Database.Connection.Select(
		"Id", "PassengerName", "PassengerLastName", "NumberOfPassenger",
		"DateTrip", "StartingPlace", "FinishingPlace",
	).Create(tripData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *RegisterLogsRepository) UpdateLogTrip(status int, tripId string) error {

	result := repository.Database.Connection.Model(&models.RequestTrip{}).Where("id = ?", tripId).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
