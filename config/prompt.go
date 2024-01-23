package config

import (
	"fmt"
	"strings"
)

// type Prompt struct {
// 	comments []string
// }

func FormatPromptData(videoNumber int, channelName, title, description string, comments []string) string {

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

	video := fmt.Sprintf("video %d", videoNumber)

	//fmt.Sprintf with video+ videoNumber

	return "\n\n" + channelName + "\n" + video + "\n" + title + "\n" + description + "\n" + strings.Join(comments, "\n")
}
