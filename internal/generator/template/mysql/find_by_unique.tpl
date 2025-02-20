{{define "findByUnique"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{$varName := .TableName | singular | lowCamel }}
    {{range $k,$v := uniques}}
        {{$funcName := uniqueFuncName $v}}
        // FindBy{{$funcName}} 根据{{$funcName}}进行查询
        func (r *{{$modelName}}Repo) FindBy{{$funcName}}({{uniqueParams $v}}) ({{$modelName}},error) {
        var {{$varName}} {{$modelName}}
        res:= r.db.Where("{{uniqueWhereCondition $v}}", {{uniqueWhereArgs $v}}).First(&{{$varName}})
        return {{$varName}},res.Error
        }
    {{end}}
{{end}}

{{define "findByUniqueWithCache"}}
    {{ $tableName := .TableName }}
    {{ $modelName := .TableName | singular | upCamel }}
    {{$varName := .TableName | singular | lowCamel }}
    {{range $k,$v := uniques}}
        {{$funcName := uniqueFuncName $v}}
        // FindBy{{$funcName}} 根据{{$funcName}}进行查询
        func (r *{{$modelName}}Repo) FindBy{{$funcName}}({{uniqueParams $v}}) ({{$modelName}},error) {
           var {{$varName}} {{$modelName}}
           // Try to get id from unique field cache
           uniqueKey := fmt.Sprintf("{{$tableName}}_{{$funcName | snake}}_pk_{{uniqueCountParams $v}}", {{uniqueWhereArgs $v}})
           var pk string
           if err := cache.GetOnce().Get(uniqueKey).String(&pk); err == nil {
              if err := cache.GetOnce().Get(pk).Json(&{{$varName}}); err == nil {
                  return {{$varName}}, nil
              }
           }

           // If no cached id found, query from database
           res := r.db.Where("{{uniqueWhereCondition $v}}", {{uniqueWhereArgs $v}}).First(&{{$varName}})
           if res.Error != nil {
               return {{$varName}}, res.Error
           }
           // Cache both the unique->id mapping and the data
           idKey:=fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{}{ {{$varName|pksFields}} }))
           _ = cache.GetOnce().Set(uniqueKey, idKey, {{cacheTtl}})
           _ = cache.GetOnce().Set(idKey, {{$varName}}, {{cacheTtl}})
           return {{$varName}}, nil
        }
    {{end}}
{{end}}

{{define "findByUniques"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{ $varName := .TableName | singular | lowCamel }}
    {{range $k,$v := uniques}}
        {{$funcName := uniqueFuncName $v}}
        // FindBy{{$funcName}}Arr 根据主键{{$funcName}}进行查询
        func (r *{{$modelName}}Repo) FindBy{{$funcName}}Arr(arr []{{uniquesType $v}}) ([]{{$modelName}},error) {
        var {{$varName}}Arr []{{$modelName}}
        res :=r.db.Where("{{uniquesWhereCondition $v}}", arr).Find(&{{$varName}}Arr)
        return {{$varName}}Arr,res.Error
        }
    {{end}}
{{end}}