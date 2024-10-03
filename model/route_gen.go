//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Route struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	Slug        string                `json:"slug" `                //
	ObjectId    int64                 `json:"object_id" `           //关联的ID，比如关联的酒店，景点，餐饮等
	StartTime   string                `json:"start_time" `          //
	EndTime     string                `json:"end_time" `            //
	PlanId      int64                 `json:"plan_id" `             //对应的套餐服务的ID
	NextId      int64                 `json:"next_id" `             //下一个目的地
	HowTo       int64                 `json:"how_to" `              //如何到达，使用什么样子的交通工具
	ObjectType  int64                 `json:"object_type" `         //1 景点，2酒店 3 吃饭 4 其他
	Day         int64                 `json:"day" `                 //第几天
	Description string                `json:"description" `         //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const RouteCacheKey = "routes_pk_%s"

func (Route) TableName() string {
	return "routes"
}

// BatchUpsert 批量插入或更新
func (r *RouteRepo) BatchUpsert(insertSlice ...*Route) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":        db.Raw("values(slug)"),
				"object_id":   db.Raw("values(object_id)"),
				"start_time":  db.Raw("values(start_time)"),
				"end_time":    db.Raw("values(end_time)"),
				"plan_id":     db.Raw("values(plan_id)"),
				"next_id":     db.Raw("values(next_id)"),
				"how_to":      db.Raw("values(how_to)"),
				"object_type": db.Raw("values(object_type)"),
				"day":         db.Raw("values(day)"),
				"description": db.Raw("values(description)"),
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

func (r *RouteRepo) BatchInsert(insertSlice ...*Route) (int64, error) {
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
func (r *RouteRepo) Insert(insert *Route) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *RouteRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"object_id",
		"start_time",
		"end_time",
		"plan_id",
		"next_id",
		"how_to",
		"object_type",
		"day",
		"description",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *RouteRepo) Omit(filter []string) []string {
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
func (r *RouteRepo) FindById(id int64) (Route, error) {
	var route Route
	res := r.db.Where("id = ?", id).First(&route)
	return route, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *RouteRepo) FindByIdArr(idArr []int64) []Route {
	var routeArr = make([]Route, 0)
	r.db.Where("id IN ?", idArr).Find(&routeArr)
	return routeArr
}

// FindBySlug 根据Slug进行查询
func (r *RouteRepo) FindBySlug(slug string) (Route, error) {
	var route Route
	res := r.db.Where("slug = ?", slug).First(&route)
	return route, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *RouteRepo) FindBySlugArr(arr []string) ([]Route, error) {
	var routeArr []Route
	res := r.db.Where("slug IN ?", arr).Find(&routeArr)
	return routeArr, res.Error
}

// FindByPlanId 根据PlanId进行查询
func (r *RouteRepo) FindByPlanId(planId int64, limit int, orderBy string) ([]Route, error) {
	var routeArr []Route
	res := r.db.Where("plan_id = ?", planId).Limit(limit).Order(orderBy).Find(&routeArr)
	return routeArr, res.Error
}
