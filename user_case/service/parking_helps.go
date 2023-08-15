package service

import (
	"ProjetoEstacionamento/entity"
	"ProjetoEstacionamento/user_case/response"
	"errors"
)

func setUpParking(TotalParking, TotalLargeParking int) {
	for n := 0; n < TotalParking; n++ {
		parking[n] = &entity.ModelParking{
			ParkingSpot: false,
			VehicleType: "",
		}
	}

	for n := 0; n < TotalLargeParking; n++ {
		largeParking[n] = &entity.ModelParking{
			ParkingSpot: false,
			VehicleType: "",
		}
	}
}

func (p ParkingService) checkVehicle(vehicle string) (entity.ModelVehicle, error) {
	var data = entity.ModelVehicle{}
	switch vehicle {
	case entity.CAR:
		data = entity.ModelVehicle{
			Name:    entity.CAR,
			Parking: 1,
		}
	case entity.MOTORBIKE:
		data = entity.ModelVehicle{
			Name:    entity.MOTORBIKE,
			Parking: 1,
		}
	case entity.VAN:
		data = entity.ModelVehicle{
			Name:    entity.VAN,
			Parking: 3,
		}
	default:
		return data, errors.New("wrong vehicle type")
	}

	return data, nil
}

func (p ParkingService) countVehicleLargeParking() response.InfoLargeParking {
	var countVanLargeParking int
	var countLargeParkingOccupied int
	var countLargeParkingFree int
	for _, item := range largeParking {
		if item.VehicleType == entity.VAN {
			countVanLargeParking++
		}

		if item.ParkingSpot {
			countLargeParkingOccupied++
			continue
		}
		countLargeParkingFree++
	}

	return response.InfoLargeParking{
		Free:     countLargeParkingFree,
		Occupied: countLargeParkingOccupied,
		Van:      countVanLargeParking,
	}
}
