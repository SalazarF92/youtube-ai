package service

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
	GetCommentsByVideoId(videoId string) ([]*yt.CommentThread, error)
}

type YoutubeService struct {
	YoutubeService *yt.Service
}

func NewYoutube(youtube *yt.Service) *YoutubeService {
	return &YoutubeService{
		YoutubeService: youtube,
	}
}

func (y *YoutubeService) GetChannelByHandle(handle string) (*yt.SearchResult, error) {
	call := y.YoutubeService.Search.List([]string{"snippet"}).Q(fmt.Sprintf("@%s", handle)).Type("channel").MaxResults(1)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil
}

func (y *YoutubeService) GetChannelById(channelId string) (*yt.Channel, error) {
	call := y.YoutubeService.Channels.List([]string{"snippet"}).Id(channelId)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil
}

func (y *YoutubeService) GetUploadsPlaylistId(channelId string) (string, error) {
	call := y.YoutubeService.Channels.List([]string{"contentDetails"}).Id(channelId)

	response, err := call.Do()
	if err != nil {
		return "", err
	}

	if len(response.Items) == 0 {
		return "", errors.New("nenhum canal encontrado com o ID fornecido")
	}

	return response.Items[0].ContentDetails.RelatedPlaylists.Uploads, nil
}

func (y *YoutubeService) GetVideosByChannelId(channelId string) ([]*yt.PlaylistItem, error) {
	uploadsPlaylistId, err := y.GetUploadsPlaylistId(channelId)
	if err != nil {
		return nil, err
	}

	// call := y.Youtube.Search.List([]string{"snippet"}).Type("video").MaxResults(10).ChannelId(channelId).Order("date")
	call := y.YoutubeService.PlaylistItems.List([]string{"snippet", "contentDetails"}).PlaylistId(uploadsPlaylistId).MaxResults(5)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (y *YoutubeService) GetCommentsByVideoId(videoId string) ([]*yt.CommentThread, error) {
	call := y.YoutubeService.CommentThreads.List([]string{"snippet"}).VideoId(videoId).MaxResults(10).Order("relevance")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (y *YoutubeService) GetCommentsById(commentId string) (*yt.Comment, error) {
	call := y.YoutubeService.Comments.List([]string{"snippet"}).Id(commentId)

	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], nil

}

func (y *YoutubeService) GetVideoById(videoId string) (*yt.Video, error) {
	call := y.YoutubeService.Videos.List([]string{"snippet", "contentDetails"}).Id(videoId)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items[0], nil
}
