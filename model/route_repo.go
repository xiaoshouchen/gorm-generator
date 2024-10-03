package model

import "gorm.io/gorm"

type RouteRepo struct {
	db *gorm.DB
}

func NewRouteRepo(db *gorm.DB) *RouteRepo {
	return &RouteRepo{db: db}
}
