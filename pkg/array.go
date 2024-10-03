package pkg

import (
	"strings"
)

func ArrayContains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func ArrayJoins(array []string, sep string, format func(string) string) string {
	for k, v := range array {
		if format != nil {
			array[k] = format(v)
		}
	}
	return strings.Join(array, sep)
}

func ArrayParamJoins(array [][]string) string {
	var temp []string
	for _, v := range array {
		if len(v) != 2 {
			panic("参数错误")
		}
		temp = append(temp, LineToLowCamel(v[0])+" "+v[1])
	}
	return strings.Join(temp, ",")
}
