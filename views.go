package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/docker/engine-api/client"
)


type RunContainersJSON struct {
	Image string `json:"image" binding:"required"`
	Expires int `json:"expires" binding:"max=1800,min=0"`
	Memory int `json:"memory" binding:"max=2048,min=1"`
}


func listImagesView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"images": listImages(),
	})
}


func listContainersView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"containers": listContainers(),
	})
}


func runContainersView(c *gin.Context) {
	var json RunContainersJSON
	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	memory := json.Memory
	if memory == 0 {
		memory = 16
	}

	ci, err := runContainer(json.Image, memory)
	if err != nil {
		if client.IsErrImageNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no such image",
			})
			return
		} else {
			panic(err)
		}
	}
	expires := json.Expires
	if expires != 0 {
		go timeoutKill(ci.ID, expires)  // To kill containers as async.
	}

	c.JSON(http.StatusOK, gin.H{
		"containerId": ci.ID,
		"image": json.Image,
		"expires": expires,
		"port": ci.Port,
		"ports": ci.Ports,
	})
}


func killContainerView(c *gin.Context) {
	id := c.Param("id")
	killContainer(id)
	c.Status(http.StatusNoContent)
}
