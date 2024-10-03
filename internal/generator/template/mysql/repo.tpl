package model

import "gorm.io/gorm"

{{ $modelName :=.TableName | singular | upCamel }}
type {{$modelName}}Repo struct {
    db *gorm.DB
}

func New{{$modelName}}Repo(db *gorm.DB) *{{$modelName}}Repo {
    return &{{$modelName}}Repo{db: db}
}