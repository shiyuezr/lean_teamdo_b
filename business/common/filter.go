package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ConvertToBeegoOrmFilter(filters map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	var key string
	var match string

	for k, v := range filters {
		if !(strings.HasPrefix(k, "__f")) {
			params[k] = v
		} else {
			keyString := strings.Split(k, "-")
			if len(keyString) == 3 {
				switch keyString[2] {
				case "equal":
					match = "exact"
				case "contain":
					match = "contains"
				case "gt", "gte", "lt", "lte", "in":
					match = keyString[2]
				case "range":
					val := v.([]interface{})
					start, _ := val[0].(json.Number).Int64()
					stop, _ := val[1].(json.Number).Int64()

					values := make([]int64, 0)
					for i := start; i < stop; i++ {
						values = append(values, i)
					}

					match = "in"
					v = values
				}
				key = fmt.Sprintf("%s__%s", keyString[1], match)
			} else {
				key = fmt.Sprintf("%s", keyString[1])
			}
			params[key] = v
		}
	}
	return params
}
