package repository

import (
	"context"
	"sharing-app/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx,
		`SELECT * FROM users
		ORDER BY users.id`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(
			&u.ID,
			&u.Fullname,
			&u.Phone,
			&u.Birthday,
			&u.RegistrationDate,
			&u.Lon,
			&u.Lat,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) Create(ctx context.Context, user model.User) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO users (fio, phone, birthday, lon, lat)
		VALUES ($1, $2, $3, $4, $5)`,
		user.Fullname,
		user.Phone,
		user.Birthday,
		user.Lon,
		user.Lat)

	return err
}

func (r *UserRepository) Update(ctx context.Context, user model.User) error {
	_, err := r.db.Exec(ctx,
		`UPDATE users
		SET fio = $1,
			phone = $2,
			birthday = $3,
			lon = $4,
			lat = $5
		WHERE id = $6 `,
		user.Fullname,
		user.Phone,
		user.Birthday,
		user.Lon,
		user.Lat,
		user.ID)

	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id)

	return err

}

func (r *UserRepository) GetTripStats(ctx context.Context) ([]model.UserTripStats, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			u.fio,
			COUNT(t.id) AS trip_count
		FROM users u
		LEFT JOIN trips t ON t.user_id = u.id
		GROUP BY u.id, u.fio
		ORDER BY trip_count
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.UserTripStats

	for rows.Next() {
		var s model.UserTripStats

		if err := rows.Scan(&s.UserName, &s.TripCount); err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
