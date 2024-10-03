package model

import "gorm.io/gorm"

type TourGuideSkillRepo struct {
	db *gorm.DB
}

func NewTourGuideSkillRepo(db *gorm.DB) *TourGuideSkillRepo {
	return &TourGuideSkillRepo{db: db}
}
