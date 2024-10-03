//go:gen DON'T EDIT !
package model

import (
	"time"

	"gorm.io/gorm/clause"
)

type TravelTimeRecord struct {
	Id         int64     `json:"id" gorm:"primaryKey"` //
	StartTime  int64     `json:"start_time" `          //
	EndTime    int64     `json:"end_time" `            //
	TravelDate time.Time `json:"travel_date" `         //
	TravelId   int64     `json:"travel_id" `           //
}

const TravelTimeRecordCacheKey = "travel_time_records_pk_%s"

func (TravelTimeRecord) TableName() string {
	return "travel_time_records"
}

// BatchUpsert 批量插入或更新
func (r *TravelTimeRecordRepo) BatchUpsert(insertSlice ...*TravelTimeRecord) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"start_time":  db.Raw("values(start_time)"),
				"end_time":    db.Raw("values(end_time)"),
				"travel_date": db.Raw("values(travel_date)"),
				"travel_id":   db.Raw("values(travel_id)"),
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

func (r *TravelTimeRecordRepo) BatchInsert(insertSlice ...*TravelTimeRecord) (int64, error) {
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
func (r *TravelTimeRecordRepo) Insert(insert *TravelTimeRecord) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TravelTimeRecordRepo) getAllFields() []string {
	return []string{
		"id",
		"start_time",
		"end_time",
		"travel_date",
		"travel_id",
	}
}

// Omit 过滤自己不想要的字段
func (r *TravelTimeRecordRepo) Omit(filter []string) []string {
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
func (r *TravelTimeRecordRepo) FindById(id int64) (TravelTimeRecord, error) {
	var travelTimeRecord TravelTimeRecord
	res := r.db.Where("id = ?", id).First(&travelTimeRecord)
	return travelTimeRecord, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TravelTimeRecordRepo) FindByIdArr(idArr []int64) []TravelTimeRecord {
	var travelTimeRecordArr = make([]TravelTimeRecord, 0)
	r.db.Where("id IN ?", idArr).Find(&travelTimeRecordArr)
	return travelTimeRecordArr
}
