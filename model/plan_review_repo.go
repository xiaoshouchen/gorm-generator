package model

import "gorm.io/gorm"

type PlanReviewRepo struct {
	db *gorm.DB
}

func NewPlanReviewRepo(db *gorm.DB) *PlanReviewRepo {
	return &PlanReviewRepo{db: db}
}
