package utils

import "encoding/json"

func ToJson(data interface{}) string {
	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return ""
	}
	return string(json)
}
