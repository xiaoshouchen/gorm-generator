package connector

import (
	"gorm-generator/internal/model"
	"gorm.io/gorm"
)

type Connector interface {
	Initialize() error
	DB() *gorm.DB
}

func NewConnector(config model.Config) Connector {
	var connector Connector
	switch config.Connect.Type {
	case "mysql":
		connector = NewMysql(config)
	default:
		panic("unknown connector type")
	}
	return connector
}
