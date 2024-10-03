//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type PlanFee struct {
	Id        int64 `json:"id" gorm:"primaryKey"` //
	ServiceId int64 `json:"service_id" `          //
	FeeTypeId int64 `json:"fee_type_id" `         //
}

const PlanFeeCacheKey = "plan_fees_pk_%s"

func (PlanFee) TableName() string {
	return "plan_fees"
}

// BatchUpsert 批量插入或更新
func (r *PlanFeeRepo) BatchUpsert(insertSlice ...*PlanFee) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"service_id":  db.Raw("values(service_id)"),
				"fee_type_id": db.Raw("values(fee_type_id)"),
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

func (r *PlanFeeRepo) BatchInsert(insertSlice ...*PlanFee) (int64, error) {
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
func (r *PlanFeeRepo) Insert(insert *PlanFee) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *PlanFeeRepo) getAllFields() []string {
	return []string{
		"id",
		"service_id",
		"fee_type_id",
	}
}

// Omit 过滤自己不想要的字段
func (r *PlanFeeRepo) Omit(filter []string) []string {
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
func (r *PlanFeeRepo) FindById(id int64) (PlanFee, error) {
	var planFee PlanFee
	res := r.db.Where("id = ?", id).First(&planFee)
	return planFee, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *PlanFeeRepo) FindByIdArr(idArr []int64) []PlanFee {
	var planFeeArr = make([]PlanFee, 0)
	r.db.Where("id IN ?", idArr).Find(&planFeeArr)
	return planFeeArr
}
