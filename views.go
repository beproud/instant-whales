package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/fsouza/go-dockerclient"
)


type RunContainersJSON struct {
	Image string `json:"image" binding:"required"`
	Expires string `json:"expires" binding:"required,numeric"`
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
			"error": "invalid_params",
		})
		return

	}
	expires, _ := strconv.Atoi(json.Expires)
	if expires <= 0 || expires > 600 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "expires shoud be in 0 < expires <= 600",
		})
		return
	}

	ci, err := runContainer(json.Image)
	if err != nil {
		if err == docker.ErrNoSuchImage {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "no such image",
			})
			return
		} else {
			panic(err)
		}
	}
	go timeoutKill(ci.ID, expires)  // To kill containers as async.

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
