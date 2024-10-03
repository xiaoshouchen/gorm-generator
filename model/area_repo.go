package model

import "gorm.io/gorm"

type AreaRepo struct {
	db *gorm.DB
}

func NewAreaRepo(db *gorm.DB) *AreaRepo {
	return &AreaRepo{db: db}
}
