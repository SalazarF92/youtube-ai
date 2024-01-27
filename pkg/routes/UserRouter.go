package routes

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	ytRouter := router.Group("/user")
	{
		ytRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello User",
			})
		})
	}

}

// var videoData VideoData
// for index, videoId := range arrayVideos {

// 	video, err := ytService.GetVideoById(videoId)
// 	if err != nil {
// 		continue
// 	}
// 	fmt.Println(video.ContentDetails.Duration)

// 	if index == 2 {
// 		videoData.title = video.Snippet.Title
// 		videoData.desc = video.Snippet.Description
// 		videoData.thumb = video.Snippet.Thumbnails.Maxres.Url

// 		videoComments, err := ytService.GetCommentsByVideoId(videoId)
// 		if err != nil {
// 			panic(err)
// 		}

// 		for _, comment := range videoComments {
// 			videoData.comments = append(videoData.comments, comment.Snippet.TopLevelComment.Snippet.TextDisplay)
// 			ytService.GetCommentsById(comment.Id)
// 			// myComment, err := ytService.GetCommentsById(comment.Id)
// 			if err != nil {
// 				panic(err)
// 			}
// 			// fmt.Println(myComment.Snippet.TextDisplay)
// 			// fmt.Println(myComment.Snippet.LikeCount)

// 		}

// 		prompt := config.FormatPromptData(video.Snippet.ChannelTitle, videoData.title, videoData.desc, videoData.comments)
// 		enc, err := tokenizer.Get(tokenizer.Cl100kBase)
// 		if err != nil {
// 			panic("oh oh")
// 		}

// 		ids, _, _ := enc.Encode(prompt)
// 		fmt.Println("maqueisso", len(ids))
// 	}

// 	// res := opnService.ChatCompletion(promptTemplate, prompt, videoData.thumb)
// 	// fmt.Println("azedou", res)
// }
