package models

import (
	"strconv"
	"time"
)

type BookingModel struct {
	Id               int       `json:"id" db:"id"`
	UserId           int       `json:"user_id" db:"user_id"`
	ScheduledClassId int       `json:"scheduled_class_id" db:"scheduled_class_id"`
	BookingTime      time.Time `json:"booking_time" db:"booking_time"`
	Status           int       `json:"status" db:"status"`
}

func ToProto(bookingModel *BookingModel) *Booking {
	return &Booking{
		Id:               strconv.Itoa(bookingModel.Id),
		UserId:           strconv.Itoa(bookingModel.UserId),
		ScheduledClassId: strconv.Itoa(bookingModel.ScheduledClassId),
		BookingTime:      time.Time.Format(bookingModel.BookingTime, "2006-01-02 15:04:05"),
		Status:           strconv.Itoa(bookingModel.Status),
	}
}
