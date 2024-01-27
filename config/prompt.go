package config

import (
	"strings"
)

// type Prompt struct {
// 	comments []string
// }

func FormatPromptData(channelName, title, description string, comments []string) string {

	if len(comments) == 0 {
		return ""
	}

	if title != "" {
		title = "title - " + title
	}

	if description != "" {
		description = "description - " + description
	}

	if channelName != "" {
		channelName = "channelName - " + channelName
	}

	//fmt.Sprintf with video+ videoNumber

	return "\n\n" + channelName + "\n\n" + title + "\n" + description + "\n\n" + "comments\n" + strings.Join(comments, "\n\n")
}
