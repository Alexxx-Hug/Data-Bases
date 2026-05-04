package model

import "time"

type Trip struct {
	ID             int64
	UserID         int64
	TransportID    int64
	TariffID       int64
	StartParkingID int64
	EndParkingID   *int64
	StartTime      time.Time
	EndTime        *time.Time
	TripStatus     string
	CancelReason   *string
}
