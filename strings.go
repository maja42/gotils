package gotils

import "strings"

// Indent the given string by an additional 4 spaces.
// Returns an empty string if the input is also empty.
func Indent(s string) string {
	if len(s) == 0 {
		return ""
	}
	prefix := "    "
	return prefix + strings.ReplaceAll(s, "\n", "\n"+prefix)
}
