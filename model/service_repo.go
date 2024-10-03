package model

import "gorm.io/gorm"

type ServiceRepo struct {
	db *gorm.DB
}

func NewServiceRepo(db *gorm.DB) *ServiceRepo {
	return &ServiceRepo{db: db}
}
