package model

import "gorm.io/gorm"

type PlanLikeRepo struct {
	db *gorm.DB
}

func NewPlanLikeRepo(db *gorm.DB) *PlanLikeRepo {
	return &PlanLikeRepo{db: db}
}
