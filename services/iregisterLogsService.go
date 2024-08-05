package services

import "template_soa/models"

type IRegisterLogsService interface {
	CreateLog(tripData *models.RequestTrip) error
	UpdateLog(status int, tripId string) error
}
