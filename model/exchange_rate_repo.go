package model

import "gorm.io/gorm"

type ExchangeRateRepo struct {
	db *gorm.DB
}

func NewExchangeRateRepo(db *gorm.DB) *ExchangeRateRepo {
	return &ExchangeRateRepo{db: db}
}
