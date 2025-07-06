package router

import (
	"fmt"
	"panda-trip/internal/enum/cache_enum"
	"panda-trip/pkg/cache"
	"strings"
	"sync"

	canalServer "panda-trip/internal/server/canal"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/siddontang/go/log"
)

type SyncEventHandler struct {
	canal.DummyEventHandler
}

func NewSyncEventHandler() *SyncEventHandler {
	return new(SyncEventHandler)
}

func (h *SyncEventHandler) OnRow(e *canal.RowsEvent) error {
	var err error
	if handler, ok := mysqlMap[e.Table.String()]; ok {
		handler.SetEvent(e)
		switch strings.ToLower(e.Action) {
		case "update":
			err = handler.Update()
		case "insert":
			err = handler.Insert()
		case "delete":
			err = handler.Delete()
		}
		if err != nil {
			log.Error(err.Error())
		}
		return err
	}
	return nil
}

// OnPosSynced 同步binlog的位置
func (h *SyncEventHandler) OnPosSynced(header *replication.EventHeader, pos mysql.Position, set mysql.GTIDSet, force bool) error {
	key := cache_enum.NewCanal().GetBinlogPositionKey()
	err := cache.GetOnce().Set(key, fmt.Sprintf("%s,%d", pos.Name, pos.Pos), -1)
	return err
}

func (h *SyncEventHandler) String() string {
	return "SyncEventHandler"
}

type CanalRouter struct {
	c *canal.Canal
}

var mysqlMap = make(map[string]canalServer.Canal)

func NewCanalRouter(c *canal.Canal) *CanalRouter {
	return &CanalRouter{
		c: c,
	}
}

func (r *CanalRouter) InitRouter() {
	var once sync.Once
	once.Do(func() {
		{{- range $k,$v := .tables}}
		mysqlMap["panda-trip.{{$k}}"] = canalServer.New{{$v}}()
		{{- end}}
	})

	// Mysql 事件
	r.c.SetEventHandler(NewSyncEventHandler())
}
