package utils

import (
	"fmt"

	"msbase/pkg/conf"
)

//var log log.Logger

var (
	lexicon map[string][]conf.Chars
)

// Terminate when program interrupts
func Terminate() {
	// defer imagick.Terminate()
}
func init() {
	// imagick.Initialize()
	lexicon = conf.GetConfigObject().GetLexicon()
}

// IsEmpty checks if the input is not empty
func IsEmpty(input interface{}) bool {
	switch input.(type) {
	case string:
		if input.(string) == "" {
			return true
		}
	case interface{}:
		if input == nil {
			return true
		}
	}
	return false
}

// NormalizeInputValue normalize input
func NormalizeInputValue(input interface{}) interface{} {
	var result string

	switch input.(type) {
	case string:
		result = input.(string)
	case int:
		result = fmt.Sprintf("%d", input.(int))
	case float64:
		result = fmt.Sprintf("%f", input.(float64))

	}

	return result
}
