package main

import (
	"context"
	"fmt"
	"log"
	"market-openai/config"
	"market-openai/pkg/entity"

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

// func getYoutubeData(apiKey, channelId string) (*YoutubeResponse, error) {
// 	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&channelId=%s&part=snippet", apiKey, channelId)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var ytResponse YoutubeResponse
// 	err = json.Unmarshal(body, &ytResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ytResponse, nil
// }

func getChannelInfoByHandle(service *youtube.Service, channelHandle string) (*youtube.SearchResult, error) {

	call := service.Search.List([]string{"snippet"}).Q(fmt.Sprintf("@%s", channelHandle)).Type("channel").MaxResults(1)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil

}

func getVideosFromPlaylist(service *youtube.Service, playlistId string) ([]*youtube.PlaylistItem, error) {
	call := service.PlaylistItems.List([]string{"snippet", "contentDetails"}).PlaylistId(playlistId)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func main() {

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

	channelId, err := ytService.GetChannelByHandle("SuperSaladin")

	if err != nil {
		panic(err)
	}

	videoList, err := ytService.GetVideosByChannelId(channelId.Id.ChannelId)

	for _, video := range videoList {
		fmt.Println(video.ContentDetails.VideoId)

		videoComments, err := ytService.GetCommentsByVideoId(video.ContentDetails.VideoId)
		if err != nil {
			panic(err)
		}

		for _, comment := range videoComments {
			fmt.Println(comment.Snippet.TopLevelComment.Snippet.TextDisplay)
		}

	}

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	// r.Run(":5000") // Por padrão, ele ouve na porta 8080
}
