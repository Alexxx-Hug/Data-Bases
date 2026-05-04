package service

import (
	"context"
	"sharing-app/internal/model"
	"sharing-app/internal/repository"
	"time"
)

type PromotionService struct {
	repo *repository.PromotionRepository
}

func NewPromotionService(repo *repository.PromotionRepository) *PromotionService {
	return &PromotionService{repo: repo}
}

func (s *PromotionService) CreatePromotion(ctx context.Context, code string, discount int) error {
	return s.repo.Create(ctx, model.Promotion{
		PromoCode:       code,
		DiscountPercent: discount,
		StartDate:       time.Now(),
		EndDate:         time.Now().AddDate(0, 1, 0),
	})
}

func (s *PromotionService) AssignToUser(ctx context.Context, userID, promoID int64) error {
	return s.repo.AssignToUser(ctx, userID, promoID)
}

func (s *PromotionService) GetUserPromotions(ctx context.Context, userID int64) ([]model.UserPromotionView, error) {
	return s.repo.GetUserPromotions(ctx, userID)
}

func (s *PromotionService) ListPromotions(ctx context.Context) ([]model.Promotion, error) {
	return s.repo.GetAll(ctx)
}

func (s *PromotionService) RemoveFromUser(ctx context.Context, userID, promoID int64) error {
	return s.repo.RemoveFromUser(ctx, userID, promoID)
}
