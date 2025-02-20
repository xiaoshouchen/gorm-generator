package func_map

import (
	"strings"

	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"github.com/xiaoshouchen/gorm-generator/internal/parser"
	"github.com/xiaoshouchen/gorm-generator/pkg"
)

type Unique struct {
	table model.Table
	pa    parser.Parser
}

func NewUnique(table model.Table, pa parser.Parser) *Unique {
	return &Unique{
		table: table,
		pa:    pa,
	}
}

func (u *Unique) FuncName(cols []model.Column) string {
	var colNames []string
	for _, col := range cols {
		colNames = append(colNames, pkg.LineToUpCamel(col.ColumnName))
	}
	return strings.Join(colNames, "And")
}

func (u *Unique) Params(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName)+" "+u.pa.TranslateDataType(col))
	}
	return strings.Join(tempArr, ",")
}

func (u *Unique) CountParams(cols []model.Column) string {
	var tempArr []string
	for range cols {
		tempArr = append(tempArr, "%v")
	}
	return strings.Join(tempArr, "_")
}

func (u *Unique) WhereCondition(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, col.ColumnName+" = ?")
	}
	return strings.Join(tempArr, " AND ")
}

func (u *Unique) WhereArgs(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		if strings.HasPrefix(u.pa.TranslateDataType(col), "*") {
			tempArr = append(tempArr, "*"+pkg.LineToLowCamel(col.ColumnName))
		} else {
			tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName))
		}
	}
	return strings.Join(tempArr, ",")
}

func (u *Unique) CacheKeyFmtArgs(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		if strings.HasPrefix(u.pa.TranslateDataType(col), "*") {
			tempArr = append(tempArr, "*"+pkg.LineToLowCamel(col.ColumnName))
		} else {
			tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName))
		}
	}
	return strings.Join(tempArr, "_")
}

func (u *Unique) CacheKeyFmtValues(cols []model.Column) string {
	var arr []string
	for range cols {
		arr = append(arr, "%v")
	}
	return strings.Join(arr, "_")
}

func (u *Unique) UniquesType(cols []model.Column) string {
	if len(cols) == 1 {
		return u.pa.TranslateDataType(cols[0])
	}
	return "interface{}"
}

func (u *Unique) UniquesWhereCondition(cols []model.Column) string {
	var tempArr []string
	if len(cols) == 1 {
		return cols[0].ColumnName + " IN ?"
	}
	for _, col := range cols {
		tempArr = append(tempArr, col.ColumnName)
	}
	return "(" + strings.Join(tempArr, ",") + ") IN ?"
}

func (u *Unique) UniquesWhereArgs(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName))
	}
	return strings.Join(tempArr, ",")
}
