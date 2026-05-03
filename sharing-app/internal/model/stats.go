package model

type ParkingStats struct {
	Address    string
	TripsCount int64
}

type TripStatusStats struct {
	Status string
	Count  int64
}
