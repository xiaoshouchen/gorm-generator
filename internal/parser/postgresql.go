package parser

import (
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"gorm.io/gorm"
	"strings"
)

type Postgresql struct {
	config model.Config
	db     *gorm.DB
}

func NewPostgresql(config model.Config, db *gorm.DB) *Postgresql {
	return &Postgresql{
		config: config,
		db:     db,
	}
}

func (p *Postgresql) GetTableNames() []string {
	var tables []string
	p.db.Raw(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = ? 
		AND table_type = 'BASE TABLE'`, 
		p.config.Scheme).Pluck("table_name", &tables)
	tables = p.config.FilterTables(tables)
	return tables
}

func (p *Postgresql) GetTables(tableNames []string) []model.Table {
	var tables []model.Table
	var columns []model.Column
	
	// Get table columns
	p.db.Raw(`
		SELECT 
			table_name as table_name,
			column_name as column_name,
			data_type as data_type,
			is_nullable as is_nullable,
			col_description((table_schema || '.' || table_name)::regclass::oid, ordinal_position) as column_comment,
			CASE 
				WHEN pk.column_name IS NOT NULL THEN 'PRI'
				ELSE ''
			END as column_key,
			column_default as column_default
		FROM information_schema.columns c
		LEFT JOIN (
			SELECT ku.column_name
			FROM information_schema.table_constraints tc
			JOIN information_schema.key_column_usage ku
				ON tc.constraint_name = ku.constraint_name
			WHERE tc.constraint_type = 'PRIMARY KEY'
			AND tc.table_schema = ?
			AND tc.table_name = ANY(?)
		) pk ON c.column_name = pk.column_name
		WHERE c.table_schema = ?
		AND c.table_name = ANY(?)
		ORDER BY ordinal_position`,
		p.config.Scheme, tableNames, p.config.Scheme, tableNames).
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
	// Get table indexes
	p.db.Raw(`
		SELECT 
			t.relname as table_name,
			CASE 
				WHEN i.indisunique = false THEN 1
				ELSE 0
			END as non_unique,
			ic.relname as index_name,
			a.attnum as seq_in_index,
			a.attname as column_name,
			CASE 
				WHEN i.indisclustered = true THEN 'CLUSTERED'
				ELSE 'BTREE'
			END as index_type,
			CASE 
				WHEN a.attnotnull = true THEN 'NO'
				ELSE 'YES'
			END as nullable
		FROM pg_class t
		INNER JOIN pg_index i ON t.oid = i.indrelid
		INNER JOIN pg_class ic ON i.indexrelid = ic.oid
		INNER JOIN pg_attribute a ON t.oid = a.attrelid AND a.attnum = ANY(i.indkey)
		INNER JOIN pg_namespace n ON n.oid = t.relnamespace
		WHERE n.nspname = ?
		AND t.relname = ANY(?)
		ORDER BY t.relname, ic.relname, a.attnum`,
		p.config.Scheme, tableNames).
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

func (p *Postgresql) TranslateDataType(column model.Column) string {
	dataType := strings.ToLower(column.DataType)
	var goType string
	switch dataType {
	case "smallint", "integer", "int", "bigint", "serial", "bigserial", "smallserial":
		goType = "int64"
	case "decimal", "numeric", "real", "double precision", "money", "float":
		goType = "float64"
	case "text", "varchar", "character varying", "character", "char", "time":
		goType = "string"
	case "boolean":
		goType = "bool"
	case "date", "timestamp", "timestamp without time zone", "timestamp with time zone":
		goType = "time.Time"
	case "bytea", "bit", "bit varying":
		goType = "[]byte"
	case "json", "jsonb":
		goType = "json.RawMessage"
	case "uuid":
		goType = "uuid.UUID"
	default:
		goType = "interface{}"
	}
	
	if strings.ToLower(column.IsNullable) == "yes" {
		goType = "*" + goType
	}
	
	// Handle soft delete
	if strings.Contains(column.ColumnName, "deleted_at") {
		if goType == "*time.Time" {
			goType = "gorm.DeletedAt"
		}
		if goType == "int64" {
			goType = "soft_delete.DeletedAt"
		}
	}
	return goType
}
