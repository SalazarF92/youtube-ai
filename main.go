package main

import (
	"context"
	"fmt"
	"log"
	"market-openai/config"
	"market-openai/pkg/db"
	"market-openai/pkg/entity"
	"market-openai/storage"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/api/youtube/v3"
)

type YoutubeResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
	} `json:"items"`
}

type VideoData struct {
	title    string
	desc     string
	comments []string
	thumb    string
}

func main() {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	db := db.ConnectToDB(
		db_host,
		db_port,
		db_user,
		db_password,
		db_name,
	)

	defer db.Close()

	envVarName := "GOOGLE_APPLICATION_CREDENTIALS"

	err := config.SetEnv(envVarName, "youtube.json")
	if err != nil {
		log.Fatalf("Erro ao configurar a variável de ambiente: %v", err)
	}

	yt, err := youtube.NewService(context.Background())

	if err != nil {
		panic(err)
	}
	ytService := entity.NewYoutube(yt)

	var videoData VideoData

	channel, err := ytService.GetChannelByHandle("SuperSaladin")

	if err != nil {
		panic(err)
	}

	videoList, err := ytService.GetVideosByChannelId(channel.Id.ChannelId)
	channelName := channel.Snippet.Title
	fmt.Println(channelName)

	if err != nil {
		fmt.Print(err)
	}

	for index, playlistItem := range videoList {

		video, err := ytService.GetVideoById(playlistItem.ContentDetails.VideoId)
		if err != nil {
			continue
		}

		fmt.Println(video.ContentDetails.Duration)

		if index == 2 {
			videoData.title = playlistItem.Snippet.Title
			videoData.desc = playlistItem.Snippet.Description
			videoData.thumb = playlistItem.Snippet.Thumbnails.Maxres.Url

			videoComments, err := ytService.GetCommentsByVideoId(playlistItem.ContentDetails.VideoId)
			if err != nil {
				panic(err)
			}

			for _, comment := range videoComments {
				videoData.comments = append(videoData.comments, comment.Snippet.TopLevelComment.Snippet.TextDisplay)
			}

		}

	}

	// fmt.Println(videoData.title)
	// fmt.Println("descrição", videoData.desc)
	// fmt.Println(videoData.comments)
	videoNumber := 1 //aqui seria o index percorrido  do array para numerar os videos por canal para comparar a thumb

	// fmt.Println(result)

	promptTemplate := storage.PROMPTS["NO_IMAGE"]

	fmt.Println(promptTemplate)

	prompt := config.FormatPromptData(videoNumber, channelName, videoData.title, videoData.desc, videoData.comments)

	opnStart := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	opnService := entity.NewGPT(opnStart)
	res := opnService.ChatCompletion(promptTemplate, prompt, videoData.thumb)
	fmt.Println("azedou", res)

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	// r.Run(":5000") // Por padrão, ele ouve na porta 8080
}
