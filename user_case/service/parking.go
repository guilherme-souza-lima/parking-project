package service

import (
	"ProjetoEstacionamento/entity"
	"ProjetoEstacionamento/user_case/response"
)

var parking = make(map[int]*entity.ModelParking)
var largeParking = make(map[int]*entity.ModelParking)

type InterfaceParkingService interface {
	Info() response.InfoParking
	Occupy(vehicle string) (string, error)
	Release(vehicle string) (string, error)
}

type ParkingService struct {
	TotalParking      int
	TotalLargeParking int
}

func NewParkingService(TotalParking, TotalLargeParking int) ParkingService {
	setUpParking(TotalParking, TotalLargeParking)
	return ParkingService{TotalParking, TotalLargeParking}
}

func (p ParkingService) Info() response.InfoParking {
	var countFree int
	var countOccupied int
	var countCar int
	var countMotorbike int
	var countVan int

	resultLargeParking := p.countVehicleLargeParking()

	for _, item := range parking {
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
		InfoLargeParking: response.InfoLargeParking{
			Free:     resultLargeParking.Free,
			Occupied: resultLargeParking.Occupied,
			Van:      resultLargeParking.Van,
		},
		Free:     countFree,
		Occupied: countOccupied,
		VehicleInformation: response.VehicleInformation{
			Car:       countCar,
			Motorbike: countMotorbike,
			Van:       countVan / 3,
		},
	}
}

func (p ParkingService) Occupy(vehicle string) (string, error) {
	data, err := p.checkVehicle(vehicle)
	if err != nil {
		return "", err
	}

	if data.Name == entity.VAN {
		for n := 0; n < p.TotalLargeParking; n++ {
			if !largeParking[n].ParkingSpot {
				largeParking[n].ParkingSpot = true
				largeParking[n].VehicleType = data.Name
				return "success", nil
			}
		}
	}

	parkingControl := 0
	for n := 0; n < p.TotalParking; n++ {
		if !parking[n].ParkingSpot {
			parking[n].ParkingSpot = true
			parking[n].VehicleType = data.Name
			parkingControl++
			if parkingControl >= data.Parking {
				return "success", nil
			}
		}
	}
	return "all parking are occupied", nil
}

func (p ParkingService) Release(vehicle string) (string, error) {
	data, err := p.checkVehicle(vehicle)
	if err != nil {
		return "", err
	}
	parkingControl := 0
	for n := 0; n < p.TotalParking; n++ {
		if parking[n].ParkingSpot && parking[n].VehicleType == data.Name {
			parking[n].ParkingSpot = false
			parking[n].VehicleType = ""
			parkingControl++
			if parkingControl >= data.Parking {
				return "success", nil
			}
		}
	}

	if data.Name == entity.VAN {
		for n := 0; n < p.TotalLargeParking; n++ {
			if largeParking[n].ParkingSpot && largeParking[n].VehicleType == data.Name {
				largeParking[n].ParkingSpot = false
				largeParking[n].VehicleType = ""
				return "success", nil
			}
		}
	}

	return "all parking are free", nil
}
