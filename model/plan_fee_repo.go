package model

import "gorm.io/gorm"

type PlanFeeRepo struct {
	db *gorm.DB
}

func NewPlanFeeRepo(db *gorm.DB) *PlanFeeRepo {
	return &PlanFeeRepo{db: db}
}
