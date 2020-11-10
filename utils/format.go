package utils

import (
	"fmt"
	"html"
	"strings"
	"time"
)

func formatX(x interface{}) interface{} {
	switch t := x.(type) {
	case time.Time:
		return t.UnixNano() / 1e6
	case int64:
		return t
	case float64:
		return t
	case string:
		return html.UnescapeString(t)
	case []byte:
		// this html unescapestring is used to decode the string if any special character comes
		return html.UnescapeString(string(t))
	default:
		return t
	}
}

func formatNumber(y interface{}) interface{} {
	switch t := y.(type) {
	case time.Time:
		return t.UnixNano() / 1e6
	case int64:
		return t
	case float64:
		return t
	case float32:
		return t
	case string:
		return 0
	case []byte:
		return 0
	default:
		return 0
	}
}

func parseFloat(n interface{}) float64 {
	switch t := n.(type) {
	case int64:
		return float64(t)
	case float64:
		return t
	default:
		return 0
	}
}

func formatString(d interface{}) string {
	switch t := d.(type) {
	case time.Time:
		return t.String()
	case int64:
		return fmt.Sprint(t)
	case float64:
		return fmt.Sprint(t)
	case float32:
		return fmt.Sprint(t)
	case string:
		return t
	case []byte:
		return string(t)
	default:
		return ""
	}
}

func getKey(m map[string]interface{}, keys []string, joinKey string) string {
	if len(keys) == 0 {
		return ""
	}

	oKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		sd := formatString(m[k])
		if len(sd) == 0 {
			continue
		}
		oKeys = append(oKeys, sd)
	}

	return strings.Join(oKeys, joinKey)
}
