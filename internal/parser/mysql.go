package parser

import (
	"gorm-generator/internal/model"
	"gorm.io/gorm"
	"strings"
)

type Mysql struct {
	config model.Config
	db     *gorm.DB
}

func NewMysql(config model.Config, db *gorm.DB) *Mysql {
	return &Mysql{
		config: config,
		db:     db,
	}
}

func (m *Mysql) GetTableNames() []string {
	var tables []string
	m.db.Table("INFORMATION_SCHEMA.COLUMNS").
		Where("TABLE_SCHEMA = ?", m.config.Scheme).Pluck("TABLE_NAME", &tables)
	tables = m.config.FilterTables(tables)
	return tables
}

func (m *Mysql) GetTables(tableNames []string) []model.Table {
	var tables []model.Table
	var columns []model.Column
	// 获取表字段
	m.db.Table("INFORMATION_SCHEMA.COLUMNS").
		Select("TABLE_NAME,COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_COMMENT,COLUMN_KEY,COLUMN_DEFAULT").
		Where("TABLE_SCHEMA = ? and TABLE_NAME in (?)", m.config.Scheme, tableNames).
		Order("ORDINAL_POSITION").
		Find(&columns)
	var columnMap = make(map[string][]model.Column)
	for _, v := range columns {
		columnMap[v.TableName] = append(columnMap[v.TableName], v)
	}

	for k, v := range columnMap {
		tables = append(tables, model.Table{
			TableName: k,
			Columns:   v,
		})
	}

	var indexes []model.Index
	// 获取表索引
	m.db.Table("INFORMATION_SCHEMA.STATISTICS").
		Select("TABLE_NAME,NON_UNIQUE,INDEX_NAME,SEQ_IN_INDEX,COLUMN_NAME,INDEX_TYPE,NULLABLE").
		Where("TABLE_SCHEMA = ? and TABLE_NAME in (?)", m.config.Scheme, tableNames).
		Find(&indexes)
	var indexMap = make(map[string][]model.Index)
	for _, v := range indexes {
		indexMap[v.TableName] = append(indexMap[v.TableName], v)
	}

	for i, t := range tables {
		t.Indexes = indexMap[t.TableName]
		tables[i] = t
	}
	return tables
}

func (m *Mysql) TranslateDataType(column model.Column) string {
	dataType := strings.ToLower(column.DataType)
	var goType string
	switch dataType {
	case "smallint", "integer", "int", "bigint", "serial", "bigserial", "smallserial", "tinyint", "mediumint":
		goType = "int64"
	case "decimal", "numeric", "real", "double precision", "money", "float", "double":
		goType = "float64"
	case "text", "varchar", "character varying", "character", "char", "mediumtext", "time":
		goType = "string"
	case "boolean":
		goType = "bool"
	case "date", "timestamp", "datetime":
		goType = "time.Time"
	case "bit":
		goType = "[]byte"
	default:
		goType = "interface{}"
	}
	if strings.ToLower(column.IsNullable) == "yes" {
		goType = "*" + goType
	}
	// 处理软删除
	if strings.Contains(column.ColumnName, "deleted_at") {
		// 时间类型时，为NULL，则说明未删除
		if goType == "*time.Time" {
			goType = "gorm.DeletedAt"
		}
		// 整数类型时，必须不能为NULL
		if goType == "int64" {
			goType = "soft_delete.DeletedAt"
		}
	}
	return goType
}
