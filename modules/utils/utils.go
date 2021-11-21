package utils

import (
	"fmt"
	"sort"
)

type Array struct {
}

func (u Array) GetColumnByKey(array []map[string][]byte, key string) [][]byte {
	result := make([][]byte, 0)
	for _, item := range array {
		result = append(result, item[key])
	}
	return result
}

func (u Array) Join(array [][]byte, joinKey string) string {
	var result string
	for _, item := range array {
		result = fmt.Sprintf("%s%s%s", result, joinKey, item)
	}
	return result
}

func (u Array) GenerateMapByKey(array []map[string][]byte, key string) map[string]map[string][]byte {
	result := make(map[string]map[string][]byte)
	for _, item := range array {
		result[string(item[key])] = item
	}
	return result
}

func (u Array) In(target string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}
