package model

import "gorm.io/gorm"

type TourGuideRepo struct {
	db *gorm.DB
}

func NewTourGuideRepo(db *gorm.DB) *TourGuideRepo {
	return &TourGuideRepo{db: db}
}
