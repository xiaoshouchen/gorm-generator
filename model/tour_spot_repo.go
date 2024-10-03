package model

import "gorm.io/gorm"

type TourSpotRepo struct {
	db *gorm.DB
}

func NewTourSpotRepo(db *gorm.DB) *TourSpotRepo {
	return &TourSpotRepo{db: db}
}
