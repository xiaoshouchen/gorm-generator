package model

import "gorm.io/gorm"

type CheckPointRepo struct {
	db *gorm.DB
}

func NewCheckPointRepo(db *gorm.DB) *CheckPointRepo {
	return &CheckPointRepo{db: db}
}
