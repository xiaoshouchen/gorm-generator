package func_map

import (
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"github.com/xiaoshouchen/gorm-generator/internal/parser"
	"github.com/xiaoshouchen/gorm-generator/pkg"
	"strings"
)

type Index struct {
	table model.Table
	pa    parser.Parser
}

func NewIndex(table model.Table, pa parser.Parser) *Index {
	return &Index{
		table: table,
		pa:    pa,
	}
}

func (i *Index) FuncName(cols []model.Column) string {
	var colNames []string
	for _, col := range cols {
		colNames = append(colNames, pkg.LineToUpCamel(col.ColumnName))
	}
	return strings.Join(colNames, "And")
}
func (i *Index) Params(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName)+" "+i.pa.TranslateDataType(col))
	}
	return strings.Join(tempArr, ",")
}
func (i *Index) WhereCondition(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, col.ColumnName+" = ?")
	}
	return strings.Join(tempArr, " AND ")
}
func (i *Index) WhereArgs(cols []model.Column) string {
	var tempArr []string
	for _, col := range cols {
		tempArr = append(tempArr, pkg.LineToLowCamel(col.ColumnName))
	}
	return strings.Join(tempArr, ",")
}
