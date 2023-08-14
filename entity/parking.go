package entity

const (
	CAR       = "car"
	MOTORBIKE = "motorbike"
	VAN       = "van"
)

type ModelParking struct {
	ParkingSpot bool
	VehicleType string
}

type ModelVehicle struct {
	Name    string
	Parking int
}
