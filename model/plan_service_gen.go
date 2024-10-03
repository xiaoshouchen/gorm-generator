//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type PlanService struct {
	Id            int64                 `json:"id" gorm:"primaryKey"` //
	ServiceId     int64                 `json:"service_id" `          //
	ServiceTypeId int64                 `json:"service_type_id" `     //
	Price         float64               `json:"price" `               //
	CreatedAt     int64                 `json:"created_at" `          //
	UpdatedAt     int64                 `json:"updated_at" `          //
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at" `          //
}

const PlanServiceCacheKey = "plan_services_pk_%s"

func (PlanService) TableName() string {
	return "plan_services"
}

// BatchUpsert 批量插入或更新
func (r *PlanServiceRepo) BatchUpsert(insertSlice ...*PlanService) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"service_id":      db.Raw("values(service_id)"),
				"service_type_id": db.Raw("values(service_type_id)"),
				"price":           db.Raw("values(price)"),
				"created_at":      db.Raw("values(created_at)"),
				"updated_at":      db.Raw("values(updated_at)"),
				"deleted_at":      db.Raw("values(deleted_at)"),
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

func (r *PlanServiceRepo) BatchInsert(insertSlice ...*PlanService) (int64, error) {
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
func (r *PlanServiceRepo) Insert(insert *PlanService) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *PlanServiceRepo) getAllFields() []string {
	return []string{
		"id",
		"service_id",
		"service_type_id",
		"price",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *PlanServiceRepo) Omit(filter []string) []string {
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
func (r *PlanServiceRepo) FindById(id int64) (PlanService, error) {
	var planService PlanService
	res := r.db.Where("id = ?", id).First(&planService)
	return planService, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *PlanServiceRepo) FindByIdArr(idArr []int64) []PlanService {
	var planServiceArr = make([]PlanService, 0)
	r.db.Where("id IN ?", idArr).Find(&planServiceArr)
	return planServiceArr
}

// FindByServiceId 根据ServiceId进行查询
func (r *PlanServiceRepo) FindByServiceId(serviceId int64, limit int, orderBy string) ([]PlanService, error) {
	var planServiceArr []PlanService
	res := r.db.Where("service_id = ?", serviceId).Limit(limit).Order(orderBy).Find(&planServiceArr)
	return planServiceArr, res.Error
}
