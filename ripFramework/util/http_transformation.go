package util

import (
	"encoding/json"
	"strings"
)

func GetHttpVerb(request string) string {

	splitText := strings.Fields(request)

	if len(splitText) > 0 {
		return splitText[0]
	}

	return ""
}

func GetUri(request string) string {

	splitText := strings.Fields(request)

	if len(splitText) > 2 {
		return splitText[1]
	}

	return ""
}

func MapToJSON(data map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
