{{define "insert"}}
{{ $modelName := .TableName | singular | upCamel }}
// BatchUpsert 批量插入或更新
func (r *{{$modelName}}Repo) BatchUpsert(ctx context.Context,insertSlice ...*{{$modelName}})(int64, error) {
	db := r.db.WithContext(ctx)

		db = db.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
{{- range $index,$value:= .Columns}}
	{{- if eq $value.ColumnKey "PRI"}}
					{Name: "{{$value.ColumnName}}"},
	{{- end}}
{{- end}}
				},
				DoUpdates: clause.Assignments(map[string]interface{}{
{{- range $index,$value:= .Columns}}
	{{- if ne $value.ColumnKey "PRI"}}
		"{{$value.ColumnName}}": db.Raw("values({{$value.ColumnName}})"),
	{{- end}}
{{- end}}
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

func (r *{{$modelName}}Repo) BatchInsert(ctx context.Context,insertSlice ...*{{$modelName}})(int64, error) {
db := r.db.WithContext(ctx)
if len(insertSlice) > 1000 {
db = db.CreateInBatches(insertSlice, 1000)

} else {
db = db.Create(insertSlice)
}
return db.RowsAffected, db.Error
}

// Insert 插入单个
// return id
func (r *{{$modelName}}Repo) Insert(ctx context.Context,insert *{{$modelName}})error {
db := r.db.WithContext(ctx)
db = db.Create(insert)
return db.Error
}

{{end}}