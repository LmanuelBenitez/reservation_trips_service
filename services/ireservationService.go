package services

import "template_soa/models"

type IReservationService interface {
	ReserveTrip(tripData *models.RequestTrip) (*models.Response, error)
}
