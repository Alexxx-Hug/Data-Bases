package model

import "time"

type Promotion struct {
	ID              int64
	PromoCode       string
	DiscountPercent int
	StartDate       time.Time
	EndDate         time.Time
	IsActive        bool
}

type UserPromotionView struct {
	UserName  string
	PromoCode string
	Discount  int
}

