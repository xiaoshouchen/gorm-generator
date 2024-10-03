//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type TravelRate struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	UserId      int64                 `json:"user_id" `             //发布评价的ID
	TravelId    int64                 `json:"travel_id" `           //被评价的导游ID
	TourGuideId int64                 `json:"tour_guide_id" `       //
	Level       int64                 `json:"level" `               //1-5
	Content     string                `json:"content" `             //评论内容
	Images      interface{}           `json:"images" `              //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const TravelRateCacheKey = "travel_rates_pk_%s"

func (TravelRate) TableName() string {
	return "travel_rates"
}

// BatchUpsert 批量插入或更新
func (r *TravelRateRepo) BatchUpsert(insertSlice ...*TravelRate) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"user_id":       db.Raw("values(user_id)"),
				"travel_id":     db.Raw("values(travel_id)"),
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"level":         db.Raw("values(level)"),
				"content":       db.Raw("values(content)"),
				"images":        db.Raw("values(images)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
				"deleted_at":    db.Raw("values(deleted_at)"),
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

func (r *TravelRateRepo) BatchInsert(insertSlice ...*TravelRate) (int64, error) {
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
func (r *TravelRateRepo) Insert(insert *TravelRate) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TravelRateRepo) getAllFields() []string {
	return []string{
		"id",
		"user_id",
		"travel_id",
		"tour_guide_id",
		"level",
		"content",
		"images",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TravelRateRepo) Omit(filter []string) []string {
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
func (r *TravelRateRepo) FindById(id int64) (TravelRate, error) {
	var travelRate TravelRate
	res := r.db.Where("id = ?", id).First(&travelRate)
	return travelRate, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TravelRateRepo) FindByIdArr(idArr []int64) []TravelRate {
	var travelRateArr = make([]TravelRate, 0)
	r.db.Where("id IN ?", idArr).Find(&travelRateArr)
	return travelRateArr
}
