package model

import "gorm.io/gorm"

type FeeRepo struct {
	db *gorm.DB
}

func NewFeeRepo(db *gorm.DB) *FeeRepo {
	return &FeeRepo{db: db}
}
