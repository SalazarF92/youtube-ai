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
	//PAREI AQUI: TROQUEI PARA MULTICONTENT, CONTINUAR AGORA PARA GERAR O PROMPT E FAZER A ANALISE DO CONTEUDO
	response, err := openai.OpenAi.CreateChatCompletion(
		context.Background(),
		opn.ChatCompletionRequest{
			Model: opn.GPT4,
			Messages: []opn.ChatCompletionMessage{
				{
					Role:    opn.ChatMessageRoleUser,
					Content: prompt,
					MultiContent: []opn.ChatMessagePart{
						{
							Text:     prompt,
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
