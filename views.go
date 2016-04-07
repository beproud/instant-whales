package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fsouza/go-dockerclient"
)


type RunContainersJSON struct {
	Image string `json:"image" binding:"required"`
	Expires int `json:"expires" binding:"required"`
	Memory int `json:"memory"`
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
	expires := json.Expires
	if expires < 0 || expires > 1800 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "expires shoud be in 0 <= expires <= 1800",
		})
		return
	}
	memory := json.Memory
	if memory == 0 {
		memory = 16
	}

	ci, err := runContainer(json.Image, memory)
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
