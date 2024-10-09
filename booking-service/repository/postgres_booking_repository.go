package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ddarli/gym/bookingservice/models"
	mylogger "github.com/Ddarli/gym/common/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type postgresBookingRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewPostgresBookingRepository(db *sqlx.DB) BookingRepository {
	return &postgresBookingRepository{
		db:     db,
		logger: mylogger.GetLogger(),
	}
}

func (r *postgresBookingRepository) Create(newBooking *models.BookingModel) (int, error) {
	var id int
	tx := r.db.MustBegin()
	tx.QueryRow("INSERT INTO bookings (user_id, scheduled_class_id, booking_time, status) VALUES ($1, $2, $3, $4) RETURNING id",
		newBooking.UserId, newBooking.ScheduledClassId, newBooking.BookingTime, newBooking.Status).Scan(&id)
	if err := tx.Commit(); err != nil {
		r.logger.Errorf("Failed to commit transaction to booking service: %v", err)
		return 0, err
	}
	r.logger.Infof("Created booking with id: %d", id)
	return id, nil
}
func (r *postgresBookingRepository) Get(id int) (*models.BookingModel, error) {

	var booking models.BookingModel
	err := r.db.Get(&booking, "SELECT * FROM bookings WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Errorf("Failed to get booking with ID: %d, error: %v", id, err)
			return nil, fmt.Errorf("booking with id %d not found", id)
		}
		r.logger.Errorf("Failed to get booking with ID: %d, error: %v", id, err)
		return nil, err
	}
	r.logger.Infof("Successfully retrieved booking with ID: %d", booking.Id)
	return &booking, nil
}
func (r *postgresBookingRepository) Delete(id int) error {
	tx := r.db.MustBegin()
	tx.MustExec("DELETE FROM bookings WHERE id = $1", id)
	if err := tx.Commit(); err != nil {
		r.logger.Errorf("Failed to commit transaction to booking service: %v", err)
		return err
	}
	r.logger.Infof("Successfully deleted booking with ID: %d", id)
	return nil
}

func (r *postgresBookingRepository) Update(id int, updatedBooking *models.BookingModel) error {
	tx := r.db.MustBegin()
	tx.MustExec("UPDATE bookings SET status = $1 WHERE id = $2", updatedBooking.Status, id)
	err := tx.Commit()
	if err != nil {
		r.logger.Errorf("Failed to commit transaction to booking service: %v", err)
		return err
	}
	return nil
}
