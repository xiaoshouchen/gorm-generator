//go:gen DON'T EDIT !
package model

import (
	"gorm.io/gorm/clause"
)

type Message struct {
	Id        int64   `json:"id" gorm:"primaryKey"` //
	Content   *string `json:"content" `             //
	MsgType   *string `json:"msg_type" `            //
	UserId    *int64  `json:"user_id" `             //
	CreatedAt *int64  `json:"created_at" `          //
	UpdatedAt *int64  `json:"updated_at" `          //
	IsRead    *int64  `json:"is_read" `             //
}

const MessageCacheKey = "messages_pk_%s"

func (Message) TableName() string {
	return "messages"
}

// BatchUpsert 批量插入或更新
func (r *MessageRepo) BatchUpsert(insertSlice ...*Message) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"content":    db.Raw("values(content)"),
				"msg_type":   db.Raw("values(msg_type)"),
				"user_id":    db.Raw("values(user_id)"),
				"created_at": db.Raw("values(created_at)"),
				"updated_at": db.Raw("values(updated_at)"),
				"is_read":    db.Raw("values(is_read)"),
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

func (r *MessageRepo) BatchInsert(insertSlice ...*Message) (int64, error) {
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
func (r *MessageRepo) Insert(insert *Message) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *MessageRepo) getAllFields() []string {
	return []string{
		"id",
		"content",
		"msg_type",
		"user_id",
		"created_at",
		"updated_at",
		"is_read",
	}
}

// Omit 过滤自己不想要的字段
func (r *MessageRepo) Omit(filter []string) []string {
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
func (r *MessageRepo) FindById(id int64) (Message, error) {
	var message Message
	res := r.db.Where("id = ?", id).First(&message)
	return message, res.Error
}

// FindByIdArr 根据主键获取数据
func (r *MessageRepo) FindByIdArr(idArr []int64) []Message {
	var messageArr = make([]Message, 0)
	r.db.Where("id IN ?", idArr).Find(&messageArr)
	return messageArr
}
