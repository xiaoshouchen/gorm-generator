//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type Coupon struct {
	Id         int64 `json:"id" gorm:"primaryKey"` //
	CouponType int64 `json:"coupon_type" `         //
}

const CouponCacheKey = "coupons_pk_%s"

func (Coupon) TableName() string {
	return "coupons"
}

// BatchUpsert 批量插入或更新
func (r *CouponRepo) BatchUpsert(insertSlice ...*Coupon) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"coupon_type": db.Raw("values(coupon_type)"),
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

func (r *CouponRepo) BatchInsert(insertSlice ...*Coupon) (int64, error) {
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
func (r *CouponRepo) Insert(insert *Coupon) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *CouponRepo) getAllFields() []string {
	return []string{
		"id",
		"coupon_type",
	}
}

// Omit 过滤自己不想要的字段
func (r *CouponRepo) Omit(filter []string) []string {
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
func (r *CouponRepo) FindById(id int64) (Coupon, error) {
	var coupon Coupon
	res := r.db.Where("id = ?", id).First(&coupon)
	return coupon, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *CouponRepo) FindByIdArr(idArr []int64) []Coupon {
	var couponArr = make([]Coupon, 0)
	r.db.Where("id IN ?", idArr).Find(&couponArr)
	return couponArr
}
