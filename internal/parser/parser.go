package parser

import (
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"gorm.io/gorm"
)

type Parser interface {
	TranslateDataType(column model.Column) string
	GetTableNames() []string
	GetTables(tableNames []string) []model.Table
}

func NewParser(config model.Config, db *gorm.DB) Parser {
	var parser Parser
	switch config.Connect.Type {
	case "mysql":
		parser = NewMysql(config, db)
	default:
		panic("unknown parser type")
	}
	return parser
}
