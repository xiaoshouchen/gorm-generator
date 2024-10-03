package model

import "gorm.io/gorm"

type OrderDetailRepo struct {
	db *gorm.DB
}

func NewOrderDetailRepo(db *gorm.DB) *OrderDetailRepo {
	return &OrderDetailRepo{db: db}
}
