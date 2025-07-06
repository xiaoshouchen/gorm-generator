package canal

import (
	"panda-trip/internal/model"
)

type {{.TableName | singular | upCamel}} struct {
	model.{{.TableName | singular | upCamel}}Cache
}

func New{{.TableName | singular | upCamel}}() *{{.TableName | singular | upCamel}} {
	return &{{.TableName | singular | upCamel}}{
		{{.TableName | singular | upCamel}}Cache: model.{{.TableName | singular | upCamel}}Cache{},
	}
}
