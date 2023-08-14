package infra

import (
	"ProjetoEstacionamento/user_case/handler"
	"ProjetoEstacionamento/user_case/service"
)

type Container struct {
	Config         Config
	ParkingHandler handler.ParkingHandler
	ParkingService service.ParkingService
}

func NewContainer(config Config) *Container {
	container := &Container{
		Config: config,
	}
	container.ParkingService = service.NewParkingService(
		container.Config.TotalParkingSpaces,
		container.Config.TotalLargeParkingSpaces,
	)
	container.ParkingHandler = handler.NewParkingHandler(
		container.ParkingService,
	)
	return container
}
