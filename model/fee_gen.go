//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Fee struct {
	Id           int64                 `json:"id" gorm:"primaryKey"` //
	Slug         string                `json:"slug" `                //
	Name         string                `json:"name" `                //
	CreateUserId int64                 `json:"create_user_id" `      //
	CreatedAt    int64                 `json:"created_at" `          //
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" `          //
	UpdatedAt    int64                 `json:"updated_at" `          //
}

const FeeCacheKey = "fees_pk_%s"

func (Fee) TableName() string {
	return "fees"
}

// BatchUpsert 批量插入或更新
func (r *FeeRepo) BatchUpsert(insertSlice ...*Fee) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":           db.Raw("values(slug)"),
				"name":           db.Raw("values(name)"),
				"create_user_id": db.Raw("values(create_user_id)"),
				"created_at":     db.Raw("values(created_at)"),
				"deleted_at":     db.Raw("values(deleted_at)"),
				"updated_at":     db.Raw("values(updated_at)"),
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

func (r *FeeRepo) BatchInsert(insertSlice ...*Fee) (int64, error) {
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
func (r *FeeRepo) Insert(insert *Fee) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *FeeRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"name",
		"create_user_id",
		"created_at",
		"deleted_at",
		"updated_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *FeeRepo) Omit(filter []string) []string {
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
func (r *FeeRepo) FindById(id int64) (Fee, error) {
	var fee Fee
	res := r.db.Where("id = ?", id).First(&fee)
	return fee, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *FeeRepo) FindByIdArr(idArr []int64) []Fee {
	var feeArr = make([]Fee, 0)
	r.db.Where("id IN ?", idArr).Find(&feeArr)
	return feeArr
}

// FindBySlug 根据Slug进行查询
func (r *FeeRepo) FindBySlug(slug string) (Fee, error) {
	var fee Fee
	res := r.db.Where("slug = ?", slug).First(&fee)
	return fee, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *FeeRepo) FindBySlugArr(arr []string) ([]Fee, error) {
	var feeArr []Fee
	res := r.db.Where("slug IN ?", arr).Find(&feeArr)
	return feeArr, res.Error
}
