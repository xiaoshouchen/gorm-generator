//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type Tag struct {
	Id   int64   `json:"id" gorm:"primaryKey"` //
	Name *string `json:"name" `                //
}

const TagCacheKey = "tags_pk_%s"

func (Tag) TableName() string {
	return "tags"
}

// BatchUpsert 批量插入或更新
func (r *TagRepo) BatchUpsert(insertSlice ...*Tag) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"name": db.Raw("values(name)"),
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

func (r *TagRepo) BatchInsert(insertSlice ...*Tag) (int64, error) {
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
func (r *TagRepo) Insert(insert *Tag) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TagRepo) getAllFields() []string {
	return []string{
		"id",
		"name",
	}
}

// Omit 过滤自己不想要的字段
func (r *TagRepo) Omit(filter []string) []string {
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
func (r *TagRepo) FindById(id int64) (Tag, error) {
	var tag Tag
	res := r.db.Where("id = ?", id).First(&tag)
	return tag, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TagRepo) FindByIdArr(idArr []int64) []Tag {
	var tagArr = make([]Tag, 0)
	r.db.Where("id IN ?", idArr).Find(&tagArr)
	return tagArr
}
