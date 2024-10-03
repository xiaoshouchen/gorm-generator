//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Area struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	Slug        string                `json:"slug" `                //
	Name        string                `json:"name" `                //
	Logo        string                `json:"logo" `                //
	Status      int64                 `json:"status" `              //状态 1 开启 2 关闭
	Description string                `json:"description" `         //
	Images      string                `json:"images" `              //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const AreaCacheKey = "areas_pk_%s"

func (Area) TableName() string {
	return "areas"
}

// BatchUpsert 批量插入或更新
func (r *AreaRepo) BatchUpsert(insertSlice ...*Area) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":        db.Raw("values(slug)"),
				"name":        db.Raw("values(name)"),
				"logo":        db.Raw("values(logo)"),
				"status":      db.Raw("values(status)"),
				"description": db.Raw("values(description)"),
				"images":      db.Raw("values(images)"),
				"created_at":  db.Raw("values(created_at)"),
				"updated_at":  db.Raw("values(updated_at)"),
				"deleted_at":  db.Raw("values(deleted_at)"),
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

func (r *AreaRepo) BatchInsert(insertSlice ...*Area) (int64, error) {
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
func (r *AreaRepo) Insert(insert *Area) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *AreaRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"name",
		"logo",
		"status",
		"description",
		"images",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *AreaRepo) Omit(filter []string) []string {
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
func (r *AreaRepo) FindById(id int64) (Area, error) {
	var area Area
	res := r.db.Where("id = ?", id).First(&area)
	return area, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *AreaRepo) FindByIdArr(idArr []int64) []Area {
	var areaArr = make([]Area, 0)
	r.db.Where("id IN ?", idArr).Find(&areaArr)
	return areaArr
}

// FindBySlug 根据Slug进行查询
func (r *AreaRepo) FindBySlug(slug string) (Area, error) {
	var area Area
	res := r.db.Where("slug = ?", slug).First(&area)
	return area, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *AreaRepo) FindBySlugArr(arr []string) ([]Area, error) {
	var areaArr []Area
	res := r.db.Where("slug IN ?", arr).Find(&areaArr)
	return areaArr, res.Error
}

// FindByName 根据Name进行查询
func (r *AreaRepo) FindByName(name string, limit int, orderBy string) ([]Area, error) {
	var areaArr []Area
	res := r.db.Where("name = ?", name).Limit(limit).Order(orderBy).Find(&areaArr)
	return areaArr, res.Error
}
