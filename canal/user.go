package canal

import (
	"panda-trip/internal/model"
)

type User struct {
	model.UserCache
}

func NewUser() *User {
	return &User{
		UserCache: model.UserCache{},
	}
}
