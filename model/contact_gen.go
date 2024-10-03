//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type Contact struct {
	Id        int64                 `json:"id" gorm:"primaryKey"` //
	Slug      string                `json:"slug" `                //
	UserId    int64                 `json:"user_id" `             //
	Name      string                `json:"name" `                //
	AreaCode  int64                 `json:"area_code" `           //
	Phone     string                `json:"phone" `               //
	Email     string                `json:"email" `               //
	Wechat    string                `json:"wechat" `              //
	CreatedAt int64                 `json:"created_at" `          //
	UpdatedAt int64                 `json:"updated_at" `          //
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" `          //
	IsDefault int64                 `json:"is_default" `          //
}

const ContactCacheKey = "contacts_pk_%s"

func (Contact) TableName() string {
	return "contacts"
}

// BatchUpsert 批量插入或更新
func (r *ContactRepo) BatchUpsert(insertSlice ...*Contact) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"slug":       db.Raw("values(slug)"),
				"user_id":    db.Raw("values(user_id)"),
				"name":       db.Raw("values(name)"),
				"area_code":  db.Raw("values(area_code)"),
				"phone":      db.Raw("values(phone)"),
				"email":      db.Raw("values(email)"),
				"wechat":     db.Raw("values(wechat)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
				"deleted_at": db.Raw("values(deleted_at)"),
				"is_default": db.Raw("values(is_default)"),
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

func (r *ContactRepo) BatchInsert(insertSlice ...*Contact) (int64, error) {
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
func (r *ContactRepo) Insert(insert *Contact) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *ContactRepo) getAllFields() []string {
	return []string{
		"id",
		"slug",
		"user_id",
		"name",
		"area_code",
		"phone",
		"email",
		"wechat",
		"created_at",
		"updated_at",
		"deleted_at",
		"is_default",
	}
}

// Omit 过滤自己不想要的字段
func (r *ContactRepo) Omit(filter []string) []string {
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
func (r *ContactRepo) FindById(id int64) (Contact, error) {
	var contact Contact
	res := r.db.Where("id = ?", id).First(&contact)
	return contact, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *ContactRepo) FindByIdArr(idArr []int64) []Contact {
	var contactArr = make([]Contact, 0)
	r.db.Where("id IN ?", idArr).Find(&contactArr)
	return contactArr
}

// FindBySlug 根据Slug进行查询
func (r *ContactRepo) FindBySlug(slug string) (Contact, error) {
	var contact Contact
	res := r.db.Where("slug = ?", slug).First(&contact)
	return contact, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *ContactRepo) FindBySlugArr(arr []string) ([]Contact, error) {
	var contactArr []Contact
	res := r.db.Where("slug IN ?", arr).Find(&contactArr)
	return contactArr, res.Error
}

// FindByUserId 根据UserId进行查询
func (r *ContactRepo) FindByUserId(userId int64, limit int, orderBy string) ([]Contact, error) {
	var contactArr []Contact
	res := r.db.Where("user_id = ?", userId).Limit(limit).Order(orderBy).Find(&contactArr)
	return contactArr, res.Error
}
