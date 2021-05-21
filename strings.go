package gotils

import (
	"strconv"
	"strings"
)

// Indent the given string by an additional 4 spaces.
// Returns an empty string if the input is also empty.
func Indent(s string) string {
	if len(s) == 0 {
		return ""
	}
	prefix := "    "
	return prefix + strings.ReplaceAll(s, "\n", "\n"+prefix)
}

// QuoteJoin quotes all strings and then joins them with sep.
func QuoteJoin(s []string, sep string) string {
	quoted := make([]string, len(s)) // copy (do not modify original slice)
	for idx, str := range s {
		quoted[idx] = strconv.Quote(str)
	}
	return strings.Join(quoted, sep)
}
