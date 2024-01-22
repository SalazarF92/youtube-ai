package config

import "strings"

// type Prompt struct {
// 	comments []string
// }

func FormatComments(comments []string) string {
	if len(comments) == 0 {
		return ""
	}

	return strings.Join(comments, "\n")
}
