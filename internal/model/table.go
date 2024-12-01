package model

import (
	"github.com/xiaoshouchen/gorm-generator/pkg"
)

type Table struct {
	TableName    string            `json:"table_name"`
	Columns      []Column          `json:"fields"`
	Indexes      []Index           `json:"indexes"`
	WithCache    bool              `json:"with_cache"`    // 是否需要缓存
	CacheExpires int               `json:"cache_expires"` // 缓存时间，小于1则代表永久
	columnMap    map[string]Column // 字段类型和映射
}

func (t *Table) HasNullable() bool {
	for _, field := range t.Columns {
		if field.IsNullable == "YES" {
			return true
		}
	}
	return false
}

func (t *Table) HasTime() bool {
	for _, field := range t.Columns {
		if pkg.ArrayContains([]string{"datetime", "timestamp", "date"}, field.DataType) {
			return true
		}
	}
	return false
}

func (t *Table) GetColumnDataType(column string) string {
	return t.getColumnByName(column).DataType
}

func (t *Table) GetPks() []Column {
	var pk []Column
	for _, col := range t.Columns {
		if col.ColumnKey == "PRI" {
			pk = append(pk, col)
		}
	}
	return pk
}
func (t *Table) GetUniques() map[string][]Column {
	var uniques = make(map[string][]Column)
	for _, index := range t.Indexes {
		if index.NonUnique == 0 && index.IndexName != "PRIMARY" {
			name := index.IndexName
			uniques[name] = append(uniques[name], t.getColumnByName(index.ColumnName))
		}
	}
	return uniques
}

func (t *Table) GetIndexes() map[string][]Column {
	var indexes = make(map[string][]Column)
	for _, index := range t.Indexes {
		if index.NonUnique == 1 {
			name := index.IndexName
			indexes[name] = append(indexes[name], t.getColumnByName(index.ColumnName))
		}
	}
	return indexes
}

func (t *Table) getColumnByName(colName string) Column {
	if t.columnMap == nil {
		t.columnMap = make(map[string]Column)
		for _, col := range t.Columns {
			t.columnMap[col.ColumnName] = col
		}
	}
	return t.columnMap[colName]
}
