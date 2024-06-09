package tg

import "strings"

func ExtractCmdArgs(text string) string {
	var result string

	idx := strings.Index(text, " ")
	if idx <= 0 {
		return result
	}

	return strings.TrimSpace(text[idx:])
}
