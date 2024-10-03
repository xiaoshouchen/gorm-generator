package connector

import (
	"fmt"
	"gorm-generator/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Mysql struct {
	db     *gorm.DB
	config model.Config
}

func NewMysql(config model.Config) *Mysql {
	return &Mysql{
		config: config,
	}
}

func (m *Mysql) Initialize() error {
	if m.db == nil {
		con := m.config.Connect
		dnsStr := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(dnsStr, con.User, con.Password, con.Host, con.Port, m.config.Scheme)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			log.Fatal(err)
		}
		m.db = db
	}
	return nil
}

func (m *Mysql) DB() *gorm.DB {
	return m.db
}
