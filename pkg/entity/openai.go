package entity

import (
	"context"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	opn "github.com/sashabaranov/go-openai"
)

type openai interface {
	ChatCompletion(prompt string, videoContent VideoData) string
}

type OpenAi struct {
	OpenAi *opn.Client
}

type VideoData struct {
	title    string
	desc     string
	comments []string
}

func NewGPT(openai *opn.Client) *OpenAi {
	return &OpenAi{
		OpenAi: openai,
	}
}

func (openai *OpenAi) ChatCompletion(prompt string, videoContent string, thumb string) string {
	concatPrompt := fmt.Sprintf("%s %s", prompt, videoContent)

	fmt.Println(concatPrompt)
	fmt.Println(thumb)

	response, err := openai.OpenAi.CreateChatCompletion(
		context.Background(),
		opn.ChatCompletionRequest{
			Model: opn.GPT4,
			Messages: []opn.ChatCompletionMessage{
				{
					Role: opn.ChatMessageRoleUser,
					// Content: "diga olá",
					MultiContent: []opn.ChatMessagePart{
						{
							Type:     opn.ChatMessagePartTypeText,
							Text:     "Consegue analisar a imagem que forneci aqui?",
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
	return response.Choices[0].Message.Content

}
