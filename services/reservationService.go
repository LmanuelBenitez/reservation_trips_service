package services

import (
	"template_soa/libs/logger"
	"template_soa/models"
	"template_soa/repositories"
)

type ReservationService struct {
	NotifyDriverService   INotifyDriverService
	RegisterLogsService   IRegisterLogsService
	ReservationRepository repositories.IReservationRepository
}

func NewReservationService(
	notifyDriverService *NotifyDriverService,
	registerLogsService *RegisterLogsService,
	reservationRepository *repositories.ReservationRepository,
) *ReservationService {
	return &ReservationService{
		ReservationRepository: reservationRepository,
	}
}

func (service *ReservationService) ReserveTrip(tripData *models.RequestTrip) (*models.Response, error) {
	err := service.RegisterLogsService.CreateLog(tripData)
	if err != nil {
		logger.Error("ReservationService", "ReserveTrip", err.Error())
	}

	response, err := service.ReservationRepository.MakeReservation(tripData)
	if err != nil {
		logger.Error("ReservationService", "ReserveTrip", err.Error())
		return response, err
	}

	err = service.RegisterLogsService.UpdateLog(response.Status, tripData.Id)
	if err != nil {
		logger.Error("ReservationService", "ReserveTrip", err.Error())
	}

	err = service.NotifyDriverService.NotifyDriver()
	if err != nil {
		logger.Error("ReservationService", "NotifyDriver", err.Error())
	}

	return response, nil
}
