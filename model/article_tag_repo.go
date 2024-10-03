package model

import "gorm.io/gorm"

type ArticleTagRepo struct {
	db *gorm.DB
}

func NewArticleTagRepo(db *gorm.DB) *ArticleTagRepo {
	return &ArticleTagRepo{db: db}
}
