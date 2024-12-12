{{define "count"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{$varName := .TableName | singular | lowCamel }}
    func (r *{{$modelName}}Repo) Count(where []Where) (int64,error) {
        var count int64
        db := r.db.Model(&{{$modelName}}{})
        for _, w := range where {
            db = db.Where(w.Key+" "+w.Value.Op+" ?", w.Value.Arg)
        }
        res := db.Count(&count)
        return count, res.Error
    }
{{end}}