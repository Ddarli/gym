package models

import "strconv"

type TrainerModel struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}

func ToProto(trainer *TrainerModel) *Trainer {
	return &Trainer{
		Id:             strconv.Itoa(trainer.Id),
		Name:           trainer.Name,
		Specialization: trainer.Specialization,
	}
}
