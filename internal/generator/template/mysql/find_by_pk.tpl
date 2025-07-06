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
    // 先尝试从缓存中获取
    cacheObj := cache.GetOnce().Get(key)
    if err := cacheObj.Json(&{{$varName}}); err != nil {
    // 如果缓存中没有，从数据库中获取
    res :=r.db.Where("{{pkWhereCondition}}", {{pkWhereArgsStr}}).First(&{{$varName}})
    if res.Error != nil {
        if res.Error == gorm.ErrRecordNotFound {
				cache.GetOnce().Set(key, "{}", 86400)
				return {{$varName}}, res.Error
			}
    return {{$varName}},res.Error
    }
    _=cache.GetOnce().Set(key, {{$varName}}, {{cacheTtl}})
    return {{$varName}},nil
    }else{
        var str string
		cacheObj.String(&str)
		if str == "{}" {
			return {{$varName}}, gorm.ErrRecordNotFound
		}
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
    for k, v := range resultList {
    var item {{$modelName}}
    if err1 := v.Json(&item); err1 != nil {
    continue
    } else {
        cacheKeyMap[k] = true
        var str string
        v.String(&str)
        if str == "{}" {
            continue
        }
        {{$varName}}Arr = append({{$varName}}Arr, item)
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

   var leftPkArr []{{pksType}}
	for _, pk := range {{pksParams}}Arr {
		if hasCache, ok := cacheKeyMap[fmt.Sprintf({{$modelName}}CacheKey, formatInterface(pk))]; ok {
            if !hasCache {
                leftPkArr = append(leftPkArr, pk)
            }
		}
	}

	var left{{$modelName}}Arr []{{$modelName}}
	if len(leftPkArr) > 0 {
		r.db.Where("{{pksWhereCondition}}", leftPkArr).Find(&left{{$modelName}}Arr)
	}
	{{$varName}}Arr = append({{$varName}}Arr, left{{$modelName}}Arr...)
    cacheMap := make(map[string]interface{})
    for _, item := range left{{$modelName}}Arr {
    cacheKey := fmt.Sprintf({{$modelName}}CacheKey, formatInterface([]interface{} { {{"item"|pksFields}} }))
    cacheMap[cacheKey] = item
    cacheKeyMap[cacheKey] = true
    }
    for cacheKey, hasCache := range cacheKeyMap {
        if !hasCache {
            cacheMap[cacheKey] = "{}"
        }
    }
    _ = cache.GetOnce().MultiSet(cacheMap,{{cacheTtl}})
    return {{$varName}}Arr
    }
{{end}}

