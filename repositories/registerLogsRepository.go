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
	repository.Database.Connection.Create(tripData)
	return nil
}

func (repository *RegisterLogsRepository) UpdateLogTrip(status int, tripId string) error {
	repository.Database.Connection.Model(&models.RequestTrip{}).Where("id = ?", tripId).Update("status", status)
	return nil
}
