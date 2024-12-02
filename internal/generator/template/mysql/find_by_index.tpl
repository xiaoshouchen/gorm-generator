{{define "findByIndex"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{$varName := .TableName | singular | lowCamel }}
    {{range $k,$v := indexes}}
        {{$funcName := indexFuncName $v}}
        // FindBy{{$funcName}} 根据{{$funcName}}进行查询
        func (r *{{$modelName}}Repo) FindBy{{$funcName}}(ctx context.Context,{{indexParams $v}},limit int,orderBy string) ([]{{$modelName}},error) {
        var {{$varName}}Arr []{{$modelName}}
        res:= r.db.WithContext(ctx).Where("{{indexWhereCondition $v}}", {{indexWhereArgs $v}}).Limit(limit).Order(orderBy).Find(&{{$varName}}Arr)
        return {{$varName}}Arr,res.Error
        }
    {{end}}

    func (r *{{$modelName}}Repo) FindList(ctx context.Context, where []Where, page, size int64, orderBy string) ([]{{$modelName}},error) {
        var results []{{$modelName}}
        db := r.db.WithContext(ctx)
        for _, w := range where {
            db = db.Model(&{{$modelName}}{}).Where(w.Key+" "+w.Value.Op+" ?", w.Value.Arg)
        }
        if orderBy != "" {
            db = db.Order(orderBy)
        }
        if page > 0 && size > 0 {
            db = db.Limit(int(size)).Offset(int((page - 1) * size))
        }
        res := db.Find(&results)
        return results, res.Error
    }
{{end}}