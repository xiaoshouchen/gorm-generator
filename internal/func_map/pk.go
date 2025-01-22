package func_map

import (
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"github.com/xiaoshouchen/gorm-generator/internal/parser"
	"github.com/xiaoshouchen/gorm-generator/pkg"
	"strings"
)

type Pk struct {
	table model.Table
	pa    parser.Parser
}

func NewPk(table model.Table, pa parser.Parser) *Pk {
	return &Pk{
		table: table,
		pa:    pa,
	}
}

func (p *Pk) FuncName() string {
	pks := p.table.GetPks()
	var colNames []string
	for _, col := range pks {
		colNames = append(colNames, pkg.LineToUpCamel(col.ColumnName))
	}
	return strings.Join(colNames, "And")
}

func (p *Pk) CacheKeyFmt() string {
	var arr []string
	for range p.table.GetPks() {
		arr = append(arr, "%v")
	}
	return strings.Join(arr, "_")
}

// Params findByPk params ,
// example: id int64
// example: key1 string , key2 string
func (p *Pk) Params() string {
	pks := p.table.GetPks()
	var tempArr []string
	for _, pk := range pks {
		tempArr = append(tempArr, pkg.LineToLowCamel(pk.ColumnName)+" "+p.pa.TranslateDataType(pk))
	}
	return strings.Join(tempArr, ",")
}

// WhereCondition findByPk where
// example: id = ? and key1 = ? and key2 = ?
func (p *Pk) WhereCondition() string {
	pks := p.table.GetPks()
	var tempArr []string
	for _, pk := range pks {
		tempArr = append(tempArr, pkg.LineToLowCamel(pk.ColumnName)+" = ?")
	}
	return strings.Join(tempArr, " AND ")
}

// WhereArgs findByPk where args
// example: id , key1 , key2
func (p *Pk) WhereArgs() []interface{} {
	pks := p.table.GetPks()
	var tempArr []interface{}
	for _, pk := range pks {
		tempArr = append(tempArr, pkg.LineToLowCamel(pk.ColumnName))
	}
	return tempArr
}

func (p *Pk) WhereArgsStr() string {
	pks := p.table.GetPks()
	var tempArr []string
	for _, pk := range pks {
		tempArr = append(tempArr, pkg.LineToLowCamel(pk.ColumnName))
	}
	return strings.Join(tempArr, ",")
}

func (p *Pk) PksParams() string {
	pks := p.table.GetPks()
	var colNames []string
	for _, col := range pks {
		colNames = append(colNames, pkg.LineToLowCamel(col.ColumnName))
	}
	return strings.Join(colNames, "And")
}

func (p *Pk) PksWhereCondition() string {
	pks := p.table.GetPks()
	if len(pks) == 1 {
		return pks[0].ColumnName + " IN ?"
	}
	var tempArr []string
	for _, pk := range pks {
		tempArr = append(tempArr, pk.ColumnName)
	}
	return "(" + strings.Join(tempArr, ",") + ") IN ?"
}

func (p *Pk) PksWhereArgs() string {
	pks := p.table.GetPks()
	if len(pks) == 1 {
		return pkg.LineToLowCamel(pks[0].ColumnName)
	}
	return "pks"
}

func (p *Pk) PksType() string {
	pks := p.table.GetPks()
	if len(pks) == 1 {
		return p.pa.TranslateDataType(pks[0])
	}
	return "interface{}"
}

func (p *Pk) PksFields(prefix string) string {
	pks := p.table.GetPks()
	var tempArr []string
	for _, pk := range pks {
		tempArr = append(tempArr, prefix+"."+pkg.LineToUpCamel(pk.ColumnName))
	}
	return strings.Join(tempArr, ",")
}
