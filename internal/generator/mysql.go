package generator

import (
	_ "embed"
)

type Mysql struct {
}

func NewMysql() *Mysql {
	return &Mysql{}
}

//go:embed template/mysql/model.tpl
var mysqlModel string

//go:embed template/mysql/field.tpl
var mysqlField string

//go:embed template/mysql/insert.tpl
var mysqlInsert string

//go:embed template/mysql/omit.tpl
var mysqlOmit string

//go:embed template/mysql/find.tpl
var mysqlFind string

//go:embed template/mysql/repo.tpl
var mysqlRepo string

//go:embed template/mysql/find_by_index.tpl
var mysqlFindByIndex string

//go:embed template/mysql/find_by_pk.tpl
var mysqlFindByPk string

//go:embed template/mysql/find_by_unique.tpl
var mysqlFindByUnique string

//go:embed template/mysql/count.tpl
var mysqlCount string

//go:embed template/mysql/delete.tpl
var mysqlDelete string

func (m *Mysql) DbTpl() string {
	return mysqlModel + mysqlField + mysqlInsert + mysqlOmit +
		mysqlFind + mysqlFindByPk + mysqlFindByUnique + mysqlFindByIndex + mysqlCount + mysqlDelete
}

func (m *Mysql) RepoTpl() string {
	return mysqlRepo
}
