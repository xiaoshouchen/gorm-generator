package model

type Column struct {
	TableName     string `json:"table_name" gorm:"column:TABLE_NAME"`
	ColumnName    string `json:"column_name" gorm:"column:COLUMN_NAME"`
	DataType      string `json:"data_type" gorm:"column:DATA_TYPE"`
	IsNullable    string `json:"is_nullable" gorm:"column:IS_NULLABLE"`
	ColumnComment string `json:"column_comment" gorm:"column:COLUMN_COMMENT"`
	ColumnKey     string `json:"column_key" gorm:"column:COLUMN_KEY"`
	ColumnDefault string `json:"column_default" gorm:"column:COLUMN_DEFAULT"`
}
