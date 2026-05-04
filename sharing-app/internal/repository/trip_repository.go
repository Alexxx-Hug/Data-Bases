package repository

import (
	"context"
	"sharing-app/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TripRepository struct {
	db *pgxpool.Pool
}

func NewTripRepository(db *pgxpool.Pool) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) Create(ctx context.Context, trip model.Trip) error {
	query := `
		INSERT INTO trips (
			user_id,
			transport_id,
			tariff_id,
			start_parking_id,
			trip_status
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		trip.UserID,
		trip.TransportID,
		trip.TariffID,
		trip.StartParkingID,
		trip.TripStatus,
	)

	return err
}

func (r *TripRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status string,
	cancelReason *string,
) error {

	query := `
		UPDATE trips
		SET trip_status = $1,
		    cancel_reason = $2,
		    end_time = CASE
		        WHEN $1 IN ('finished','cancelled') THEN now()
		        ELSE end_time
		    END
		WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, status, cancelReason, id)
	return err
}

func (r *TripRepository) GetAllWithDetails(ctx context.Context) ([]model.TripView, error) {
	query := `
	SELECT t.id, t.start_time, t.end_time, t.trip_status, u.fio, sp.address AS start_parking, ep.address AS end_parking FROM trips t
         JOIN users u ON t.user_id = u.id
         JOIN parkings sp ON t.start_parking_id = sp.id
         LEFT JOIN parkings ep ON t.end_parking_id = ep.id
    ORDER BY u.id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []model.TripView

	for rows.Next() {
		var trip model.TripView

		err := rows.Scan(
			&trip.ID,
			&trip.StartTime,
			&trip.EndTime,
			&trip.TripStatus,
			&trip.UserName, // ← вот это пропущено
			&trip.StartParking,
			&trip.EndParking,
		)
		if err != nil {
			return nil, err
		}

		trips = append(trips, trip)
	}

	return trips, nil
}

func (r *TripRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM trips
		WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TripRepository) GetParkingStats(ctx context.Context) ([]model.ParkingStats, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			p.address,
			COUNT(t.id) AS trips_count
		FROM parkings p
		LEFT JOIN trips t ON t.start_parking_id = p.id
		GROUP BY p.id, p.address
		ORDER BY trips_count
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.ParkingStats

	for rows.Next() {
		var s model.ParkingStats
		if err := rows.Scan(&s.Address, &s.TripsCount); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return stats, rows.Err()
}

func (r *TripRepository) GetTripStatusStats(ctx context.Context) ([]model.TripStatusStats, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			trip_status,
			COUNT(*) AS count
		FROM trips
		GROUP BY trip_status
		ORDER BY count
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.TripStatusStats

	for rows.Next() {
		var s model.TripStatusStats
		if err := rows.Scan(&s.Status, &s.Count); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}

	return stats, rows.Err()
}
