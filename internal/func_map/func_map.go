package func_map

import (
	"github.com/gertd/go-pluralize"
	"gorm-generator/internal/model"
	"gorm-generator/internal/parser"
	"gorm-generator/pkg"
	"html/template"
	"strings"
)

func GetFuncMap(config model.Config, table model.Table, pa parser.Parser) template.FuncMap {
	pl := pluralize.NewClient()
	fm := template.FuncMap{
		"lowCamel":       pkg.LineToLowCamel,
		"upCamel":        pkg.LineToUpCamel,
		"inline":         pkg.Inline,
		"singular":       pl.Singular,
		"plural":         pl.Plural,
		"containsNumber": pkg.ContainsNumber,
		"joins":          pkg.ArrayJoins,
		"paramJoins":     pkg.ArrayParamJoins,
	}
	fm["transType"] = pa.TranslateDataType
	fm["hasTime"] = table.HasTime

	fm["cacheTtl"] = func() int {
		return table.CacheExpires
	}

	// pk func
	pk := NewPk(table, pa)
	fm["pkFuncName"] = pk.FuncName
	fm["pkParams"] = pk.Params
	fm["pkWhereCondition"] = pk.WhereCondition
	fm["pkWhereArgs"] = pk.WhereArgs

	// pks func
	fm["pksParams"] = pk.PksParams
	fm["pksWhereCondition"] = pk.PksWhereCondition
	fm["pksWhereArgs"] = pk.PksWhereArgs
	fm["pksType"] = pk.PksType
	fm["pksFields"] = pk.PksFields

	uni := NewUnique(table, pa)

	fm["uniques"] = table.GetUniques
	fm["uniqueFuncName"] = uni.FuncName
	fm["uniqueParams"] = uni.Params
	fm["uniqueWhereCondition"] = uni.WhereCondition
	fm["uniqueWhereArgs"] = uni.WhereArgs
	fm["uniquesType"] = uni.UniquesType
	fm["uniquesWhereCondition"] = uni.UniquesWhereCondition
	fm["uniquesWhereArgs"] = uni.UniquesWhereArgs

	index := NewIndex(table, pa)

	fm["indexes"] = table.GetIndexes
	fm["indexFuncName"] = index.FuncName
	fm["indexParams"] = index.Params
	fm["indexWhereCondition"] = index.WhereCondition
	fm["indexWhereArgs"] = index.WhereArgs

	fm["withCache"] = func() bool {
		return config.WithCache(table.TableName)
	}
	fm["softDelete"] = func() int {
		var has int
		for _, v := range table.Columns {
			t := pa.TranslateDataType(v)
			if strings.Contains(t, "soft_delete.DeletedAt") {
				has = 1
			}
			if strings.Contains(t, "gorm.DeletedAt") {
				has = 2
			}
		}
		return has
	}

	return fm
}
