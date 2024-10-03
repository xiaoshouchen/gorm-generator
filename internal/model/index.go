package model

type Index struct {
	TableName  string `gorm:"column:TABLE_NAME"  json:"table_name"`
	NonUnique  int    `gorm:"column:NON_UNIQUE"  json:"non_unique"`
	IndexName  string `gorm:"column:INDEX_NAME"  json:"index_name"`
	SeqInIndex int    `gorm:"column:SEQ_IN_INDEX" json:"seq_in_index"`
	ColumnName string `gorm:"column:COLUMN_NAME"  json:"column_name"`
	Nullable   string `gorm:"column:NULLABLE"    json:"nullable"`
	IndexType  string `gorm:"column:INDEX_TYPE"  json:"index_type"`
}
