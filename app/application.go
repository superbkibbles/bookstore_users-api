package app

import (
	"github.com/gin-gonic/gin"
	"github.com/superbkibbles/bookstore_users-api/logger"
)

// StartApplication: getting called in main.go

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start application")
	router.Run(":8080")
}
