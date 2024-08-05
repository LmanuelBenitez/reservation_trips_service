package repositories

import "template_soa/models"

type IReservationRepository interface {
	MakeReservation(tripData *models.RequestTrip) (*models.Response, error)
}
