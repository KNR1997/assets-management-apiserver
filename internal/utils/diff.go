package utils

import (
	"encoding/json"
	"regexp"
	"strings"
)

func CalculateDiff(oldData, newData []byte) string {
	diffMap := make(map[string]interface{})
	oldMap := make(map[string]interface{})
	newMap := make(map[string]interface{})
	_ = json.Unmarshal(oldData, &oldMap)
	_ = json.Unmarshal(newData, &newMap)

	oldMap = normalizeKeys(oldMap)
	newMap = normalizeKeys(newMap)

	for key, oldValue := range oldMap {
		if newValue, exists := newMap[key]; exists && oldValue != newValue {
			diffMap[key] = map[string]interface{}{
				"old": oldValue,
				"new": newValue,
			}
		}
	}

	diffJSON, _ := json.Marshal(diffMap)
	return string(diffJSON)
}

func convertToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func normalizeKeys(data map[string]interface{}) map[string]interface{} {
	normalized := make(map[string]interface{})
	for key, value := range data {
		normalized[convertToSnakeCase(key)] = value
	}
	return normalized
}
