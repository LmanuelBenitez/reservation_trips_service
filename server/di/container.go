package di

import (
	"template_soa/config"
	"template_soa/controllers"
	"template_soa/repositories"
	"template_soa/services"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.NewConfiguration)
	_ = container.Provide(controllers.NewReservationTripController)
	_ = container.Provide(services.NewReservationService)
	_ = container.Provide(services.NewRegisterLogsService)
	_ = container.Provide(repositories.NewReservationRepository)
	_ = container.Provide(repositories.NewRegisterLogsRepository)

	return container
}
