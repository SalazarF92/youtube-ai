package service

import (
	"fmt"
	"log"
	"market-openai/config"
	"market-openai/storage"

	"github.com/tiktoken-go/tokenizer"
)

type VideoDataAnalyser struct {
	channelTitle string
	title        string
	desc         string
	thumb        string
	comments     []string
}

type analyser interface {
	CheckTokens(prompt string) (value int)
	Run(videoId string, yt *YoutubeService, openai *OpenAiService)
	SetPrompt(channelTitle, videoTitle, videoDesc string, videoComments []string) (prompt string)
}

type AnalyserService struct {
	AnalyserService *analyser
}

func NewAnalyserService() *AnalyserService {
	return &AnalyserService{}
}

func (s *AnalyserService) CheckTokens(prompt string) (value int) {
	enc, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		log.Fatal(err)
	}
	tokens, _, _ := enc.Encode(prompt)
	return len(tokens)
}

func (s *AnalyserService) Run(videoId string, yt *YoutubeService, openai *OpenAiService) (bool, error) {
	videoData := VideoDataAnalyser{}

	video, err := yt.GetVideoById(videoId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(video.ContentDetails.Duration)

	videoData.channelTitle = video.Snippet.ChannelTitle
	videoData.title = video.Snippet.Title
	videoData.desc = video.Snippet.Description
	videoData.thumb = video.Snippet.Thumbnails.Maxres.Url

	videoComments, err := yt.GetCommentsByVideoId(videoId)
	if err != nil {
		panic(err)
	}

	for _, comment := range videoComments {
		videoData.comments = append(videoData.comments, comment.Snippet.TopLevelComment.Snippet.TextDisplay)
		yt.GetCommentsById(comment.Id)
		// myComment, err := yt.GetCommentsById(comment.Id)
		if err != nil {
			panic(err)
		}
		// fmt.Println(myComment.Snippet.TextDisplay)
		// fmt.Println(myComment.Snippet.LikeCount)

	}

	prompt := s.SetPrompt(videoData)
	fmt.Println(prompt)
	promptTemplate := storage.PROMPTS["IMAGE"]

	openaiAnalyse := openai.ChatCompletion(promptTemplate, prompt, videoData.thumb)

	fmt.Println(openaiAnalyse)

	return true, nil
}

func (s *AnalyserService) SetPrompt(props VideoDataAnalyser) string {
	prompt := config.FormatPromptData(props.channelTitle, props.title, props.desc, props.comments)
	return prompt
}
