{{define "delete" }}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) Delete({{pkParams}})(int64, error) {
        res:= r.db.Where("{{pkWhereCondition}}", {{pkWhereArgsStr}}).Delete(&{{$modelName}}{})
        return res.RowsAffected, res.Error
    }

    
    {{if withCache}}
        {{template "delete_cache" .}}
    {{end}}

{{end}}

{{define "delete_cache" }}
    {{ $modelName := .TableName | singular | upCamel }}
    func (r *{{$modelName}}Repo) DeleteCacheBy{{pkFuncName}}({{pkParams}}) error {
        cacheKey := fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{}{ {{pkWhereArgsStr}} } ))
        err := cache.GetOnce().Del(cacheKey)
        return err
    }

    func (r *{{$modelName}}Repo) DeleteCacheBy{{pkFuncName}}Arr(arr []{{pksType}}) error {
        for _, v := range arr {
            cacheKey := fmt.Sprintf({{$modelName}}CacheKey, formatInterface(v))
            err := cache.GetOnce().Del(cacheKey)
            if err != nil {
                return err
            }
        }
        return nil
    }
{{end}}
    