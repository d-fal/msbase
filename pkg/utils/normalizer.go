package utils

import (
	"math"
	"msbase/pkg/conf"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	ptime "github.com/yaa110/go-persian-calendar"
)

// LexiconSearch searchs for a character equivalence in lexicon
func LexiconSearch(char rune) string {
	if char <= 'z' {
		return string(char)
	}

	for _, c := range conf.GetConfigObject().GetLexicon()["lexicon"] {

		s, _ := strconv.Unquote(`"` + c.Unicode + `"`)
		if string(char) == s {
			return c.Name
		}
	}
	return ""
}

// NormalizeString normalize input string
func NormalizeString(strr string) string {
	result := ""

	for i := 0; i < utf8.RuneCountInString(strr); i++ {

		result += string(LexiconSearch([]rune(strr)[i]))
	}
	return result
}

// NormalizeTime time normalizer, it accepts various forms of date
func NormalizeTime(tm interface{}) interface{} {
	switch tm.(type) {

	case string:

		timeString := tm.(string)
		match, _ := regexp.MatchString(`^[0-9]{2}\/[0-9]{2}\/[0-9]{4}\s[0-9]{2}:[0-9]{2}:[0-9]{2}$`, timeString)

		if match {
			tm_, _ := time.Parse("01/02/2006 00:00:00", timeString)
			return tm_.Format(time.RFC3339)
		}
		match, _ = regexp.MatchString(`^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}[.]*([0-9]{3}|)Z$`, timeString)
		if match {
			tm_, _ := time.Parse(time.RFC3339, timeString)
			return tm_.Format(time.RFC3339)
		}
		match, _ = regexp.MatchString(`^[0-9]{4}\/[0-9]{1,2}\/[0-9]{1,2}$`, timeString)

		if match {
			tm_, _ := time.Parse("2006/01/02", timeString)
			if math.Abs(time.Now().Sub(tm_).Hours()) > 365*24*5 {
				values := strings.Split(timeString, "/")
				if len(values) > 2 {
					y, _ := strconv.Atoi(values[0])
					m, _ := strconv.Atoi(values[1])
					d, _ := strconv.Atoi(values[2])
					tm_ := ptime.Date(y, ptime.Month(m), d, 0, 0, 0, 0, ptime.Iran())
					return tm_.Time().Format(time.RFC3339)
				}

			}
		}

	}
	return nil
}
