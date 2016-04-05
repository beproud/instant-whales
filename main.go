package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.POST("/containers", runContainersView)
	r.DELETE("/containers/:uid")
	r.Run() // listen and server on 0.0.0.0:8080
}
