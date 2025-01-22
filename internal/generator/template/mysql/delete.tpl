{{define "delete" }}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) Delete({{pkParams}})(int64, error) {
        res:= r.db.Where("{{pkWhereCondition}}", {{pkWhereArgsStr}}).Delete(&{{$modelName}}{})
        return res.RowsAffected, res.Error
    }
{{end}}