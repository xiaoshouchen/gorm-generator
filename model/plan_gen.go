//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Plan struct {
	Id            int64                 `json:"id" gorm:"primaryKey"` //
	Slug          string                `json:"slug" `                //
	Name          string                `json:"name" `                //服务的名称
	Days          int64                 `json:"days" `                //天数
	BasePrice     float64               `json:"base_price" `          //套餐的基础价格，包含了base_people_num 的基本费用
	AdultPrice    float64               `json:"adult_price" `         //套餐外，添加一个成年人的价格
	ChildPrice    float64               `json:"child_price" `         //套餐外，添加一个儿童的价格
	Description   string                `json:"description" `         //
	Images        string                `json:"images" `              //展示照片
	TourGuideId   int64                 `json:"tour_guide_id" `       //
	ServiceType   int64                 `json:"service_type" `        //1 是 套餐服务 ，2 是定制服务
	AreaId        int64                 `json:"area_id" `             //
	BasePeopleNum int64                 `json:"base_people_num" `     //套餐基础包含的人数，比如默认包含2个人的费用
	MaxPeopleNum  int64                 `json:"max_people_num" `      //由于交通等各种因素，套餐允许的最大人数
	CreatedAt     int64                 `json:"created_at" `          //
	UpdatedAt     int64                 `json:"updated_at" `          //
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at" `          //
}

const PlanCacheKey = "plans_pk_%s"

func (Plan) TableName() string {
	return "plans"
}

// BatchUpsert 批量插入或更新
func (r *PlanRepo) BatchUpsert(insertSlice ...*Plan) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":            db.Raw("values(slug)"),
				"name":            db.Raw("values(name)"),
				"days":            db.Raw("values(days)"),
				"base_price":      db.Raw("values(base_price)"),
				"adult_price":     db.Raw("values(adult_price)"),
				"child_price":     db.Raw("values(child_price)"),
				"description":     db.Raw("values(description)"),
				"images":          db.Raw("values(images)"),
				"tour_guide_id":   db.Raw("values(tour_guide_id)"),
				"service_type":    db.Raw("values(service_type)"),
				"area_id":         db.Raw("values(area_id)"),
				"base_people_num": db.Raw("values(base_people_num)"),
				"max_people_num":  db.Raw("values(max_people_num)"),
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

func (r *PlanRepo) BatchInsert(insertSlice ...*Plan) (int64, error) {
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
func (r *PlanRepo) Insert(insert *Plan) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *PlanRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"name",
		"days",
		"base_price",
		"adult_price",
		"child_price",
		"description",
		"images",
		"tour_guide_id",
		"service_type",
		"area_id",
		"base_people_num",
		"max_people_num",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *PlanRepo) Omit(filter []string) []string {
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
func (r *PlanRepo) FindById(id int64) (Plan, error) {
	var plan Plan
	res := r.db.Where("id = ?", id).First(&plan)
	return plan, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *PlanRepo) FindByIdArr(idArr []int64) []Plan {
	var planArr = make([]Plan, 0)
	r.db.Where("id IN ?", idArr).Find(&planArr)
	return planArr
}
