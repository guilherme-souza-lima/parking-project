package service

import (
	"ProjetoEstacionamento/entity"
	"ProjetoEstacionamento/user_case/response"
)

var parking = make(map[int]*entity.ModelParking)
var largeParking = make(map[int]*entity.ModelParking)

type InterfaceParkingService interface {
	Info() response.Info
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

func (p ParkingService) Info() response.Info {
	resultLargeParking := p.countVehicleLargeParking()
	resultParking := p.countVehicleParking()
	return response.Info{
		InfoLargeParking: response.InfoLargeParking{
			Free:     resultLargeParking.Free,
			Occupied: resultLargeParking.Occupied,
			Van:      resultLargeParking.Van,
		},
		InfoParking: response.InfoParking{
			Free:     resultParking.Free,
			Occupied: resultParking.Occupied,
			VehicleInformation: response.VehicleInformation{
				Car:       resultParking.VehicleInformation.Car,
				Motorbike: resultParking.VehicleInformation.Motorbike,
				Van:       resultParking.VehicleInformation.Van,
			},
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
