package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	_ "github.com/joho/godotenv/autoload"
	opn "github.com/sashabaranov/go-openai"
)

type openai interface {
	ChatCompletion(prompt string, videoContent VideoData) string
}

type OpenAiService struct {
	OpenAi *opn.Client
}

type VideoData struct {
	title    string
	desc     string
	comments []string
}

func NewGPT(openai *opn.Client) *OpenAiService {
	return &OpenAiService{
		OpenAi: openai,
	}
}

func (openai *OpenAiService) ChatCompletion(prompt string, videoContent string, thumb string) string {
	concatPrompt := fmt.Sprintf("%s %s", prompt, videoContent)

	fmt.Println(concatPrompt)
	fmt.Println(thumb)

	stream, err := openai.OpenAi.CreateChatCompletionStream(
		context.Background(),
		opn.ChatCompletionRequest{
			MaxTokens: 300,
			Stream:    true,
			Model:     opn.GPT4VisionPreview,
			Messages: []opn.ChatCompletionMessage{
				{
					Role: opn.ChatMessageRoleUser,
					MultiContent: []opn.ChatMessagePart{
						{
							Type: opn.ChatMessagePartTypeText,
							Text: concatPrompt,
						},
						{
							Type:     opn.ChatMessagePartTypeImageURL,
							ImageURL: &opn.ChatMessageImageURL{URL: thumb},
						},
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	var result string

	for {
		response, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			fmt.Println("stream finished")
			break
		}

		result += response.Choices[0].Delta.Content
	}

	return result

}
