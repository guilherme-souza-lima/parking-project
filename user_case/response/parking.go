package response

type InfoParking struct {
	Free               int                `json:"free_parking_spaces"`
	Occupied           int                `json:"occupied_parking_spaces"`
	VehicleInformation VehicleInformation `json:"vehicle_information"`
}

type VehicleInformation struct {
	Car       int `json:"car"`
	Motorbike int `json:"motorbike"`
	Van       int `json:"van"`
}
