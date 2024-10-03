//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type OrderDetail struct {
	Id        int64   `json:"id" gorm:"primaryKey"` //
	OrderId   int64   `json:"order_id" `            //
	Price     float64 `json:"price" `               //
	Num       int64   `json:"num" `                 //
	Amount    float64 `json:"amount" `              //
	GoodsType int64   `json:"goods_type" `          //
	GoodsId   int64   `json:"goods_id" `            //
	ExtraData string  `json:"extra_data" `          //
}

const OrderDetailCacheKey = "order_details_pk_%s"

func (OrderDetail) TableName() string {
	return "order_details"
}

// BatchUpsert 批量插入或更新
func (r *OrderDetailRepo) BatchUpsert(insertSlice ...*OrderDetail) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"order_id":   db.Raw("values(order_id)"),
				"price":      db.Raw("values(price)"),
				"num":        db.Raw("values(num)"),
				"amount":     db.Raw("values(amount)"),
				"goods_type": db.Raw("values(goods_type)"),
				"goods_id":   db.Raw("values(goods_id)"),
				"extra_data": db.Raw("values(extra_data)"),
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

func (r *OrderDetailRepo) BatchInsert(insertSlice ...*OrderDetail) (int64, error) {
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
func (r *OrderDetailRepo) Insert(insert *OrderDetail) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *OrderDetailRepo) getAllFields() []string {
	return []string{
		"id",
		"order_id",
		"price",
		"num",
		"amount",
		"goods_type",
		"goods_id",
		"extra_data",
	}
}

// Omit 过滤自己不想要的字段
func (r *OrderDetailRepo) Omit(filter []string) []string {
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
func (r *OrderDetailRepo) FindById(id int64) (OrderDetail, error) {
	var orderDetail OrderDetail
	res := r.db.Where("id = ?", id).First(&orderDetail)
	return orderDetail, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *OrderDetailRepo) FindByIdArr(idArr []int64) []OrderDetail {
	var orderDetailArr = make([]OrderDetail, 0)
	r.db.Where("id IN ?", idArr).Find(&orderDetailArr)
	return orderDetailArr
}
