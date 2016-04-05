package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/fsouza/go-dockerclient"
)


type RunContainersJSON struct {
	Image string `json:"image" binding:"required"`
	Timeout string `json:"timeout" binding:"required,numeric"`
}


func runContainersView(c *gin.Context) {
	var json RunContainersJSON
	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_params",
		})
		return

	}
	timeout, _ := strconv.Atoi(json.Timeout)
	if timeout <= 0 || timeout > 600 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "timeout shoud be in 0 < timeout <= 600",
		})
		return
	}

	id, err := runContainer(json.Image)
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
	go timeoutKill(id, timeout)  // To kill containers as async.

	c.JSON(http.StatusOK, gin.H{
		"containerId": id,
	})
}


func killContainerView(c *gin.Context) {
	id := c.Param("id")
	killContainer(id)
	c.Status(http.StatusNoContent)
}
