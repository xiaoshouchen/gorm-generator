//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type Bill struct {
	Id        int64    `json:"id" gorm:"primaryKey"` //
	UserId    *int64   `json:"user_id" `             //
	UserSlug  *string  `json:"user_slug" `           //
	Amount    *float64 `json:"amount" `              //
	CreatedAt *int64   `json:"created_at" `          //
	UpdatedAt *int64   `json:"updated_at" `          //
	DeletedAt *int64   `json:"deleted_at" `          //
	Status    *int64   `json:"status" `              //
}

const BillCacheKey = "bills_pk_%s"

func (Bill) TableName() string {
	return "bills"
}

// BatchUpsert 批量插入或更新
func (r *BillRepo) BatchUpsert(insertSlice ...*Bill) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"user_id":    db.Raw("values(user_id)"),
				"user_slug":  db.Raw("values(user_slug)"),
				"amount":     db.Raw("values(amount)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
				"deleted_at": db.Raw("values(deleted_at)"),
				"status":     db.Raw("values(status)"),
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

func (r *BillRepo) BatchInsert(insertSlice ...*Bill) (int64, error) {
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
func (r *BillRepo) Insert(insert *Bill) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *BillRepo) getAllFields() []string {
	return []string{
		"id",
		"user_id",
		"user_slug",
		"amount",
		"created_at",
		"updated_at",
		"deleted_at",
		"status",
	}
}

// Omit 过滤自己不想要的字段
func (r *BillRepo) Omit(filter []string) []string {
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
func (r *BillRepo) FindById(id int64) (Bill, error) {
	var bill Bill
	res := r.db.Where("id = ?", id).First(&bill)
	return bill, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *BillRepo) FindByIdArr(idArr []int64) []Bill {
	var billArr = make([]Bill, 0)
	r.db.Where("id IN ?", idArr).Find(&billArr)
	return billArr
}
