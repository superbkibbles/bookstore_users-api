package app

import "github.com/gin-gonic/gin"

// StartApplication: getting called in main.go

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	router.Run(":8080")
}
