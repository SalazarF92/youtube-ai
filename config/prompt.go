package config

import "strings"

// type Prompt struct {
// 	comments []string
// }

func FormatPromptData(title string, description string, comments []string) string {

	if len(comments) == 0 {
		return ""
	}

	if title != "" {
		title = "title\n" + title
	}

	if description != "" {
		description = "description\n" + description
	}

	return strings.Join(comments, "\n") + "\n" + title + "\n" + description + "\n"
}
