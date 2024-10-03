//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type TourGuideSkill struct {
	Id          int64                 `json:"id" gorm:"primaryKey"` //
	SkillType   int64                 `json:"skill_type" `          //1 语言 2 驾驶 3导游证
	Slug        string                `json:"slug" `                //
	TourGuideId int64                 `json:"tour_guide_id" `       //
	Name        string                `json:"name" `                //如果是语言，Engliash，如果是驾驶 C1
	Level       string                `json:"level" `               //
	Status      int64                 `json:"status" `              //1 未认证，2 已验证
	CreatedAt   int64                 `json:"created_at" `          //
	UpdatedAt   int64                 `json:"updated_at" `          //
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" `          //
}

const TourGuideSkillCacheKey = "tour_guide_skills_pk_%s"

func (TourGuideSkill) TableName() string {
	return "tour_guide_skills"
}

// BatchUpsert 批量插入或更新
func (r *TourGuideSkillRepo) BatchUpsert(insertSlice ...*TourGuideSkill) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"skill_type":    db.Raw("values(skill_type)"),
				"slug":          db.Raw("values(slug)"),
				"tour_guide_id": db.Raw("values(tour_guide_id)"),
				"name":          db.Raw("values(name)"),
				"level":         db.Raw("values(level)"),
				"status":        db.Raw("values(status)"),
				"created_at":    db.Raw("values(created_at)"),
				"updated_at":    db.Raw("values(updated_at)"),
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

func (r *TourGuideSkillRepo) BatchInsert(insertSlice ...*TourGuideSkill) (int64, error) {
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
func (r *TourGuideSkillRepo) Insert(insert *TourGuideSkill) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TourGuideSkillRepo) getAllFields() []string {
	return []string{
		"id",
		"skill_type",
		"slug",
		"tour_guide_id",
		"name",
		"level",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TourGuideSkillRepo) Omit(filter []string) []string {
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
func (r *TourGuideSkillRepo) FindById(id int64) (TourGuideSkill, error) {
	var tourGuideSkill TourGuideSkill
	res := r.db.Where("id = ?", id).First(&tourGuideSkill)
	return tourGuideSkill, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TourGuideSkillRepo) FindByIdArr(idArr []int64) []TourGuideSkill {
	var tourGuideSkillArr = make([]TourGuideSkill, 0)
	r.db.Where("id IN ?", idArr).Find(&tourGuideSkillArr)
	return tourGuideSkillArr
}

// FindBySlug 根据Slug进行查询
func (r *TourGuideSkillRepo) FindBySlug(slug string) (TourGuideSkill, error) {
	var tourGuideSkill TourGuideSkill
	res := r.db.Where("slug = ?", slug).First(&tourGuideSkill)
	return tourGuideSkill, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *TourGuideSkillRepo) FindBySlugArr(arr []string) ([]TourGuideSkill, error) {
	var tourGuideSkillArr []TourGuideSkill
	res := r.db.Where("slug IN ?", arr).Find(&tourGuideSkillArr)
	return tourGuideSkillArr, res.Error
}

// FindByTourGuideId 根据TourGuideId进行查询
func (r *TourGuideSkillRepo) FindByTourGuideId(tourGuideId int64, limit int, orderBy string) ([]TourGuideSkill, error) {
	var tourGuideSkillArr []TourGuideSkill
	res := r.db.Where("tour_guide_id = ?", tourGuideId).Limit(limit).Order(orderBy).Find(&tourGuideSkillArr)
	return tourGuideSkillArr, res.Error
}
