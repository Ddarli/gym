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

func ToBookingModel(model *Booking) BookingModel {
	formattedBookingTime := model.BookingTime.AsTime()
	id, _ := strconv.Atoi(model.Id)
	userId, _ := strconv.Atoi(model.UserId)
	scheduledClassId, _ := strconv.Atoi(model.ScheduledClassId)
	return BookingModel{
		Id:               id,
		UserId:           userId,
		ScheduledClassId: scheduledClassId,
		BookingTime:      formattedBookingTime,
		Status:           0,
	}
}
