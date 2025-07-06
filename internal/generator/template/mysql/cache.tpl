package model

import (
	"panda-trip/pkg/custom_array"
	"panda-trip/pkg/db"

	"github.com/go-mysql-org/go-mysql/canal"
)

{{ $modelName :=.TableName | singular | upCamel }}
{{ $modelVarName :=.TableName | singular | lowCamel }}

type {{$modelName}}Cache struct {
	tool  *Tool
	event *canal.RowsEvent
}

func New{{$modelName}}Cache(tool *Tool, event *canal.RowsEvent) *{{$modelName}}Cache {
	return &{{$modelName}}Cache{
		tool:  tool,
		event: event,
	}
}

func (s *{{$modelName}}Cache) SetEvent(event *canal.RowsEvent) {
	s.event = event
}

func (s *{{$modelName}}Cache) Insert() error {
	// 获取返回的数组
	dataMap := s.tool.toMap(s.event)
	if len(dataMap) == 0 {
		return nil
	}
	// 插入
	var {{$modelVarName}}PkArr []{{pksType}}
	// 生成新的缓存
	for _, data := range dataMap {
		if id, ok := data["id"]; ok && custom_array.IsInteger(id) {
			{{$modelVarName}}PkArr = append({{$modelVarName}}PkArr, custom_array.ToInt64(id))
		}
	}

	// 删除旧缓存,防止空缓存的存在
	New{{$modelName}}Repo(db.Manager).DeleteCacheBy{{pkFuncName}}Arr({{$modelVarName}}PkArr)
	// 生成新的缓存
	New{{$modelName}}Repo(db.Manager).FindBy{{pkFuncName}}Arr({{$modelVarName}}PkArr)
	return nil
}

func (s *{{$modelName}}Cache) Update() error {
	// 获取返回的数组
	dataMap := s.tool.toMap(s.event)
	if len(dataMap) == 0 {
		return nil
	}
	// 更新
	var {{$modelVarName}}PkArr []{{pksType}}
	
	for _, data := range dataMap {
		if id, ok := data["id"]; ok && custom_array.IsInteger(id) {
			{{$modelVarName}}PkArr = append({{$modelVarName}}PkArr, custom_array.ToInt64(id))
		}
	}
	// 删除旧缓存
	New{{$modelName}}Repo(db.Manager).DeleteCacheBy{{pkFuncName}}Arr({{$modelVarName}}PkArr)
	// 生成新的缓存
	New{{$modelName}}Repo(db.Manager).FindBy{{pkFuncName}}Arr({{$modelVarName}}PkArr)
	return nil
}

func (s *{{$modelName}}Cache) Delete() error {
	// 获取返回的数组
	dataMap := s.tool.toMap(s.event)
	if len(dataMap) == 0 {
		return nil
	}
	// 删除
	var {{$modelVarName}}PkArr []{{pksType}}
	// 删除缓存
	for _, data := range dataMap {
		if id, ok := data["id"]; ok && custom_array.IsInteger(id) {
			{{$modelVarName}}PkArr = append({{$modelVarName}}PkArr, custom_array.ToInt64(id))
		}
	}
	New{{$modelName}}Repo(db.Manager).DeleteCacheBy{{pkFuncName}}Arr({{$modelVarName}}PkArr)
	return nil
}
