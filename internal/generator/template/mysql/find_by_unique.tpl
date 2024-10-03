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


{{define "findByUniqueWithCache"}}
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