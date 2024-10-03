//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type TourGuideArea struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	AreaId      int64                 `json:"area_id" `             //
	TourGuideId int64                 `json:"tour_guide_id" `       //
	Status      int64                 `json:"status" `              //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const TourGuideAreaCacheKey = "tour_guide_areas_pk_%s"

func (TourGuideArea) TableName() string {
	return "tour_guide_areas"
}

// BatchUpsert 批量插入或更新
func (r *TourGuideAreaRepo) BatchUpsert(insertSlice ...*TourGuideArea) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"area_id":       db.Raw("values(area_id)"),
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"status":        db.Raw("values(status)"),
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

func (r *TourGuideAreaRepo) BatchInsert(insertSlice ...*TourGuideArea) (int64, error) {
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
func (r *TourGuideAreaRepo) Insert(insert *TourGuideArea) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TourGuideAreaRepo) getAllFields() []string {
	return []string{
		"id",
		"area_id",
		"tour_guide_id",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TourGuideAreaRepo) Omit(filter []string) []string {
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
func (r *TourGuideAreaRepo) FindById(id int64) (TourGuideArea, error) {
	var tourGuideArea TourGuideArea
	res := r.db.Where("id = ?", id).First(&tourGuideArea)
	return tourGuideArea, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TourGuideAreaRepo) FindByIdArr(idArr []int64) []TourGuideArea {
	var tourGuideAreaArr = make([]TourGuideArea, 0)
	r.db.Where("id IN ?", idArr).Find(&tourGuideAreaArr)
	return tourGuideAreaArr
}

// FindByAreaId 根据AreaId进行查询
func (r *TourGuideAreaRepo) FindByAreaId(areaId int64, limit int, orderBy string) ([]TourGuideArea, error) {
	var tourGuideAreaArr []TourGuideArea
	res := r.db.Where("area_id = ?", areaId).Limit(limit).Order(orderBy).Find(&tourGuideAreaArr)
	return tourGuideAreaArr, res.Error
}

// FindByTourGuideId 根据TourGuideId进行查询
func (r *TourGuideAreaRepo) FindByTourGuideId(tourGuideId int64, limit int, orderBy string) ([]TourGuideArea, error) {
	var tourGuideAreaArr []TourGuideArea
	res := r.db.Where("tour_guide_id = ?", tourGuideId).Limit(limit).Order(orderBy).Find(&tourGuideAreaArr)
	return tourGuideAreaArr, res.Error
}
