package model

import "gorm.io/gorm"

type GoogleUserRepo struct {
	db *gorm.DB
}

func NewGoogleUserRepo(db *gorm.DB) *GoogleUserRepo {
	return &GoogleUserRepo{db: db}
}
