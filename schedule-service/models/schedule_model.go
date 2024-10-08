package models

import (
	"strconv"
	"time"
)

type ScheduleModel struct {
	Id        int       `json:"id" db:"id"`
	ClassId   int       `json:"class_id" db:"class_id"`
	TrainerId int       `json:"trainer_id" db:"trainer_id"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
}

func ToProto(model *ScheduleModel) *Schedule {
	return &Schedule{
		Id:        strconv.Itoa(model.Id),
		ClassId:   strconv.Itoa(model.ClassId),
		TrainerId: strconv.Itoa(model.TrainerId),
		StartTime: time.Time.Format(model.StartTime, "2006-01-02 15:04:05"),
		EndTime:   time.Time.Format(model.EndTime, "2006-01-02 15:04:05"),
	}
}

func ToModel(request *CreateScheduleRequest) *ScheduleModel {
	classId, _ := strconv.Atoi(request.ClassId)
	trainerId, _ := strconv.Atoi(request.TrainerId)
	startTime, _ := time.Parse("2006-01-02 15:04:05", request.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", request.EndTime)
	return &ScheduleModel{
		ClassId:   classId,
		TrainerId: trainerId,
		StartTime: startTime,
		EndTime:   endTime,
	}
}
