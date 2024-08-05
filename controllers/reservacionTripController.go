package controllers

import (
	"encoding/json"
	"template_soa/libs/logger"
	"template_soa/models"
	"template_soa/services"

	"github.com/labstack/echo/v4"
)

type ReservationTripController struct {
	ReservationService services.IReservationService
}

func NewReservationTripController(reservationService *services.ReservationService) *ReservationTripController {
	return &ReservationTripController{
		ReservationService: reservationService,
	}
}

func (controller *ReservationTripController) GetReservation(ctx echo.Context) error {
	requestBody := ctx.Get("body").(string)

	requestTrip := new(models.RequestTrip)
	err := json.Unmarshal([]byte(requestBody), &requestTrip)
	if err != nil {
		logger.Error("ReservationTripController", "GetReservation", err.Error())
	}

	response, err := controller.ReservationService.ReserveTrip(requestTrip)
	if err != nil {
		logger.Error("ReservationTripController", "GetReservation", err.Error())
	}

	return ctx.JSON(response.Status, response.Message)
}
