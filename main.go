package main

import (
	"context"
	"log"
	"market-openai/config"
	"market-openai/pkg/db"
	"market-openai/pkg/service"
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

type Services struct {
	ytService       *service.YoutubeService
	opnService      *service.OpenAiService
	analyserService *service.AnalyserService
}

func main() {
	envVarName := "GOOGLE_APPLICATION_CREDENTIALS"
	err := config.SetEnv(envVarName, "youtube.json")
	if err != nil {
		log.Fatalf("Erro ao configurar a variável de ambiente do google developers: %v", err)
	}

	var arrayVideos []string
	arrayVideos = append(arrayVideos, "1jbJpG3LufA", "UcTxU--Ga0U", "mN_Bxld5lg0")

	db := db.ConnectToDB(
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
	)
	defer db.Close()

	yt, err := youtube.NewService(context.Background())
	ytService := service.NewYoutube(yt)
	// promptTemplate := storage.PROMPTS["IMAGE"]
	opnStart := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	opnService := service.NewGPT(opnStart)
	analyserService := service.NewAnalyserService()

	services := Services{
		ytService:       ytService,
		opnService:      opnService,
		analyserService: analyserService,
	}
	RouterInitializer(&services)

}
