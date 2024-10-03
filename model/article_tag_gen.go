//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type ArticleTag struct {
	Id        int64 `json:"id" gorm:"primaryKey"` //
	TagId     int64 `json:"tag_id" `              //
	ArticleId int64 `json:"article_id" `          //
}

const ArticleTagCacheKey = "article_tags_pk_%s"

func (ArticleTag) TableName() string {
	return "article_tags"
}

// BatchUpsert 批量插入或更新
func (r *ArticleTagRepo) BatchUpsert(insertSlice ...*ArticleTag) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"tag_id":     db.Raw("values(tag_id)"),
				"article_id": db.Raw("values(article_id)"),
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

func (r *ArticleTagRepo) BatchInsert(insertSlice ...*ArticleTag) (int64, error) {
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
func (r *ArticleTagRepo) Insert(insert *ArticleTag) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *ArticleTagRepo) getAllFields() []string {
	return []string{
		"id",
		"tag_id",
		"article_id",
	}
}

// Omit 过滤自己不想要的字段
func (r *ArticleTagRepo) Omit(filter []string) []string {
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
func (r *ArticleTagRepo) FindById(id int64) (ArticleTag, error) {
	var articleTag ArticleTag
	res := r.db.Where("id = ?", id).First(&articleTag)
	return articleTag, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *ArticleTagRepo) FindByIdArr(idArr []int64) []ArticleTag {
	var articleTagArr = make([]ArticleTag, 0)
	r.db.Where("id IN ?", idArr).Find(&articleTagArr)
	return articleTagArr
}
