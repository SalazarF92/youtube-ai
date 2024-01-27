package routes

import (
	"fmt"
	"market-openai/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnalyserRoutes(router *gin.Engine, analyser *service.AnalyserService, yt *service.YoutubeService, opn *service.OpenAiService) {
	ytRouter := router.Group("/analyser")
	{
		ytRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello Youtube",
			})
		})

		ytRouter.POST("/videos-analyze", func(c *gin.Context) {
			var data VideoData

			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// result, err := yt.GetVideoById(data.VideoIds[0])
			result, err := analyser.Run(data.VideoIds[0], yt, opn)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Dados recebidos com sucesso!", "data": result})

		})
	}

}
