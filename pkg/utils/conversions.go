package utils

import (
	"strconv"
)

func ConvertPayable(v interface{}) interface{} {

	switch v.(type) {
	case int:

		return v.(int)
	case float64:
		return int(v.(float64))
	case string:

		cval, err := strconv.ParseFloat(v.(string), 64)

		if err == nil {
			return cval
		}
	}
	return nil
}
