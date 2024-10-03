package model

import "gorm.io/gorm"

type BillRepo struct {
	db *gorm.DB
}

func NewBillRepo(db *gorm.DB) *BillRepo {
	return &BillRepo{db: db}
}
