package repositories

import "template_soa/models"

type IRegisterLogsRepository interface {
	CreateLogTrip(tripData *models.RequestTrip) error
	UpdateLogTrip(status int, tripId string) error
}
