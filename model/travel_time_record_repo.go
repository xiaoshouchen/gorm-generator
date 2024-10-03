package model

import "gorm.io/gorm"

type TravelTimeRecordRepo struct {
	db *gorm.DB
}

func NewTravelTimeRecordRepo(db *gorm.DB) *TravelTimeRecordRepo {
	return &TravelTimeRecordRepo{db: db}
}
