package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ddarli/gym/bookingservice/models"
	mylogger "github.com/Ddarli/gym/common/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func toProto(model models.BookingModel) *models.Booking {
	bookingTimeProto := timestamppb.New(model.BookingTime)
	return &models.Booking{
		Id:               strconv.Itoa(model.Id),
		UserId:           strconv.Itoa(model.UserId),
		ScheduledClassId: strconv.Itoa(model.ScheduledClassId),
		BookingTime:      bookingTimeProto,
		Status:           0,
	}
}

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

func (r *postgresBookingRepository) Create(newBooking *models.Booking) error {
	booking := models.ToBookingModel(newBooking)
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO bookings (user_id, scheduled_class_id, booking_time, status) VALUES ($1, $2, $3, $4)",
		booking.UserId, booking.ScheduledClassId, booking.BookingTime, booking.Status)
	if err := tx.Commit(); err != nil {
		r.logger.Errorf("Failed to commit transaction to booking service: %v", err)
		return err
	}
	r.logger.Infof("Created booking with id: %d", booking.Id)
	return nil
}
func (r *postgresBookingRepository) Get(bookingId string) (*models.Booking, error) {
	var booking models.BookingModel
	id, err := strconv.Atoi(bookingId)
	if err != nil {
		return nil, err
	}
	err = r.db.Get(&booking, "SELECT * FROM bookings WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Errorf("Failed to get booking with ID: %d, error: %v", id, err)
			return nil, fmt.Errorf("booking with id %d not found", id)
		}
		r.logger.Errorf("Failed to get booking with ID: %d, error: %v", id, err)
		return nil, err
	}
	r.logger.Infof("Successfully retrieved booking with ID: %d", booking.Id)
	return toProto(booking), nil
}
func (r *postgresBookingRepository) Delete(id string) error {
	bookingId, err := strconv.Atoi(id)
	tx := r.db.MustBegin()
	tx.MustExec("DELETE FROM bookings WHERE id = $1", bookingId)
	if err = tx.Commit(); err != nil {
		r.logger.Errorf("Failed to commit transaction to booking service: %v", err)
		return err
	}
	r.logger.Infof("Successfully deleted booking with ID: %d", bookingId)
	return nil
}
