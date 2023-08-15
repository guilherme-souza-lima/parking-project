package response

type Info struct {
	InfoLargeParking InfoLargeParking `json:"info_large_parking"`
	InfoParking      InfoParking      `json:"info_parking"`
}

type VehicleInformation struct {
	Car       int `json:"car"`
	Motorbike int `json:"motorbike"`
	Van       int `json:"van"`
}

type InfoParking struct {
	Free               int                `json:"free_parking_spaces"`
	Occupied           int                `json:"occupied_parking_spaces"`
	VehicleInformation VehicleInformation `json:"vehicle_information"`
}

type InfoLargeParking struct {
	Free     int `json:"free_large_parking_spaces"`
	Occupied int `json:"occupied_large_parking_spaces"`
	Van      int `json:"vehicle_van"`
}
