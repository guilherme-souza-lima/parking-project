package service

import (
	"ProjetoEstacionamento/entity"
	"ProjetoEstacionamento/user_case/response"
	"errors"
)

var m = make(map[int]*entity.ModelParking)

type InterfaceParkingService interface {
	Info() response.InfoParking
	Occupy(vehicle string) error
	Release(vehicle string) error
}

type ParkingService struct {
	TotalParking int
}

func NewParkingService(TotalParking int) ParkingService {
	for n := 0; n < TotalParking; n++ {
		m[n] = &entity.ModelParking{
			ParkingSpot: false,
			VehicleType: "",
		}
	}
	return ParkingService{TotalParking}
}

func (p ParkingService) Info() response.InfoParking {
	var countFree int
	var countOccupied int
	var countCar int
	var countMotorbike int
	var countVan int
	for _, item := range m {

		if item.VehicleType == "car" {
			countCar++
		}
		if item.VehicleType == "motorbike" {
			countMotorbike++
		}
		if item.VehicleType == "van" {
			countVan++
		}

		if item.ParkingSpot {
			countOccupied++
			continue
		}
		countFree++
	}

	return response.InfoParking{
		Free:     countFree,
		Occupied: countOccupied,
		VehicleInformation: response.VehicleInformation{
			Car:       countCar,
			Motorbike: countMotorbike,
			Van:       countVan / 3,
		},
	}
}

func (p ParkingService) Occupy(vehicle string) error {
	data, err := p.checkVehicle(vehicle)
	if err != nil {
		return err
	}
	parkingControl := 0
	for n := 0; n < p.TotalParking; n++ {
		if !m[n].ParkingSpot {
			m[n].ParkingSpot = true
			m[n].VehicleType = data.Name
			parkingControl++
			if parkingControl >= data.Parking {
				break
			}
		}
	}
	return nil
}

func (p ParkingService) Release(vehicle string) error {
	data, err := p.checkVehicle(vehicle)
	if err != nil {
		return err
	}
	parkingControl := 0
	for n := 0; n < p.TotalParking; n++ {
		if m[n].ParkingSpot && m[n].VehicleType == data.Name {
			m[n].ParkingSpot = false
			m[n].VehicleType = ""
			parkingControl++
			if parkingControl >= data.Parking {
				break
			}
		}
	}
	return nil
}

func (p ParkingService) checkVehicle(vehicle string) (entity.ModelVehicle, error) {
	var data = entity.ModelVehicle{}
	switch vehicle {
	case "car":
		data = entity.ModelVehicle{
			Name:    "car",
			Parking: 1,
		}
	case "motorbike":
		data = entity.ModelVehicle{
			Name:    "motorbike",
			Parking: 1,
		}
	case "van":
		data = entity.ModelVehicle{
			Name:    "van",
			Parking: 3,
		}
	default:
		return data, errors.New("wrong vehicle type")
	}

	return data, nil
}
