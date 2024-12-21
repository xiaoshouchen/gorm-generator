//go:gen DON'T EDIT !
package model

import (
{{if withCache}}
    "fmt"
    "panda-trip/pkg/cache"
{{end}}
{{if hasTime}}"time"{{end}}

"gorm.io/gorm"
"gorm.io/gorm/clause"
{{if eq softDelete 1}}"gorm.io/plugin/soft_delete"{{end}}
)

{{ $modelName :=.TableName | singular | upCamel }}
type {{$modelName}} struct {
{{range $k,$v := .Columns}}{{template "field" $v}}{{end}}
}

const {{$modelName}}CacheKey="{{.TableName}}_pk_%s"

func ({{$modelName}}) TableName() string {
return "{{.TableName}}"
}

func (r *{{$modelName}}Repo) DB() *gorm.DB {
return r.db
}

{{template "insert" .}}
{{template "omit" .}}
{{template "find" .}}
{{template "count" .}}
{{template "delete" .}}
