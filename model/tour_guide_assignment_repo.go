package model

import "gorm.io/gorm"

type TourGuideAssignmentRepo struct {
	db *gorm.DB
}

func NewTourGuideAssignmentRepo(db *gorm.DB) *TourGuideAssignmentRepo {
	return &TourGuideAssignmentRepo{db: db}
}
