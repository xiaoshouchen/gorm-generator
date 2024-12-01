package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/xiaoshouchen/gorm-generator/internal/connector"
	"github.com/xiaoshouchen/gorm-generator/internal/func_map"
	"github.com/xiaoshouchen/gorm-generator/internal/generator"
	"github.com/xiaoshouchen/gorm-generator/internal/model"
	"github.com/xiaoshouchen/gorm-generator/internal/parser"
	"github.com/xiaoshouchen/gorm-generator/pkg"
	"go/format"
	"html/template"
	"log"
	"os"
)

var dataSource = flag.String("f", "data_source.json", "文件地址")
var outputSource = flag.String("o", "./model/", "文件输出位置")

func main() {
	var configs = make([]model.Config, 0)
	getConfig(&configs)
	pl := pluralize.NewClient()
	for _, config := range configs {
		conn := connector.NewConnector(config)
		err := conn.Initialize()
		if err != nil {
			return
		}
		pa := parser.NewParser(config, conn.DB())
		tableNames := pa.GetTableNames()
		tables := pa.GetTables(tableNames)
		gen := generator.NewGenerator(config)
		for _, table := range tables {
			var fm = func_map.GetFuncMap(config, table, pa)
			dbPath := *outputSource + pl.Singular(table.TableName) + "_gen.go"
			generateAndWrite(gen.DbTpl(), dbPath, fm, table, true)
			repoPath := *outputSource + pl.Singular(table.TableName) + "_repo.go"
			generateAndWrite(gen.RepoTpl(), repoPath, fm, table, false)
		}
	}

}

func generateAndWrite(tpl, path string, f template.FuncMap, table model.Table, overwrite bool) {
	t, err := template.New("tplFile").Funcs(f).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}
	var buf = new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "tplFile", table)
	if err != nil {
		log.Fatal(table.TableName, err)
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(table.TableName, err)
	}
	if res, err1 := pkg.FileExists(*outputSource); err1 == nil && !res {
		_ = os.Mkdir(*outputSource, 0777)
	}
	if !overwrite {
		exists, err := pkg.FileExists(path)
		if err != nil {
			log.Println(table.TableName, err)
		}
		if exists {
			return
		}
	}
	err = os.WriteFile(path, source, 0664)
	if err != nil {
		log.Println(table.TableName, err)
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
	if configs, ok := configMap["databases"]; ok {
		*j = configs
	}
}
