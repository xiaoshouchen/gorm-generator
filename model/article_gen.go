//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type Article struct {
	Id        int64       `json:"id" gorm:"primaryKey"` //
	UserId    int64       `json:"user_id" `             //
	UserSlug  string      `json:"user_slug" `           //
	Title     string      `json:"title" `               //
	Content   string      `json:"content" `             //
	Images    interface{} `json:"images" `              //
	CreatedAt *int64      `json:"created_at" `          //
	UpdatedAt *int64      `json:"updated_at" `          //
}

const ArticleCacheKey = "articles_pk_%s"

func (Article) TableName() string {
	return "articles"
}

// BatchUpsert 批量插入或更新
func (r *ArticleRepo) BatchUpsert(insertSlice ...*Article) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"user_id":    db.Raw("values(user_id)"),
				"user_slug":  db.Raw("values(user_slug)"),
				"title":      db.Raw("values(title)"),
				"content":    db.Raw("values(content)"),
				"images":     db.Raw("values(images)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
			}),
		},
	)

	if len(insertSlice) > 1000 {
		db = db.CreateInBatches(insertSlice, 1000)

	} else {
		db = db.Create(insertSlice)
	}
	return db.RowsAffected, db.Error
}

func (r *ArticleRepo) BatchInsert(insertSlice ...*Article) (int64, error) {
	db := r.db
	if len(insertSlice) > 1000 {
		db = db.CreateInBatches(insertSlice, 1000)

	} else {
		db = db.Create(insertSlice)
	}
	return db.RowsAffected, db.Error
}

// Insert 插入单个
// return id
func (r *ArticleRepo) Insert(insert *Article) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *ArticleRepo) getAllFields() []string {
	return []string{
		"id",
		"user_id",
		"user_slug",
		"title",
		"content",
		"images",
		"created_at",
		"updated_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *ArticleRepo) Omit(filter []string) []string {
	fields := r.getAllFields()
	result := make([]string, 0, len(fields))
	filterSet := make(map[string]bool)

	// 将需要过滤的值添加到 filterSet 中
	for _, v := range filter {
		filterSet[v] = true
	}

	// 遍历原始切片，将不在 filterSet 中的值添加到结果切片中
	for _, v := range fields {
		if _, ok := filterSet[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

// FindById 根据主键进行查询
func (r *ArticleRepo) FindById(id int64) (Article, error) {
	var article Article
	res := r.db.Where("id = ?", id).First(&article)
	return article, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *ArticleRepo) FindByIdArr(idArr []int64) []Article {
	var articleArr = make([]Article, 0)
	r.db.Where("id IN ?", idArr).Find(&articleArr)
	return articleArr
}
