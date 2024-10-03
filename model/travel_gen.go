//go:gen DON'T EDIT !
package model

import (
	"time"

	"gorm.io/gorm/clause"
)

type Travel struct {
	Id          int64     `json:"id" gorm:"primaryKey"` //
	Slug        string    `json:"slug" `                //
	AreaId      int64     `json:"area_id" `             //
	UserId      int64     `json:"user_id" `             //
	TourGuideId int64     `json:"tour_guide_id" `       //
	Status      int64     `json:"status" `              //1.待支付 2.待确认 3.旅游中，4 已结束 5.已取消
	StartDate   time.Time `json:"start_date" `          //
	EndDate     time.Time `json:"end_date" `            //
	Days        string    `json:"days" `                //
	TotalPrice  float64   `json:"total_price" `         //
	PlanId      int64     `json:"plan_id" `             //如果是0，则是自由行
	CreatedAt   int64     `json:"created_at" `          //
	UpdatedAt   int64     `json:"updated_at" `          //
}

const TravelCacheKey = "travels_pk_%s"

func (Travel) TableName() string {
	return "travels"
}

// BatchUpsert 批量插入或更新
func (r *TravelRepo) BatchUpsert(insertSlice ...*Travel) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":          db.Raw("values(slug)"),
				"area_id":       db.Raw("values(area_id)"),
				"user_id":       db.Raw("values(user_id)"),
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"status":        db.Raw("values(status)"),
				"start_date":    db.Raw("values(start_date)"),
				"end_date":      db.Raw("values(end_date)"),
				"days":          db.Raw("values(days)"),
				"total_price":   db.Raw("values(total_price)"),
				"plan_id":       db.Raw("values(plan_id)"),
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

func (r *TravelRepo) BatchInsert(insertSlice ...*Travel) (int64, error) {
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
func (r *TravelRepo) Insert(insert *Travel) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TravelRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"area_id",
		"user_id",
		"tour_guide_id",
		"status",
		"start_date",
		"end_date",
		"days",
		"total_price",
		"plan_id",
		"created_at",
		"updated_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TravelRepo) Omit(filter []string) []string {
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
func (r *TravelRepo) FindById(id int64) (Travel, error) {
	var travel Travel
	res := r.db.Where("id = ?", id).First(&travel)
	return travel, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TravelRepo) FindByIdArr(idArr []int64) []Travel {
	var travelArr = make([]Travel, 0)
	r.db.Where("id IN ?", idArr).Find(&travelArr)
	return travelArr
}

// FindBySlug 根据Slug进行查询
func (r *TravelRepo) FindBySlug(slug string) (Travel, error) {
	var travel Travel
	res := r.db.Where("slug = ?", slug).First(&travel)
	return travel, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *TravelRepo) FindBySlugArr(arr []string) ([]Travel, error) {
	var travelArr []Travel
	res := r.db.Where("slug IN ?", arr).Find(&travelArr)
	return travelArr, res.Error
}
