//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type TourGuide struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	Slug        string                `json:"slug" `                //
	UserId      int64                 `json:"user_id" `             //
	Name        string                `json:"name" `                //导游姓名
	Sex         int64                 `json:"sex" `                 //1男2女
	Selfie      string                `json:"selfie" `              //自拍照
	WorkYears   int64                 `json:"work_years" `          //
	Description string                `json:"description" `         //个人简介
	Status      int64                 `json:"status" `              //状态，1草稿，2审核，3通过，4封禁
	Rate        int64                 `json:"rate" `                //
	ServiceFee  float64               `json:"service_fee" `         //服务费，每天的费用
	Images      string                `json:"images" `              //其他的生活照片
	Sort        int64                 `json:"sort" `                //排名
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
	WorkStartAt string                `json:"work_start_at" `       //
	WorkEndAt   string                `json:"work_end_at" `         //
}

const TourGuideCacheKey = "tour_guides_pk_%s"

func (TourGuide) TableName() string {
	return "tour_guides"
}

// BatchUpsert 批量插入或更新
func (r *TourGuideRepo) BatchUpsert(insertSlice ...*TourGuide) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":          db.Raw("values(slug)"),
				"user_id":       db.Raw("values(user_id)"),
				"name":          db.Raw("values(name)"),
				"sex":           db.Raw("values(sex)"),
				"selfie":        db.Raw("values(selfie)"),
				"work_years":    db.Raw("values(work_years)"),
				"description":   db.Raw("values(description)"),
				"status":        db.Raw("values(status)"),
				"rate":          db.Raw("values(rate)"),
				"service_fee":   db.Raw("values(service_fee)"),
				"images":        db.Raw("values(images)"),
				"sort":          db.Raw("values(sort)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
				"deleted_at":    db.Raw("values(deleted_at)"),
				"work_start_at": db.Raw("values(work_start_at)"),
				"work_end_at":   db.Raw("values(work_end_at)"),
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

func (r *TourGuideRepo) BatchInsert(insertSlice ...*TourGuide) (int64, error) {
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
func (r *TourGuideRepo) Insert(insert *TourGuide) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TourGuideRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"user_id",
		"name",
		"sex",
		"selfie",
		"work_years",
		"description",
		"status",
		"rate",
		"service_fee",
		"images",
		"sort",
		"created_at",
		"updated_at",
		"deleted_at",
		"work_start_at",
		"work_end_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TourGuideRepo) Omit(filter []string) []string {
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
func (r *TourGuideRepo) FindById(id int64) (TourGuide, error) {
	var tourGuide TourGuide
	res := r.db.Where("id = ?", id).First(&tourGuide)
	return tourGuide, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TourGuideRepo) FindByIdArr(idArr []int64) []TourGuide {
	var tourGuideArr = make([]TourGuide, 0)
	r.db.Where("id IN ?", idArr).Find(&tourGuideArr)
	return tourGuideArr
}

// FindBySlug 根据Slug进行查询
func (r *TourGuideRepo) FindBySlug(slug string) (TourGuide, error) {
	var tourGuide TourGuide
	res := r.db.Where("slug = ?", slug).First(&tourGuide)
	return tourGuide, res.Error
}

// FindByUserId 根据UserId进行查询
func (r *TourGuideRepo) FindByUserId(userId int64) (TourGuide, error) {
	var tourGuide TourGuide
	res := r.db.Where("user_id = ?", userId).First(&tourGuide)
	return tourGuide, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *TourGuideRepo) FindBySlugArr(arr []string) ([]TourGuide, error) {
	var tourGuideArr []TourGuide
	res := r.db.Where("slug IN ?", arr).Find(&tourGuideArr)
	return tourGuideArr, res.Error
}

// FindByUserIdArr 根据主键UserId进行查询
func (r *TourGuideRepo) FindByUserIdArr(arr []int64) ([]TourGuide, error) {
	var tourGuideArr []TourGuide
	res := r.db.Where("user_id IN ?", arr).Find(&tourGuideArr)
	return tourGuideArr, res.Error
}
