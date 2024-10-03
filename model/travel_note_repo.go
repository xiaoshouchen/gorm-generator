package model

import "gorm.io/gorm"

type TravelNoteRepo struct {
	db *gorm.DB
}

func NewTravelNoteRepo(db *gorm.DB) *TravelNoteRepo {
	return &TravelNoteRepo{db: db}
}
