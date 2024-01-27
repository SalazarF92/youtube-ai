package main

import (
	"market-openai/pkg/routes"

	"github.com/gin-gonic/gin"
)

func RouterInitializer(s *Services) {

	router := gin.Default()
	routes.UserRoutes(router)
	// routes.YoutubeRoutes(router, yt, opn)
	routes.AnalyserRoutes(router, s.analyserService, s.ytService, s.opnService)

	router.Run(":5000")
}
