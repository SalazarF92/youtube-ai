package entity

import (
	"errors"
	"fmt"

	yt "google.golang.org/api/youtube/v3"
)

type youtube interface {
	GetChannelById(channelId string) (*yt.Channel, error)
	GetChannelByHandle(handle string) (*yt.SearchResult, error)
	GetUploadsPlaylistId(channelId string) (string, error)
	GetVideosByChannelId(channelId string) ([]*yt.PlaylistItem, error)
}

type Youtube struct {
	Youtube *yt.Service
}

func NewYoutube(youtube *yt.Service) *Youtube {
	return &Youtube{
		Youtube: youtube,
	}
}

func (y *Youtube) GetChannelByHandle(handle string) (*yt.SearchResult, error) {
	call := y.Youtube.Search.List([]string{"snippet"}).Q(fmt.Sprintf("@%s", handle)).Type("channel").MaxResults(1)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil
}

func (y *Youtube) GetChannelById(channelId string) (*yt.Channel, error) {
	call := y.Youtube.Channels.List([]string{"snippet"}).Id(channelId)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil
}

func (y *Youtube) GetUploadsPlaylistId(channelId string) (string, error) {
	call := y.Youtube.Channels.List([]string{"contentDetails"}).Id(channelId)

	response, err := call.Do()
	if err != nil {
		return "", err
	}

	if len(response.Items) == 0 {
		return "", errors.New("nenhum canal encontrado com o ID fornecido")
	}

	return response.Items[0].ContentDetails.RelatedPlaylists.Uploads, nil
}

func (y *Youtube) GetVideosByChannelId(channelId string) ([]*yt.PlaylistItem, error) {
	uploadsPlaylistId, err := y.GetUploadsPlaylistId(channelId)
	if err != nil {
		return nil, err
	}

	fmt.Println(uploadsPlaylistId)

	// call := y.Youtube.Search.List([]string{"snippet"}).Type("video").MaxResults(10).ChannelId(channelId).Order("date")
	call := y.Youtube.PlaylistItems.List([]string{"snippet"}).PlaylistId(uploadsPlaylistId).MaxResults(50)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

// func (y *Youtube) GetVideos(channelId string) ([]string, error) {
// 	return nil, nil
// }
