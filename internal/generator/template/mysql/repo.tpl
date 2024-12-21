package model

import (
	"context"
	"gorm.io/gorm"
)

{{ $modelName :=.TableName | singular | upCamel }}
type {{$modelName}}Repo struct {
    db *gorm.DB
}

func New{{$modelName}}Repo(db *gorm.DB, ctxArr ...context.Context) *{{$modelName}}Repo {
	if len(ctxArr) > 0 {
		db = db.WithContext(ctxArr[0])
	}
     return &{{$modelName}}Repo{db: db}
}
