package services

import (
	"template_soa/models"
	"template_soa/repositories"
)

type RegisterLogsService struct {
	RegisterLogsRepository repositories.IRegisterLogsRepository
}

func NewRegisterLogsService(registerLogsRepository *repositories.RegisterLogsRepository) *RegisterLogsService {
	return &RegisterLogsService{
		RegisterLogsRepository: registerLogsRepository,
	}
}

func (service *RegisterLogsService) CreateLog(tripData *models.RequestTrip) error {

	err := service.RegisterLogsRepository.CreateLogTrip(tripData)
	if err != nil {
		return err
	}

	return nil
}

func (service *RegisterLogsService) UpdateLog(status int, tripId string) error {

	err := service.RegisterLogsRepository.UpdateLogTrip(status, tripId)
	if err != nil {
		return err
	}

	return nil
}
