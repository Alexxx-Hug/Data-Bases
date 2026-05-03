package model

import "time"

type TripView struct {
	ID           int64
	StartTime    time.Time
	EndTime      *time.Time
	TripStatus   string
	UserName     string
	StartParking string
	EndParking   *string
}
