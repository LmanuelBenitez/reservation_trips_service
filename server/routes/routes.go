package routes

import (
	"template_soa/controllers"
	"template_soa/models"

	"go.uber.org/dig"
)

var Routes []models.Route

func SetRoutes(container *dig.Container) {
	_ = container.Invoke(func(reservationTripController *controllers.ReservationTripController) {
		Routes = append(Routes,
			models.Route{
				Method:      "POST",
				Name:        "ReservationTripController",
				Pattern:     "/",
				HandlerFunc: reservationTripController.GetReservation,
			})
	})
}
