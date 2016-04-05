package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/images", listImagesView)
	r.POST("/containers", runContainersView)
	r.DELETE("/containers/:id", killContainerView)
	r.Run() // listen and server on 0.0.0.0:8080
}
