package repository

import (
	"context"
	"sharing-app/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PromotionRepository struct {
	db *pgxpool.Pool
}

func NewPromotionRepository(db *pgxpool.Pool) *PromotionRepository {
	return &PromotionRepository{db: db}
}

func (r *PromotionRepository) Create(ctx context.Context, p model.Promotion) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO promotions (promo_code, discount_percent, start_date, end_date)
		VALUES ($1, $2, $3, $4)
	`,
		p.PromoCode,
		p.DiscountPercent,
		p.StartDate,
		p.EndDate,
	)
	return err
}

func (r *PromotionRepository) GetAll(ctx context.Context) ([]model.Promotion, error) {
	rows, err := r.db.Query(ctx, `SELECT id, promo_code, discount_percent, start_date, end_date, is_active FROM promotions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Promotion

	for rows.Next() {
		var p model.Promotion
		err := rows.Scan(&p.ID, &p.PromoCode, &p.DiscountPercent, &p.StartDate, &p.EndDate, &p.IsActive)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}

func (r *PromotionRepository) AssignToUser(ctx context.Context, userID, promoID int64) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO user_promotions (user_id, promotion_id)
		VALUES ($1, $2)
	`, userID, promoID)
	return err
}

func (r *PromotionRepository) RemoveFromUser(ctx context.Context, userID, promoID int64) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM user_promotions
		WHERE user_id = $1 AND promotion_id = $2
	`, userID, promoID)
	return err
}

func (r *PromotionRepository) GetUserPromotions(ctx context.Context, userID int64) ([]model.UserPromotionView, error) {
	rows, err := r.db.Query(ctx, `
		SELECT u.fio, p.promo_code, p.discount_percent
		FROM user_promotions up
		JOIN users u ON up.user_id = u.id
		JOIN promotions p ON up.promotion_id = p.id
		WHERE u.id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.UserPromotionView

	for rows.Next() {
		var v model.UserPromotionView
		err := rows.Scan(&v.UserName, &v.PromoCode, &v.Discount)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}

	return list, nil
}
