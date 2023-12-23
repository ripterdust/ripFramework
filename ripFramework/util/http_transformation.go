package util

import "strings"

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
