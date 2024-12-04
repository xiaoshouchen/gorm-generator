package model

import (
	"context"
	"gorm.io/gorm"
)

{{ $modelName :=.TableName | singular | upCamel }}
type {{$modelName}}Repo struct {
    db *gorm.DB
}

func New{{$modelName}}Repo(db *gorm.DB, ctxs ...context.Contextt) *{{$modelName}}Repo {
	if len(ctxs) > 0 {
		db = db.WithContext(ctxs[0])
	}
     return &{{$modelName}}Repo{db: db}
}
