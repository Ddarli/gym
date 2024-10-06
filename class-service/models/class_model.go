package models

type ClassModel struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Capacity    int    `json:"capacity" db:"capacity"`
}
