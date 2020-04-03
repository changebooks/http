package http

import (
	"bytes"
	"strings"
)

// t.cn ["a" => "1", "b" => "2"] => t.cn?a=1&b=2
func UrlPath(path string, params map[string]string) string {
	if path == "" {
		return UrlQuery(params)
	}

	if strings.ContainsRune(path, '?') {
		path += "&"
	} else {
		path += "?"
	}

	return path + UrlQuery(params)
}

// ["a" => "1", "b" => "2"] => a=1&b=2
func UrlQuery(params map[string]string) string {
	if params == nil {
		return ""
	}

	var bucket bytes.Buffer

	first := true
	for key, value := range params {
		if first {
			first = false
		} else {
			bucket.WriteByte('&')
		}

		bucket.WriteString(key)
		bucket.WriteByte('=')
		bucket.WriteString(value)
	}

	return bucket.String()
}
