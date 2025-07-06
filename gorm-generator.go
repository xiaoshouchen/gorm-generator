package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/xiaoshouchen/gorm-generator/internal/connector"
	"github.com/xiaoshouchen/gorm-generator/internal/func_map"
	"github.com/xiaoshouchen/gorm-generator/internal/generator"
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"github.com/xiaoshouchen/gorm-generator/internal/parser"
	"github.com/xiaoshouchen/gorm-generator/pkg"
)

var dataSource = flag.String("f", "data_source.json", "文件地址")

// 删除文件，支持模糊匹配
func deleteFile(path string) {
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		os.Remove(file)
	}
}

func main() {
	var configs = make([]model.Config, 0)
	getConfig(&configs)
	pl := pluralize.NewClient()

	for _, config := range configs {
		deleteFile(config.OutputPath + "*" + "_gen.go")
		deleteFile(config.OutputPath + "*" + "_cache.go")
		conn := connector.NewConnector(config)
		err := conn.Initialize()
		if err != nil {
			return
		}
		pa := parser.NewParser(config, conn.DB())
		tableNames := pa.GetTableNames()
		tables := pa.GetTables(tableNames)
		gen := generator.NewGenerator(config)
		var routerData = make(map[string]string)
		for _, table := range tables {
			var fm = func_map.GetFuncMap(config, table, pa)
			dbPath := config.OutputPath + pl.Singular(table.TableName) + "_gen.go"
			generateAndWrite(gen.DbTpl(), dbPath, fm, table, true)
			repoPath := config.OutputPath + pl.Singular(table.TableName) + "_repo.go"
			generateAndWrite(gen.RepoTpl(), repoPath, fm, table, false)
			if config.WithCache(table.TableName) {
				cachePath := config.OutputPath + pl.Singular(table.TableName) + "_cache.go"
				generateAndWrite(gen.CacheTpl(), cachePath, fm, table, true)
				routerData[table.TableName] = pkg.LineToUpCamel(pl.Singular(table.TableName))

				if config.CanalPath != "" {
					canalPath := config.CanalPath + pl.Singular(table.TableName) + ".go"
					generateAndWrite(gen.CanalTpl(), canalPath, fm, table, false)
				}
			}
		}
		if config.CanalRouterPath != "" {
			canalRouterPath := config.CanalRouterPath + "canal.go"
			generateAndWrite(gen.CanalRouterTpl(), canalRouterPath, nil, map[string]interface{}{
				"tables": routerData,
			}, true)
		}
	}
}

func generateAndWrite(tpl, path string, f template.FuncMap, data interface{}, overwrite bool) {
	t, err := template.New(path).Funcs(f).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	var buf = new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, path, data)
	if err != nil {
		log.Fatal(path, err)
	}
	source, err := format.Source(buf.Bytes())
	//source := buf.Bytes()
	if err != nil {
		log.Fatal(path, err)
	}
	// path 过滤文件名，文明的格式为xx.xx
	dirPath := path[:strings.LastIndex(path, "/")]

	if res, err1 := pkg.FileExists(dirPath); err1 == nil && !res {
		_ = os.Mkdir(dirPath, 0777)
	}
	if !overwrite {
		exists, err := pkg.FileExists(path)
		if err != nil {
			log.Println(path, err)
		}
		if exists {
			return
		}
	}
	err = os.WriteFile(path, source, 0664)
	if err != nil {
		log.Println(path, err)
	}
}

func getConfig(j *[]model.Config) {
	flag.Parse()
	tables, err := os.ReadFile(*dataSource)
	if err != nil {
		fmt.Println("please create a data_source.json first")
	}
	var configMap = make(map[string][]model.Config)
	err = json.Unmarshal(tables, &configMap)
	if err != nil {
		log.Fatal(err)
	}
	if configs, ok := configMap["databases"]; ok {
		*j = configs
	}
}
