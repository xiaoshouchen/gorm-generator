package model

import "gorm-generator/pkg"

type Config struct {
	Scheme      string `json:"scheme"`
	PackageName string `json:"package_name"`
	Connect     struct {
		Type     string `json:"type"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"connect"`
	TableFilterOption string   `json:"table_filter_option"`
	Tables            []string `json:"tables"`
	CacheExpiration   []struct {
		Name string `json:"name"`
		Time int    `json:"time"`
	} `json:"cache_expiration"`
}

// WithCache 是否开启缓存
func (c Config) WithCache(tableName string) bool {
	for _, v := range c.CacheExpiration {
		if v.Name == tableName {
			return true
		}
	}
	return false
}

// CacheExpires 获取缓存过期时间
func (c Config) CacheExpires(tableName string) int {
	for _, v := range c.CacheExpiration {
		if v.Name == tableName {
			return v.Time
		}
	}
	return 0
}

// FilterTables 过滤表
func (c Config) FilterTables(tableNames []string) []string {
	if c.TableFilterOption == "all" {
		return tableNames
	}
	// 白名单模式
	if c.TableFilterOption == "whitelist" {
		temp := make([]string, 0)
		for _, v := range tableNames {
			if pkg.ArrayContains(c.Tables, v) {
				temp = append(temp, v)
			}
		}
		return temp
	}

	// 黑名单模式
	if c.TableFilterOption == "blacklist" {
		temp := make([]string, 0)
		for _, v := range tableNames {
			if !pkg.ArrayContains(c.Tables, v) {
				temp = append(temp, v)
			}
		}
		return temp
	}
	return tableNames
}
