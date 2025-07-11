package generator

import (
	"github.com/xiaoshouchen/gorm-generator/internal/model"
)

type Generator interface {
	DbTpl() string          // 获取数据库模板
	RepoTpl() string        // 获取仓库模板
	CacheTpl() string       // 获取缓存模板
	CanalTpl() string       // 获取 Canal 模板
	CanalRouterTpl() string // 获取 Canal Router 模板
}

func NewGenerator(config model.Config) Generator {
	var generator Generator
	switch config.Connect.Type {
	case "mysql":
		generator = NewMysql()
	default:
		panic("unknown generator type")
	}
	return generator
}
