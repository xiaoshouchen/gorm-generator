package model

import "gorm.io/gorm"

type TravelRepo struct {
	db *gorm.DB
}

func NewTravelRepo(db *gorm.DB) *TravelRepo {
	return &TravelRepo{db: db}
}
