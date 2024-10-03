package model

import "gorm.io/gorm"

type TourGuideAreaRepo struct {
	db *gorm.DB
}

func NewTourGuideAreaRepo(db *gorm.DB) *TourGuideAreaRepo {
	return &TourGuideAreaRepo{db: db}
}
