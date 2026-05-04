package service

import (
	"context"
	"errors"
	"sharing-app/internal/model"
	"sharing-app/internal/repository"
)

type TripService struct {
	repo *repository.TripRepository
}

func NewTripService(repo *repository.TripRepository) *TripService {
	return &TripService{repo: repo}
}

func (s *TripService) CreateTrip(ctx context.Context, userID int64, transportID int64, tariffID int64, startParkingID int64) error {

	if userID <= 0 || transportID <= 0 || tariffID <= 0 {
		return errors.New("invalid foreign keys")
	}

	return s.repo.Create(ctx, model.Trip{
		UserID:         userID,
		TransportID:    transportID,
		TariffID:       tariffID,
		StartParkingID: startParkingID,
		TripStatus:     "active",
	})
}

func (s *TripService) FinishTrip(ctx context.Context, id int64) error {
	return s.repo.UpdateStatus(ctx, id, "finished", nil)
}

func (s *TripService) CancelTrip(ctx context.Context, id int64, reason string) error {
	return s.repo.UpdateStatus(ctx, id, "cancelled", &reason)
}

func (s *TripService) ListTrips(ctx context.Context) ([]model.TripView, error) {
	return s.repo.GetAllWithDetails(ctx)
}

func (s *TripService) DeleteTrips(ctx context.Context, tripId int64) error {
	return s.repo.Delete(ctx, tripId)
}

func (s *TripService) GetParkingStats(ctx context.Context) ([]model.ParkingStats, error) {
	return s.repo.GetParkingStats(ctx)
}

func (s *TripService) GetTripStatusStats(ctx context.Context) ([]model.TripStatusStats, error) {
	return s.repo.GetTripStatusStats(ctx)
}
