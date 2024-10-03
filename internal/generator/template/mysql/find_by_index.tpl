{{define "findByIndex"}}
    {{ $modelName := .TableName | singular | upCamel }}
    {{$varName := .TableName | singular | lowCamel }}
    {{range $k,$v := indexes}}
        {{$funcName := indexFuncName $v}}
        // FindBy{{$funcName}} 根据{{$funcName}}进行查询
        func (r *{{$modelName}}Repo) FindBy{{$funcName}}({{indexParams $v}},limit int,orderBy string) ([]{{$modelName}},error) {
        var {{$varName}}Arr []{{$modelName}}
        res:= r.db.Where("{{indexWhereCondition $v}}", {{indexWhereArgs $v}}).Limit(limit).Order(orderBy).Find(&{{$varName}}Arr)
        return {{$varName}}Arr,res.Error
        }
    {{end}}
{{end}}
