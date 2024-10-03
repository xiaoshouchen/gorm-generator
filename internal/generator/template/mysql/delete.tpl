{{define "delete" }}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) Delete({{- range $index,$value:= .Columns}}{{- if ne $value.ColumnKey "PRI"}}{{- end}}{{- end}})(int64, error) {
    return 0, nil
    }
{{end}}