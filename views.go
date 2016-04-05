package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type ContainersJSON struct {
	Container string `json:"container" binding:"required"`
	Timeout string `json:"timeout" binding:"required,numeric"`
}


func createContainersView(c *gin.Context) {
	var json ContainersJSON
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

	id := runContainer()
	go timeoutKill(id, timeout)  // To kill containers as async.

	c.JSON(http.StatusOK, gin.H{
		"containerId": id,
	})
}
