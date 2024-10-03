package model

import "gorm.io/gorm"

type TravelRateRepo struct {
	db *gorm.DB
}

func NewTravelRateRepo(db *gorm.DB) *TravelRateRepo {
	return &TravelRateRepo{db: db}
}
