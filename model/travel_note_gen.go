//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type TravelNote struct {
	Id        int64        `json:"id" gorm:"primaryKey"` //
	UserId    *int64       `json:"user_id" `             //
	Bills     *interface{} `json:"bills" `               //
	CreatedAt *int64       `json:"created_at" `          //
	UpdatedAt *int64       `json:"updated_at" `          //
	DeletedAt *int64       `json:"deleted_at" `          //
}

const TravelNoteCacheKey = "travel_notes_pk_%s"

func (TravelNote) TableName() string {
	return "travel_notes"
}

// BatchUpsert 批量插入或更新
func (r *TravelNoteRepo) BatchUpsert(insertSlice ...*TravelNote) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"user_id":    db.Raw("values(user_id)"),
				"bills":      db.Raw("values(bills)"),
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

func (r *TravelNoteRepo) BatchInsert(insertSlice ...*TravelNote) (int64, error) {
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
func (r *TravelNoteRepo) Insert(insert *TravelNote) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *TravelNoteRepo) getAllFields() []string {
	return []string{
		"id",
		"user_id",
		"bills",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// Omit 过滤自己不想要的字段
func (r *TravelNoteRepo) Omit(filter []string) []string {
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
func (r *TravelNoteRepo) FindById(id int64) (TravelNote, error) {
	var travelNote TravelNote
	res := r.db.Where("id = ?", id).First(&travelNote)
	return travelNote, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *TravelNoteRepo) FindByIdArr(idArr []int64) []TravelNote {
	var travelNoteArr = make([]TravelNote, 0)
	r.db.Where("id IN ?", idArr).Find(&travelNoteArr)
	return travelNoteArr
}

// FindByUserId 根据UserId进行查询
func (r *TravelNoteRepo) FindByUserId(userId *int64, limit int, orderBy string) ([]TravelNote, error) {
	var travelNoteArr []TravelNote
	res := r.db.Where("user_id = ?", userId).Limit(limit).Order(orderBy).Find(&travelNoteArr)
	return travelNoteArr, res.Error
}
