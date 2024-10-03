package model

import "gorm.io/gorm"

type SkillRepo struct {
	db *gorm.DB
}

func NewSkillRepo(db *gorm.DB) *SkillRepo {
	return &SkillRepo{db: db}
}
