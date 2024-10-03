//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type PlanReview struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	TourGuideId int64                 `json:"tour_guide_id" `       //
	Rate        string                `json:"rate" `                //
	ServiceId   int64                 `json:"service_id" `          //
	Content     string                `json:"content" `             //
	UserId      int64                 `json:"user_id" `             //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdateAt    int64                 `json:"update_at" `           //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const PlanReviewCacheKey = "plan_reviews_pk_%s"

func (PlanReview) TableName() string {
	return "plan_reviews"
}

// BatchUpsert 批量插入或更新
func (r *PlanReviewRepo) BatchUpsert(insertSlice ...*PlanReview) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"rate":          db.Raw("values(rate)"),
				"service_id":    db.Raw("values(service_id)"),
				"content":       db.Raw("values(content)"),
				"user_id":       db.Raw("values(user_id)"),
				"created_at":    db.Raw("values(created_at)"),
				"update_at":     db.Raw("values(update_at)"),
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

func (r *PlanReviewRepo) BatchInsert(insertSlice ...*PlanReview) (int64, error) {
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
func (r *PlanReviewRepo) Insert(insert *PlanReview) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *PlanReviewRepo) getAllFields() []string {
	return []string{
		"id",
		"tour_guide_id",
		"rate",
		"service_id",
		"content",
		"user_id",
		"created_at",
		"update_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *PlanReviewRepo) Omit(filter []string) []string {
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
func (r *PlanReviewRepo) FindById(id int64) (PlanReview, error) {
	var planReview PlanReview
	res := r.db.Where("id = ?", id).First(&planReview)
	return planReview, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *PlanReviewRepo) FindByIdArr(idArr []int64) []PlanReview {
	var planReviewArr = make([]PlanReview, 0)
	r.db.Where("id IN ?", idArr).Find(&planReviewArr)
	return planReviewArr
}
