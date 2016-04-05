package main

import (
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	if os.Getenv("INSTANT_WHALES_RELEASE_MODE") == "1" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.GET("/images", listImagesView)
	r.GET("/containers", listContainersView)
	r.POST("/containers", runContainersView)
	r.DELETE("/containers/:id", killContainerView)
	r.Run() // listen and server on 0.0.0.0:8080
}
