//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type TourGuideAssignment struct {
	Id          int64  `json:"id" gorm:"primaryKey"` //
	TourGuideId int64  `json:"tour_guide_id" `       //
	Year        int64  `json:"year" `                //
	Month       int64  `json:"month" `               //
	Busy        []byte `json:"busy" `                //
	CreatedAt   int64  `json:"created_at" `          //
	UpdatedAt   int64  `json:"updated_at" `          //
}

const TourGuideAssignmentCacheKey = "tour_guide_assignment_pk_%s"

func (TourGuideAssignment) TableName() string {
	return "tour_guide_assignment"
}

// BatchUpsert 批量插入或更新
func (r *TourGuideAssignmentRepo) BatchUpsert(insertSlice ...*TourGuideAssignment) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"year":          db.Raw("values(year)"),
				"month":         db.Raw("values(month)"),
				"busy":          db.Raw("values(busy)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
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

func (r *TourGuideAssignmentRepo) BatchInsert(insertSlice ...*TourGuideAssignment) (int64, error) {
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
func (r *TourGuideAssignmentRepo) Insert(insert *TourGuideAssignment) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TourGuideAssignmentRepo) getAllFields() []string {
	return []string{
		"id",
		"tour_guide_id",
		"year",
		"month",
		"busy",
		"created_at",
		"updated_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TourGuideAssignmentRepo) Omit(filter []string) []string {
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
func (r *TourGuideAssignmentRepo) FindById(id int64) (TourGuideAssignment, error) {
	var tourGuideAssignment TourGuideAssignment
	res := r.db.Where("id = ?", id).First(&tourGuideAssignment)
	return tourGuideAssignment, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TourGuideAssignmentRepo) FindByIdArr(idArr []int64) []TourGuideAssignment {
	var tourGuideAssignmentArr = make([]TourGuideAssignment, 0)
	r.db.Where("id IN ?", idArr).Find(&tourGuideAssignmentArr)
	return tourGuideAssignmentArr
}

// FindByTourGuideId 根据TourGuideId进行查询
func (r *TourGuideAssignmentRepo) FindByTourGuideId(tourGuideId int64, limit int, orderBy string) ([]TourGuideAssignment, error) {
	var tourGuideAssignmentArr []TourGuideAssignment
	res := r.db.Where("tour_guide_id = ?", tourGuideId).Limit(limit).Order(orderBy).Find(&tourGuideAssignmentArr)
	return tourGuideAssignmentArr, res.Error
}
