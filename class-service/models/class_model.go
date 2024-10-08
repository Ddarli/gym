package models

import "strconv"

type ClassModel struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Capacity    int    `json:"capacity" db:"capacity"`
}

func ToProto(class *ClassModel) *Class {
	return &Class{
		Id:          strconv.Itoa(class.Id),
		Name:        class.Name,
		Description: class.Description,
		Capacity:    int32(class.Capacity),
	}
}
