{{define "findByPk"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{ $varName := .TableName | singular | lowCamel }}
    // FindBy{{pkFuncName}} 根据主键进行查询
    func (r *{{$modelName}}Repo) FindBy{{pkFuncName}}({{pkParams}}) ({{$modelName}},error) {
    var {{$varName}} {{$modelName}}
    res:= r.db.Where("{{pkWhereCondition}}", {{pkWhereArgsStr}}).First(&{{$varName}})
    return {{$varName}},res.Error
    }
{{end}}

{{define "findByPks"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{ $varName := .TableName | singular | lowCamel }}
    // FindBy{{pkFuncName}}Arr 根据主键获取数据
    func (r *{{$modelName}}Repo) FindBy{{pkFuncName}}Arr({{pksParams}}Arr []{{pksType}}) []{{$modelName}} {
    var {{$varName}}Arr = make([]{{$modelName}}, 0)
    r.db.Where("{{pksWhereCondition}}", {{pksParams}}Arr).Find(&{{$varName}}Arr)
    return {{$varName}}Arr
    }
{{end}}

{{define "findByPkWithCache"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{ $varName := .TableName | singular | lowCamel }}
    // FindBy{{pkFuncName}} 根据主键进行查询
    func (r *{{$modelName}}Repo) FindBy{{pkFuncName}}({{pkParams}}) ({{$modelName}},error) {
    var {{$varName}} {{$modelName}}
    var key = fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{}{ {{pkWhereArgsStr}} }))
    if err := cache.GetOnce().Get(key).Json(&{{$varName}}); err != nil {
    res :=r.db.Where("{{pkWhereCondition}}", {{pkWhereArgsStr}}).First(&{{$varName}})
    if res.Error != nil {
    return {{$varName}},res.Error
    }
    _=cache.GetOnce().Set(key, res, {{cacheTtl}})
        return {{$varName}},nil
    }else{
    return {{$varName}},nil
    }
    }
{{end}}

{{define "findByPksWithCache"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{ $varName := .TableName | singular | lowCamel }}
    // FindBy{{pkFuncName}}Arr 根据主键获取数据
    func (r *{{$modelName}}Repo) FindBy{{pkFuncName}}Arr({{pksParams}}Arr []{{pksType}}) []{{$modelName}} {
    var {{$varName}}Arr = make([]{{$modelName}}, 0)
    var cacheKeyList []string
    var cacheKeyMap = make(map[string]bool)
    for _, pk := range {{pksParams}}Arr {
    cacheKey := fmt.Sprintf({{$modelName}}CacheKey, formatInterface(pk))
    cacheKeyList = append(cacheKeyList, cacheKey)
    cacheKeyMap[cacheKey] = false
    }
    resultList, err := cache.GetOnce().MultiGet(cacheKeyList)
    if err == nil {
    for _, v := range resultList {
    var item {{$modelName}}
    if err1 := v.Json(&item); err1 != nil {
    continue
    } else {
    {{$varName}}Arr = append({{$varName}}Arr, item)
    cacheKeyMap[fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{} { {{"item"|pksFields}} }))] = true
    }
    }
    }

    var isReturn = true
    for _, hasCache := range cacheKeyMap {
    isReturn = isReturn && hasCache
    }
    if isReturn {
    return {{$varName}}Arr
    }

    r.db.Where("{{pksWhereCondition}}", {{pksParams}}Arr).Find(&{{$varName}}Arr)
    cacheMap := make(map[string]interface{})
    for _, item := range {{$varName}}Arr {
    cacheKey := fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{} { {{"item"|pksFields}} }))
    cacheMap[cacheKey] = item
    }
    _ = cache.GetOnce().MultiSet(cacheMap,{{cacheTtl}})
    return {{$varName}}Arr
    }
{{end}}

