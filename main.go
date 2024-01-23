package main

import (
	"context"
	"fmt"
	"log"
	"market-openai/config"
	"market-openai/pkg/db"
	"market-openai/pkg/entity"
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

	channelId, err := ytService.GetChannelByHandle("SuperSaladin")

	if err != nil {
		panic(err)
	}

	videoList, err := ytService.GetVideosByChannelId(channelId.Id.ChannelId)

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

	formatedVideoData := config.FormatPromptData(videoData.title, videoData.desc, videoData.comments)

	// fmt.Println(result)

	// promptYoutubeAnalysis := `Considerando o canal ou canais a serem analisados
	// com o vídeo ou vídeos através do título ou títulos, descrição ou descrições, e comentários abaixo,
	// faça uma análise inteligente como um todo do conteúdo fornecido com uma sugestão de vídeo que eu poderia fazer com
	// base nos dados gerados a partir do conteúdo fornecido. Além disso, estou fornecendo a thumb ou thumbs dos vídeos, na ordem respectivas aos títulos para me dar sugestões de como eu poderia`

	opnStart := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	opnService := entity.NewGPT(opnStart)
	res := opnService.ChatCompletion("Conte uma história em até 50 caracteres", formatedVideoData, videoData.thumb)
	fmt.Println("azedou", res)

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	// r.Run(":5000") // Por padrão, ele ouve na porta 8080
}
