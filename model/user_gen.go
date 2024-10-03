//go:gen DON'T EDIT !
package model

import (
	"fmt"
	"panda-trip/pkg/cache"

	"gorm.io/gorm/clause"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id                int64                 `json:"id" gorm:"primaryKey"` //
	Email             *string               `json:"email" `               //手机、邮箱二选一
	AreaCode          *int64                `json:"area_code" `           //手机不同国家的前缀编码
	Password          string                `json:"password" `            //
	Phone             *string               `json:"phone" `               //
	Slug              string                `json:"slug" `                //用户的唯一编码
	CreatedAt         int64                 `json:"created_at" `          //
	UpdatedAt         int64                 `json:"updated_at" `          //
	DeletedAt         soft_delete.DeletedAt `json:"deleted_at" `          //
	PasswordChangedAt int64                 `json:"password_changed_at" ` //用户最近一次修改密码的时间，方便jwt进行失效验证
	Nickname          string                `json:"nickname" `            //
	Avatar            string                `json:"avatar" `              //
}

const UserCacheKey = "users_pk_%s"

func (User) TableName() string {
	return "users"
}

// BatchUpsert 批量插入或更新
func (r *UserRepo) BatchUpsert(insertSlice ...*User) (int64, error) {
	db := r.db

	db = db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"email":               db.Raw("values(email)"),
				"area_code":           db.Raw("values(area_code)"),
				"password":            db.Raw("values(password)"),
				"phone":               db.Raw("values(phone)"),
				"slug":                db.Raw("values(slug)"),
				"created_at":          db.Raw("values(created_at)"),
				"updated_at":          db.Raw("values(updated_at)"),
				"deleted_at":          db.Raw("values(deleted_at)"),
				"password_changed_at": db.Raw("values(password_changed_at)"),
				"nickname":            db.Raw("values(nickname)"),
				"avatar":              db.Raw("values(avatar)"),
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

func (r *UserRepo) BatchInsert(insertSlice ...*User) (int64, error) {
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
func (r *UserRepo) Insert(insert *User) error {
	db := r.db
	db = db.Create(insert)
	return db.Error
}

func (r *UserRepo) getAllFields() []string {
	return []string{
		"id",
		"email",
		"area_code",
		"password",
		"phone",
		"slug",
		"created_at",
		"updated_at",
		"deleted_at",
		"password_changed_at",
		"nickname",
		"avatar",
	}
}

// Omit 过滤自己不想要的字段
func (r *UserRepo) Omit(filter []string) []string {
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
func (r *UserRepo) FindById(id int64) (User, error) {
	var user User
	var key = fmt.Sprintf("users_pk_%s", formatInterface([]interface{}{id}))
	if err := cache.GetOnce().Get(key).Json(&user); err != nil {
		res := r.db.Where("id = ?", id).First(&user)
		if res.Error != nil {
			return user, res.Error
		}
		_ = cache.GetOnce().Set(key, res, 0)
		return user, nil
	} else {
		return user, nil
	}
}

// FindByIdArr 根据主键获取数据
func (r *UserRepo) FindByIdArr(idArr []int64) []User {
	var userArr = make([]User, 0)
	var cacheKeyList []string
	var cacheKeyMap = make(map[string]bool)
	for _, pk := range idArr {
		cacheKey := fmt.Sprintf(UserCacheKey, formatInterface(pk))
		cacheKeyList = append(cacheKeyList, cacheKey)
		cacheKeyMap[cacheKey] = false
	}
	resultList, err := cache.GetOnce().MultiGet(cacheKeyList)
	if err == nil {
		for _, v := range resultList {
			var item User
			if err1 := v.Json(&item); err1 != nil {
				continue
			} else {
				userArr = append(userArr, item)
				cacheKeyMap[formatInterface([]interface{}{item.Id})] = true
			}
		}
	}

	var isReturn = true
	for _, hasCache := range cacheKeyMap {
		isReturn = isReturn && hasCache
	}
	if isReturn {
		return userArr
	}

	r.db.Where("id IN ?", idArr).Find(&userArr)
	cacheMap := make(map[string]interface{})
	for _, item := range userArr {
		cacheKey := fmt.Sprintf(UserCacheKey, formatInterface([]interface{}{item.Id}))
		cacheMap[cacheKey] = item
	}
	_ = cache.GetOnce().MultiSet(cacheMap, 0)
	return userArr
}

// FindByEmail 根据Email进行查询
func (r *UserRepo) FindByEmail(email *string) (User, error) {
	var user User
	res := r.db.Where("email = ?", email).First(&user)
	return user, res.Error
}

// FindByAreaCodeAndPhone 根据AreaCodeAndPhone进行查询
func (r *UserRepo) FindByAreaCodeAndPhone(areaCode *int64, phone *string) (User, error) {
	var user User
	res := r.db.Where("area_code = ? AND phone = ?", areaCode, phone).First(&user)
	return user, res.Error
}

// FindBySlug 根据Slug进行查询
func (r *UserRepo) FindBySlug(slug string) (User, error) {
	var user User
	res := r.db.Where("slug = ?", slug).First(&user)
	return user, res.Error
}

// FindByEmailArr 根据主键Email进行查询
func (r *UserRepo) FindByEmailArr(arr []*string) ([]User, error) {
	var userArr []User
	res := r.db.Where("email IN ?", arr).Find(&userArr)
	return userArr, res.Error
}

// FindByAreaCodeAndPhoneArr 根据主键AreaCodeAndPhone进行查询
func (r *UserRepo) FindByAreaCodeAndPhoneArr(arr []interface{}) ([]User, error) {
	var userArr []User
	res := r.db.Where("(area_code,phone) IN ?", arr).Find(&userArr)
	return userArr, res.Error
}

// FindBySlugArr 根据主键Slug进行查询
func (r *UserRepo) FindBySlugArr(arr []string) ([]User, error) {
	var userArr []User
	res := r.db.Where("slug IN ?", arr).Find(&userArr)
	return userArr, res.Error
}

// FindByCreatedAt 根据CreatedAt进行查询
func (r *UserRepo) FindByCreatedAt(createdAt int64, limit int, orderBy string) ([]User, error) {
	var userArr []User
	res := r.db.Where("created_at = ?", createdAt).Limit(limit).Order(orderBy).Find(&userArr)
	return userArr, res.Error
}
