//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type PlanLike struct {
	Id        int64 `json:"id" gorm:"primaryKey"` //
	ServiceId int64 `json:"service_id" `          //
	UserId    int64 `json:"user_id" `             //
	Status    int64 `json:"status" `              //
	CreatedAt int64 `json:"created_at" `          //
	UpdatedAt int64 `json:"updated_at" `          //
}

const PlanLikeCacheKey = "plan_likes_pk_%s"

func (PlanLike) TableName() string {
	return "plan_likes"
}

// BatchUpsert 批量插入或更新
func (r *PlanLikeRepo) BatchUpsert(insertSlice ...*PlanLike) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"service_id": db.Raw("values(service_id)"),
				"user_id":    db.Raw("values(user_id)"),
				"status":     db.Raw("values(status)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
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

func (r *PlanLikeRepo) BatchInsert(insertSlice ...*PlanLike) (int64, error) {
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
func (r *PlanLikeRepo) Insert(insert *PlanLike) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *PlanLikeRepo) getAllFields() []string {
	return []string{
		"id",
		"service_id",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *PlanLikeRepo) Omit(filter []string) []string {
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
func (r *PlanLikeRepo) FindById(id int64) (PlanLike, error) {
	var planLike PlanLike
	res := r.db.Where("id = ?", id).First(&planLike)
	return planLike, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *PlanLikeRepo) FindByIdArr(idArr []int64) []PlanLike {
	var planLikeArr = make([]PlanLike, 0)
	r.db.Where("id IN ?", idArr).Find(&planLikeArr)
	return planLikeArr
}
