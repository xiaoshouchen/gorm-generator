//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type CheckPoint struct {
	Id         int64                 `json:"id" gorm:"primaryKey"` //
	TravelId   int64                 `json:"travel_id" `           //
	TourSpotId int64                 `json:"tour_spot_id" `        //
	Images     interface{}           `json:"images" `              //
	CreatedAt  int64                 `json:"created_at" `          //
	UpdatedAt  int64                 `json:"updated_at" `          //
	DeletedAt  soft_delete.DeletedAt `json:"deleted_at" `          //
	GeoX       string                `json:"geo_x" `               //
	GeoY       string                `json:"geo_y" `               //
}

const CheckPointCacheKey = "check_points_pk_%s"

func (CheckPoint) TableName() string {
	return "check_points"
}

// BatchUpsert 批量插入或更新
func (r *CheckPointRepo) BatchUpsert(insertSlice ...*CheckPoint) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"travel_id":    db.Raw("values(travel_id)"),
				"tour_spot_id": db.Raw("values(tour_spot_id)"),
				"images":       db.Raw("values(images)"),
				"created_at":   db.Raw("values(created_at)"),
				"updated_at":   db.Raw("values(updated_at)"),
				"deleted_at":   db.Raw("values(deleted_at)"),
				"geo_x":        db.Raw("values(geo_x)"),
				"geo_y":        db.Raw("values(geo_y)"),
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

func (r *CheckPointRepo) BatchInsert(insertSlice ...*CheckPoint) (int64, error) {
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
func (r *CheckPointRepo) Insert(insert *CheckPoint) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *CheckPointRepo) getAllFields() []string {
	return []string{
		"id",
		"travel_id",
		"tour_spot_id",
		"images",
		"created_at",
		"updated_at",
		"deleted_at",
		"geo_x",
		"geo_y",
	}
}

// Omit 过滤自己不想要的字段
func (r *CheckPointRepo) Omit(filter []string) []string {
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
func (r *CheckPointRepo) FindById(id int64) (CheckPoint, error) {
	var checkPoint CheckPoint
	res := r.db.Where("id = ?", id).First(&checkPoint)
	return checkPoint, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *CheckPointRepo) FindByIdArr(idArr []int64) []CheckPoint {
	var checkPointArr = make([]CheckPoint, 0)
	r.db.Where("id IN ?", idArr).Find(&checkPointArr)
	return checkPointArr
}
