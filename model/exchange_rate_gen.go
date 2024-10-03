//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type ExchangeRate struct {
	Id           int64   `json:"id" gorm:"primaryKey"` //
	Rate         float64 `json:"rate" `                //
	CreatedAt    int64   `json:"created_at" `          //
	UpdatedAt    int64   `json:"updated_at" `          //
	CurrencyType int64   `json:"currency_type" `       //
}

const ExchangeRateCacheKey = "exchange_rates_pk_%s"

func (ExchangeRate) TableName() string {
	return "exchange_rates"
}

// BatchUpsert 批量插入或更新
func (r *ExchangeRateRepo) BatchUpsert(insertSlice ...*ExchangeRate) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"rate":          db.Raw("values(rate)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
				"currency_type": db.Raw("values(currency_type)"),
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

func (r *ExchangeRateRepo) BatchInsert(insertSlice ...*ExchangeRate) (int64, error) {
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
func (r *ExchangeRateRepo) Insert(insert *ExchangeRate) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *ExchangeRateRepo) getAllFields() []string {
	return []string{
		"id",
		"rate",
		"created_at",
		"updated_at",
		"currency_type",
	}
}

// Omit 过滤自己不想要的字段
func (r *ExchangeRateRepo) Omit(filter []string) []string {
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
func (r *ExchangeRateRepo) FindById(id int64) (ExchangeRate, error) {
	var exchangeRate ExchangeRate
	res := r.db.Where("id = ?", id).First(&exchangeRate)
	return exchangeRate, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *ExchangeRateRepo) FindByIdArr(idArr []int64) []ExchangeRate {
	var exchangeRateArr = make([]ExchangeRate, 0)
	r.db.Where("id IN ?", idArr).Find(&exchangeRateArr)
	return exchangeRateArr
}
