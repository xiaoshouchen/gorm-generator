package model

import "gorm.io/gorm"

type PlanServiceRepo struct {
	db *gorm.DB
}

func NewPlanServiceRepo(db *gorm.DB) *PlanServiceRepo {
	return &PlanServiceRepo{db: db}
}
