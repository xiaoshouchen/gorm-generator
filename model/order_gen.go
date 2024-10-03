//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Order struct {
	Id           int64                 `json:"id" gorm:"primaryKey"` //
	Slug         string                `json:"slug" `                //
	OrderNum     string                `json:"order_num" `           //
	UserId       int64                 `json:"user_id" `             //
	Status       int64                 `json:"status" `              //
	Amount       float64               `json:"amount" `              //
	CreatedAt    int64                 `json:"created_at" `          //
	UpdatedAt    int64                 `json:"updated_at" `          //
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" `          //
	PaidAt       int64                 `json:"paid_at" `             //
	CurrencyType int64                 `json:"currency_type" `       //
	AmountRmb    float64               `json:"amount_rmb" `          //
}

const OrderCacheKey = "orders_pk_%s"

func (Order) TableName() string {
	return "orders"
}

// BatchUpsert 批量插入或更新
func (r *OrderRepo) BatchUpsert(insertSlice ...*Order) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":          db.Raw("values(slug)"),
				"order_num":     db.Raw("values(order_num)"),
				"user_id":       db.Raw("values(user_id)"),
				"status":        db.Raw("values(status)"),
				"amount":        db.Raw("values(amount)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
				"deleted_at":    db.Raw("values(deleted_at)"),
				"paid_at":       db.Raw("values(paid_at)"),
				"currency_type": db.Raw("values(currency_type)"),
				"amount_rmb":    db.Raw("values(amount_rmb)"),
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

func (r *OrderRepo) BatchInsert(insertSlice ...*Order) (int64, error) {
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
func (r *OrderRepo) Insert(insert *Order) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *OrderRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"order_num",
		"user_id",
		"status",
		"amount",
		"created_at",
		"updated_at",
		"deleted_at",
		"paid_at",
		"currency_type",
		"amount_rmb",
	}
}

// Omit 过滤自己不想要的字段
func (r *OrderRepo) Omit(filter []string) []string {
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
func (r *OrderRepo) FindById(id int64) (Order, error) {
	var order Order
	res := r.db.Where("id = ?", id).First(&order)
	return order, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *OrderRepo) FindByIdArr(idArr []int64) []Order {
	var orderArr = make([]Order, 0)
	r.db.Where("id IN ?", idArr).Find(&orderArr)
	return orderArr
}

// FindByOrderNum 根据OrderNum进行查询
func (r *OrderRepo) FindByOrderNum(orderNum string) (Order, error) {
	var order Order
	res := r.db.Where("order_num = ?", orderNum).First(&order)
	return order, res.Error
}

// FindBySlug 根据Slug进行查询
func (r *OrderRepo) FindBySlug(slug string) (Order, error) {
	var order Order
	res := r.db.Where("slug = ?", slug).First(&order)
	return order, res.Error
}

// FindByOrderNumArr 根据主键OrderNum进行查询
func (r *OrderRepo) FindByOrderNumArr(arr []string) ([]Order, error) {
	var orderArr []Order
	res := r.db.Where("order_num IN ?", arr).Find(&orderArr)
	return orderArr, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *OrderRepo) FindBySlugArr(arr []string) ([]Order, error) {
	var orderArr []Order
	res := r.db.Where("slug IN ?", arr).Find(&orderArr)
	return orderArr, res.Error
}
