package model

import "time"

type User struct{
	ID int64
	Fullname string
	Phone string
	Birthday *time.Time
	RegistrationDate time.Time
	Lon *float64
	Lat *float64
}