package handler

import (
	"ProjetoEstacionamento/user_case/request"
	"ProjetoEstacionamento/user_case/service"
	"github.com/gofiber/fiber/v2"
)

type ParkingHandler struct {
	InterfaceParkingService service.InterfaceParkingService
}

func NewParkingHandler(InterfaceParkingService service.InterfaceParkingService) ParkingHandler {
	return ParkingHandler{InterfaceParkingService}
}

func (e ParkingHandler) Info(c *fiber.Ctx) error {
	result := e.InterfaceParkingService.Info()
	return c.Status(fiber.StatusOK).JSON(result)
}

func (e ParkingHandler) Occupy(c *fiber.Ctx) error {
	var vehicle request.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	err := e.InterfaceParkingService.Occupy(vehicle.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("success")
}

func (e ParkingHandler) Release(c *fiber.Ctx) error {
	var vehicle request.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	err := e.InterfaceParkingService.Release(vehicle.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("success")
}
