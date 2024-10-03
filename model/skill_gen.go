//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Skill struct {
	Id        int64                 `json:"id" gorm:"primaryKey"` //
	Slug      string                `json:"slug" `                //
	Name      string                `json:"name" `                //
	CreatedAt int64                 `json:"created_at" `          //
	UpdatedAt int64                 `json:"updated_at" `          //
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" `          //
}

const SkillCacheKey = "skills_pk_%s"

func (Skill) TableName() string {
	return "skills"
}

// BatchUpsert 批量插入或更新
func (r *SkillRepo) BatchUpsert(insertSlice ...*Skill) (int64, error) {
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

func (r *SkillRepo) BatchInsert(insertSlice ...*Skill) (int64, error) {
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
func (r *SkillRepo) Insert(insert *Skill) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *SkillRepo) getAllFields() []string {
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
func (r *SkillRepo) Omit(filter []string) []string {
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
func (r *SkillRepo) FindById(id int64) (Skill, error) {
	var skill Skill
	res := r.db.Where("id = ?", id).First(&skill)
	return skill, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *SkillRepo) FindByIdArr(idArr []int64) []Skill {
	var skillArr = make([]Skill, 0)
	r.db.Where("id IN ?", idArr).Find(&skillArr)
	return skillArr
}

// FindBySlug 根据Slug进行查询
func (r *SkillRepo) FindBySlug(slug string) (Skill, error) {
	var skill Skill
	res := r.db.Where("slug = ?", slug).First(&skill)
	return skill, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *SkillRepo) FindBySlugArr(arr []string) ([]Skill, error) {
	var skillArr []Skill
	res := r.db.Where("slug IN ?", arr).Find(&skillArr)
	return skillArr, res.Error
}
