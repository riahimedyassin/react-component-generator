package lib

import "strings"

// return a string in the capitalize form
func Capitalize(value string) string {
	if len(value) == 0 {
		return ""
	}
	return strings.ToUpper(string(value[0])) + value[1:]
}
