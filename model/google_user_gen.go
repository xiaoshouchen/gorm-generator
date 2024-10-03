//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type GoogleUser struct {
	Id     int64 `json:"id" gorm:"primaryKey"` //
	UserId int64 `json:"user_id" `             //
}

const GoogleUserCacheKey = "google_users_pk_%s"

func (GoogleUser) TableName() string {
	return "google_users"
}

// BatchUpsert 批量插入或更新
func (r *GoogleUserRepo) BatchUpsert(insertSlice ...*GoogleUser) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"user_id": db.Raw("values(user_id)"),
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

func (r *GoogleUserRepo) BatchInsert(insertSlice ...*GoogleUser) (int64, error) {
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
func (r *GoogleUserRepo) Insert(insert *GoogleUser) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *GoogleUserRepo) getAllFields() []string {
	return []string{
		"id",
		"user_id",
	}
}

// Omit 过滤自己不想要的字段
func (r *GoogleUserRepo) Omit(filter []string) []string {
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
func (r *GoogleUserRepo) FindById(id int64) (GoogleUser, error) {
	var googleUser GoogleUser
	res := r.db.Where("id = ?", id).First(&googleUser)
	return googleUser, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *GoogleUserRepo) FindByIdArr(idArr []int64) []GoogleUser {
	var googleUserArr = make([]GoogleUser, 0)
	r.db.Where("id IN ?", idArr).Find(&googleUserArr)
	return googleUserArr
}
