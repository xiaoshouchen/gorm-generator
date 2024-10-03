//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type TourSpot struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	Slug        string                `json:"slug" `                //
	AreaId      int64                 `json:"area_id" `             //
	Name        string                `json:"name" `                //
	Description string                `json:"description" `         //
	Images      string                `json:"images" `              //
	OpenTime    string                `json:"open_time" `           //
	CloseTime   string                `json:"close_time" `          //
	TicketPrice float64               `json:"ticket_price" `        //
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
	Location    string                `json:"location" `            //
	Latitude    float64               `json:"latitude" `            //
	Longitude   float64               `json:"longitude" `           //
}

const TourSpotCacheKey = "tour_spots_pk_%s"

func (TourSpot) TableName() string {
	return "tour_spots"
}

// BatchUpsert 批量插入或更新
func (r *TourSpotRepo) BatchUpsert(insertSlice ...*TourSpot) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":         db.Raw("values(slug)"),
				"area_id":      db.Raw("values(area_id)"),
				"name":         db.Raw("values(name)"),
				"description":  db.Raw("values(description)"),
				"images":       db.Raw("values(images)"),
				"open_time":    db.Raw("values(open_time)"),
				"close_time":   db.Raw("values(close_time)"),
				"ticket_price": db.Raw("values(ticket_price)"),
				"created_at":   db.Raw("values(created_at)"),
				"updated_at":   db.Raw("values(updated_at)"),
				"deleted_at":   db.Raw("values(deleted_at)"),
				"location":     db.Raw("values(location)"),
				"latitude":     db.Raw("values(latitude)"),
				"longitude":    db.Raw("values(longitude)"),
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

func (r *TourSpotRepo) BatchInsert(insertSlice ...*TourSpot) (int64, error) {
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
func (r *TourSpotRepo) Insert(insert *TourSpot) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TourSpotRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"area_id",
		"name",
		"description",
		"images",
		"open_time",
		"close_time",
		"ticket_price",
		"created_at",
		"updated_at",
		"deleted_at",
		"location",
		"latitude",
		"longitude",
	}
}

// Omit 过滤自己不想要的字段
func (r *TourSpotRepo) Omit(filter []string) []string {
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
func (r *TourSpotRepo) FindById(id int64) (TourSpot, error) {
	var tourSpot TourSpot
	res := r.db.Where("id = ?", id).First(&tourSpot)
	return tourSpot, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TourSpotRepo) FindByIdArr(idArr []int64) []TourSpot {
	var tourSpotArr = make([]TourSpot, 0)
	r.db.Where("id IN ?", idArr).Find(&tourSpotArr)
	return tourSpotArr
}

// FindBySlug 根据Slug进行查询
func (r *TourSpotRepo) FindBySlug(slug string) (TourSpot, error) {
	var tourSpot TourSpot
	res := r.db.Where("slug = ?", slug).First(&tourSpot)
	return tourSpot, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *TourSpotRepo) FindBySlugArr(arr []string) ([]TourSpot, error) {
	var tourSpotArr []TourSpot
	res := r.db.Where("slug IN ?", arr).Find(&tourSpotArr)
	return tourSpotArr, res.Error
}

// FindByAreaId 根据AreaId进行查询
func (r *TourSpotRepo) FindByAreaId(areaId int64, limit int, orderBy string) ([]TourSpot, error) {
	var tourSpotArr []TourSpot
	res := r.db.Where("area_id = ?", areaId).Limit(limit).Order(orderBy).Find(&tourSpotArr)
	return tourSpotArr, res.Error
}
