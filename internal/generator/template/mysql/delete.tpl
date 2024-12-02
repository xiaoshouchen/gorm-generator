{{define "delete" }}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) Delete(ctx context.Context, {{- range $index,$value:= .Columns}}{{- if ne $value.ColumnKey "PRI"}}{{- end}}{{- end}} int64) error {
    return nil
    }
{{end}}