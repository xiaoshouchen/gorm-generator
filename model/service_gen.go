//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Service struct {
	Id        int64                 `json:"id" gorm:"primaryKey"` //
	Slug      string                `json:"slug" `                //
	Name      string                `json:"name" `                //
	CreatedAt int64                 `json:"created_at" `          //
	UpdatedAt int64                 `json:"updated_at" `          //
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" `          //
}

const ServiceCacheKey = "services_pk_%s"

func (Service) TableName() string {
	return "services"
}

// BatchUpsert 批量插入或更新
func (r *ServiceRepo) BatchUpsert(insertSlice ...*Service) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":       db.Raw("values(slug)"),
				"name":       db.Raw("values(name)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
				"deleted_at": db.Raw("values(deleted_at)"),
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

func (r *ServiceRepo) BatchInsert(insertSlice ...*Service) (int64, error) {
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
func (r *ServiceRepo) Insert(insert *Service) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *ServiceRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *ServiceRepo) Omit(filter []string) []string {
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
func (r *ServiceRepo) FindById(id int64) (Service, error) {
	var service Service
	res := r.db.Where("id = ?", id).First(&service)
	return service, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *ServiceRepo) FindByIdArr(idArr []int64) []Service {
	var serviceArr = make([]Service, 0)
	r.db.Where("id IN ?", idArr).Find(&serviceArr)
	return serviceArr
}

// FindBySlug 根据Slug进行查询
func (r *ServiceRepo) FindBySlug(slug string) (Service, error) {
	var service Service
	res := r.db.Where("slug = ?", slug).First(&service)
	return service, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *ServiceRepo) FindBySlugArr(arr []string) ([]Service, error) {
	var serviceArr []Service
	res := r.db.Where("slug IN ?", arr).Find(&serviceArr)
	return serviceArr, res.Error
}
